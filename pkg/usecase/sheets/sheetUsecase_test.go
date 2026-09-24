package sheetUsecase

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mickey-mickser/google-sheets-project/pkg/config"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// --------- test http server, эмулирующий Google APIs ---------

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		// Sheets: batchUpdate
		case r.Method == http.MethodPost && r.URL.Path == "/v4/spreadsheets/spreadsheetId/values:batchUpdate":
			_, _ = w.Write([]byte(`{"spreadsheetId":"spreadsheetId","totalUpdatedCells":1}`))
			return

		// Sheets: Get default range
		case r.Method == http.MethodGet && r.URL.Path == "/v4/spreadsheets/spreadsheetId/values/A1:Z1000":
			_, _ = w.Write([]byte(`{"range":"A1:Z1000","values":[["val1"],["val2"]]}`))
			return

		// Sheets: Clear values
		case r.Method == http.MethodPost && r.URL.Path == "/v4/spreadsheets/spreadsheetId/values/A1:Z1000:clear":
			_, _ = w.Write([]byte(`{"clearedRange":"A1:Z1000"}`))
			return

		// Sheets: Create spreadsheet
		case r.Method == http.MethodPost && r.URL.Path == "/v4/spreadsheets":
			_, _ = w.Write([]byte(`{"spreadsheetId":"newId","spreadsheetUrl":"url","properties":{"title":"Test"}}`))
			return

		// Drive: create permission
		case r.Method == http.MethodPost && r.URL.Path == "/drive/v3/files/testSheetID/permissions":
			// Симулируем 409-конфликт (уже расшарено), если пришло EmailMessage=conflict
			if r.URL.Query().Get("emailMessage") == "conflict" {
				w.WriteHeader(http.StatusConflict)
				_ = json.NewEncoder(w).Encode(map[string]any{
					"error": map[string]any{
						"code":    409,
						"message": "alreadyShared",
					},
				})
				return
			}
			_, _ = w.Write([]byte(`{"id":"permissionId"}`))
			return
		}

		t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
	}))
}

// --------- вспомогательные сущности ---------

// Заглушка PermissionSetter, чтобы Create не дергал Drive
type stubPermissionSetter struct{}

func (s stubPermissionSetter) SetPermission(ctx context.Context, sheetID string) error {
	return nil
}

// --------- тесты ---------

func TestBatchUpdate_Integration(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	sheetsSvc, err := sheets.NewService(context.Background(),
		option.WithEndpoint(ts.URL),
		option.WithoutAuthentication(),
	)
	if err != nil {
		t.Fatalf("sheets.NewService: %v", err)
	}
	driveSvc, err := drive.NewService(context.Background(),
		option.WithEndpoint(ts.URL),
		option.WithoutAuthentication(),
	)
	if err != nil {
		t.Fatalf("drive.NewService: %v", err)
	}

	uc := NewSheetUse(sheetsSvc, driveSvc, logrus.New(), &config.PermissionsStruct{}, "testSheetID")

	updates := []models.UpdateRequest{
		{Range: "A1:B1", Values: [][]interface{}{{"x", "y"}}},
	}
	resp, err := uc.BatchUpdate(context.Background(), "spreadsheetId", updates)
	if err != nil {
		t.Fatalf("BatchUpdate error: %v", err)
	}
	if resp.TotalUpdatedCells != 1 {
		t.Errorf("expected TotalUpdatedCells=1, got %d", resp.TotalUpdatedCells)
	}
}

func TestBatchUpdate_InvalidParam(t *testing.T) {
	// updates == nil -> ошибка валидации, без HTTP
	uc := &sheetUseCase{log: logrus.New()}
	_, err := uc.BatchUpdate(context.Background(), "spreadsheetId", nil)
	if err == nil || !strings.Contains(err.Error(), "invalid parameter") {
		t.Fatalf("expected validation error, got %v", err)
	}
}

func TestGet_DefaultRange_Integration(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	sheetsSvc, _ := sheets.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())
	driveSvc, _ := drive.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())

	uc := NewSheetUse(sheetsSvc, driveSvc, logrus.New(), &config.PermissionsStruct{}, "testSheetID")

	values, err := uc.Get(context.Background(), "spreadsheetId", "")
	if err != nil {
		t.Fatalf("Get error: %v", err)
	}
	if len(values) != 2 || values[0][0] != "val1" {
		t.Errorf("unexpected values: %#v", values)
	}
}

func TestDelete_Integration(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	sheetsSvc, _ := sheets.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())
	driveSvc, _ := drive.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())

	uc := NewSheetUse(sheetsSvc, driveSvc, logrus.New(), &config.PermissionsStruct{}, "testSheetID")

	resp, err := uc.Delete(context.Background(), "spreadsheetId", "A1:Z1000")
	if err != nil {
		t.Fatalf("Delete error: %v", err)
	}
	if resp.ClearedRange != "A1:Z1000" {
		t.Errorf("expected cleared A1:Z1000, got %s", resp.ClearedRange)
	}
}

func TestCreate_Integration(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	sheetsSvc, _ := sheets.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())
	driveSvc, _ := drive.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())

	uc := NewSheetUse(sheetsSvc, driveSvc, logrus.New(), &config.PermissionsStruct{}, "testSheetID")
	// Заменяем реальный выставитель прав на заглушку
	uc.SetPermissionSetter(stubPermissionSetter{})

	req := models.CreateRequest{
		Properties: struct {
			Title string `json:"title" binding:"required"`
		}{Title: "Test"},
	}
	resp, err := uc.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Create error: %v", err)
	}
	if resp.SpreadsheetId != "newId" {
		t.Errorf("expected SpreadsheetId=newId, got %s", resp.SpreadsheetId)
	}
	if resp.Properties.Title != "Test" {
		t.Errorf("expected title Test, got %s", resp.Properties.Title)
	}
}

func TestCreate_ValidationError(t *testing.T) {
	uc := &sheetUseCase{log: logrus.New()}
	_, err := uc.Create(context.Background(), models.CreateRequest{
		Properties: struct {
			Title string `json:"title" binding:"required"`
		}{Title: ""},
	})
	if err == nil || !strings.Contains(err.Error(), "invalid parameter") {
		t.Fatalf("expected validation error, got %v", err)
	}
}

func TestSharePermission_Validation(t *testing.T) {
	uc := &sheetUseCase{log: logrus.New()}
	_, err := uc.SharePermission(context.Background(), models.ShareRequest{Email: ""})
	if err == nil || !strings.Contains(err.Error(), "email is required") {
		t.Fatalf("expected email is required error, got %v", err)
	}
}

func TestSharePermission_SendEmailFalse_OK(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	sheetsSvc, _ := sheets.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())
	driveSvc, _ := drive.NewService(
		context.Background(),
		option.WithEndpoint(strings.TrimRight(ts.URL, "/")+"/drive/v3/"),
		option.WithoutAuthentication(),
	)
	uc := NewSheetUse(sheetsSvc, driveSvc, logrus.New(), &config.PermissionsStruct{}, "testSheetID")

	p, err := uc.SharePermission(context.Background(), models.ShareRequest{
		Email:        "user@example.com",
		SendEmail:    false,
		EmailMessage: "",
	})
	if err != nil {
		t.Fatalf("SharePermission error: %v", err)
	}
	if p == nil || p.Id != "permissionId" {
		t.Fatalf("unexpected permission: %#v", p)
	}
}

func TestSharePermission_SendEmailTrue_WithMessage_OK(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	sheetsSvc, _ := sheets.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())
	driveSvc, _ := drive.NewService(
		context.Background(),
		option.WithEndpoint(strings.TrimRight(ts.URL, "/")+"/drive/v3/"),
		option.WithoutAuthentication(),
	)
	uc := NewSheetUse(sheetsSvc, driveSvc, logrus.New(), &config.PermissionsStruct{}, "testSheetID")

	p, err := uc.SharePermission(context.Background(), models.ShareRequest{
		Email:        "user@example.com",
		SendEmail:    true,
		EmailMessage: "hello",
	})
	if err != nil {
		t.Fatalf("SharePermission error: %v", err)
	}
	if p == nil || p.Id != "permissionId" {
		t.Fatalf("unexpected permission: %#v", p)
	}
}

func TestSharePermission_Conflict409_ReturnsEmptyID_NoError(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	sheetsSvc, _ := sheets.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())
	driveSvc, _ := drive.NewService(
		context.Background(),
		option.WithEndpoint(strings.TrimRight(ts.URL, "/")+"/drive/v3/"),
		option.WithoutAuthentication(),
	)
	uc := NewSheetUse(sheetsSvc, driveSvc, logrus.New(), &config.PermissionsStruct{}, "testSheetID")

	// Триггерим 409 через emailMessage=conflict (см. newTestServer)
	p, err := uc.SharePermission(context.Background(), models.ShareRequest{
		Email:        "user@example.com",
		SendEmail:    true,
		EmailMessage: "conflict",
	})
	if err != nil {
		t.Fatalf("expected no error on 409, got %v", err)
	}
	if p == nil || p.Id != "" {
		t.Fatalf("expected empty permission id on 409, got %#v", p)
	}
}
