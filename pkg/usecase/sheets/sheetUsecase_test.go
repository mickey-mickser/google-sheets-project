package sheetUsecase_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mickey-mickser/google-sheets-project/pkg/config"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/mickey-mickser/google-sheets-project/pkg/usecase/sheets"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// newTestServer returns httptest.Server emulating Google Sheets API
func newTestServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		// Emulate batchUpdate request
		case r.URL.Path == "/v4/spreadsheets/spreadsheetId/values:batchUpdate":
			w.Header().Set("Content-Type", "application/json")
			if _, err := w.Write([]byte(`{"spreadsheetId":"spreadsheetId","totalUpdatedCells":1}`)); err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			t.Logf("Request: %s %s", r.Method, r.URL.Path)
		// Emulate getting values ​​from a sheet range
		case r.URL.Path == "/v4/spreadsheets/spreadsheetId/values/A1:Z1000":
			w.Header().Set("Content-Type", "application/json")
			if _, err := w.Write([]byte(`{"range":"A1:Z1000","values":[["val1"],["val2"]]}`)); err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			t.Logf("Request: %s %s", r.Method, r.URL.Path)
		// Emulate clearing a range of values
		case r.URL.Path == "/v4/spreadsheets/spreadsheetId/values/A1:Z1000:clear":
			w.Header().Set("Content-Type", "application/json")
			if _, err := w.Write([]byte(`{"clearedRange":"A1:Z1000"}`)); err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			t.Logf("Request: %s %s", r.Method, r.URL.Path)
		// Emulate creating a new spreadsheet
		case r.URL.Path == "/v4/spreadsheets":
			w.Header().Set("Content-Type", "application/json")
			if _, err := w.Write([]byte(`{"spreadsheetId":"newId","spreadsheetUrl":"url","properties":{"title":"Test"}}`)); err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			t.Logf("Request: %s %s", r.Method, r.URL.Path)
		// Emulate setting permissions on a file in Google Drive (for example, to access a spreadsheet)
		case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/drive/v3/files/") && strings.HasSuffix(r.URL.Path, "/permissions"):
			w.Header().Set("Content-Type", "application/json")
			if _, err := w.Write([]byte(`{"id":"permissionId"}`)); err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			t.Logf("Request: %s %s", r.Method, r.URL.Path)

		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)

		}
	}))

}

// TestBatchUpdate_Integration checks the correct execution of the BatchUpdate request
func TestBatchUpdate_Integration(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	srv, err := sheets.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	uc := sheetUsecase.NewSheetUse(srv, logrus.New(), &config.PermissionsStruct{})

	updates := []models.UpdateRequest{
		{
			Range:  "A1:B1",
			Values: [][]interface{}{{"Test1", "Test2"}},
		},
	}
	resp, err := uc.BatchUpdate(context.Background(), "spreadsheetId", updates)
	if err != nil {
		t.Fatalf("BatchUpdate error: %v", err)
	}
	if resp.TotalUpdatedCells != 1 {
		t.Errorf("expected 1 updated cell, got %d", resp.TotalUpdatedCells)
	}
}

// TestGet_Integration checks for getting values from a table
func TestGet_Integration(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	srv, _ := sheets.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())
	uc := sheetUsecase.NewSheetUse(srv, logrus.New(), &config.PermissionsStruct{})

	values, err := uc.Get(context.Background(), "spreadsheetId", "A1:Z1000")
	if err != nil {
		t.Fatalf("Get error: %v", err)
	}
	if len(values) != 2 || values[0][0] != "val1" {
		t.Errorf("unexpected values: %v", values)
	}
}

// TestDelete_Integration checks if the range of values in the table is cleared
func TestDelete_Integration(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	srv, _ := sheets.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())
	uc := sheetUsecase.NewSheetUse(srv, logrus.New(), &config.PermissionsStruct{})

	resp, err := uc.Delete(context.Background(), "spreadsheetId", "A1:Z1000")
	if err != nil {
		t.Fatalf("Delete error: %v", err)
	}
	if resp.ClearedRange != "A1:Z1000" {
		t.Errorf("expected cleared range A1:Z1000, got %s", resp.ClearedRange)
	}
}

// mockPermissionSetter - a mock for PermissionSetter that does nothing.
// Used to replace the real permission setter in tests.
type mockPermissionSetter struct{}

func (m mockPermissionSetter) SetPermission(ctx context.Context, sheetID string) error {
	return nil
}

// TestCreate_Integration checks the creation of a new table and the setting of permissions
func TestCreate_Integration(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	srv, _ := sheets.NewService(context.Background(), option.WithEndpoint(ts.URL), option.WithoutAuthentication())
	uc := sheetUsecase.NewSheetUse(srv, logrus.New(), &config.PermissionsStruct{})

	// Replace PermissionSetter with a mock to avoid making real requests to Google Drive
	uc.SetPermissionSetter(&mockPermissionSetter{})

	req := models.CreateRequest{
		Properties: struct {
			Title string `json:"title" binding:"required"`
		}(struct{ Title string }{Title: "Test"}),
	}
	resp, err := uc.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Create error: %v", err)
	}
	if resp.SpreadsheetId != "newId" {
		t.Errorf("expected SpreadsheetId newId, got %s", resp.SpreadsheetId)
	}
	if resp.Properties.Title != "Test" {
		t.Errorf("expected Title Test, got %s", resp.Properties.Title)
	}

}
