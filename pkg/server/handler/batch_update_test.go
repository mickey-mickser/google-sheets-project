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
	"testing"
)

func TestHandler_BatchUpdate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockusecase.NewMockSheetUseCase(ctrl)

	h := Handler{
		Log:      logrus.New(),
		SheetUse: mockUseCase,
	}

	router := gin.Default()
	router.POST("/batch-update", h.BatchUpdate)

	t.Run("Success", func(t *testing.T) {
		reqBody := models.BatchUpdateRequest{
			Updates: []models.UpdateRequest{
				{Range: "A1", Value: "test"},
			},
		}
		jsonBody, _ := json.Marshal(reqBody)

		mockUseCase.EXPECT().
			BatchUpdate(gomock.Any(), "12345", reqBody.Updates).
			Return(&sheets.BatchUpdateValuesResponse{}, nil)

		req := httptest.NewRequest(http.MethodPost, "/batch-update?sheet_id=12345", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"status":"ok","data":{}}`, rec.Body.String())
	})

	t.Run("Missing sheet_id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/batch-update", bytes.NewReader([]byte(`{}`)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "sheet_id")
	})

	t.Run("Invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/batch-update?sheet_id=12345", bytes.NewReader([]byte(`{invalid json}`)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "invalid character")
	})

	t.Run("Usecase returns error", func(t *testing.T) {
		reqBody := models.BatchUpdateRequest{
			Updates: []models.UpdateRequest{
				{Range: "A1", Value: "test"},
			},
		}
		jsonBody, _ := json.Marshal(reqBody)

		mockUseCase.EXPECT().
			BatchUpdate(gomock.Any(), "12345", reqBody.Updates).
			Return(nil, assert.AnError)

		req := httptest.NewRequest(http.MethodPost, "/batch-update?sheet_id=12345", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var got map[string]string
		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatalf("failed to unmarshal response: %v; body=%s", err, rec.Body.String())
		}

		assert.Equal(t, "Failed to update the Google Sheets", got["message"])
		assert.Equal(t, assert.AnError.Error(), got["error"])
	})

}
