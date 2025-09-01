package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockusecase "github.com/mickey-mickser/google-sheets-project/pkg/mocks"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/sheets/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockusecase.NewMockSheetUseCase(ctrl)
	log := logrus.New()

	h := Handler{
		Log:      log,
		SheetUse: mockUseCase,
	}

	r := gin.Default()
	r.POST("/create", h.Create)

	t.Run("Success", func(t *testing.T) {
		reqBody := `{"properties":{"title":"Sheet1"}}`

		mockUseCase.EXPECT().
			Create(gomock.Any(), models.CreateRequest{
				Properties: models.Properties{Title: "Sheet1"},
			}).
			Return(&sheets.Spreadsheet{}, nil)

		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"status":"ok","data":{}}`, rec.Body.String())
	})

	t.Run("Invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(`{invalid json}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var got map[string]string
		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatalf("failed to unmarshal response: %v; body=%s", err, rec.Body.String())
		}
		assert.Equal(t, "006", got["code"])
		assert.Contains(t, got["detail"], "invalid character 'i' looking for beginning of object key string")
	})

	t.Run("Failed to create table", func(t *testing.T) {
		reqBody := `{"properties":{"title":"Sheet1"}}`

		mockUseCase.EXPECT().
			Create(gomock.Any(), models.CreateRequest{
				Properties: models.Properties{Title: "Sheet1"},
			}).
			Return(nil, assert.AnError)

		req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var got map[string]string
		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatalf("failed to unmarshal response: %v; body=%s", err, rec.Body.String())
		}
		assert.Equal(t, "000", got["code"])
		assert.Equal(t, "Something bad happened", got["detail"])
	})
}
