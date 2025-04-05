package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockusecase "github.com/mickey-mickser/google-sheets-project/pkg/usecase/sheets/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/sheets/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_BatchUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockusecase.NewMockSheetUseCase(ctrl)
	r := gin.Default()
	log := logrus.New()

	h := Handler{
		Log:      log,
		SheetUse: mockUseCase,
	}

	r.POST("/batch-update", h.BatchUpdate)

	tests := []struct {
		name           string
		sheetId        string
		body           string
		mockReturn     interface{}
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success",
			sheetId:        "12345",
			body:           `{"updates": [{"range": "A1", "value": "test"}]}`,
			mockReturn:     &sheets.BatchUpdateSpreadsheetResponse{},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"ok","data":{}}`,
		},
		{
			name:           "Missing sheet_id",
			sheetId:        "",
			body:           `{"updates": [{"range": "A1", "value": "test"}]}`,
			mockReturn:     nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Sheet ID parameter is required"}`,
		},
		{
			name:           "Invalid request body",
			sheetId:        "12345",
			body:           `{invalid json}`,
			mockReturn:     nil,
			mockError:      fmt.Errorf("failed to unmarshal batchUpdateRequest body: invalid character 'i' looking for beginning of value"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"failed to unmarshal batchUpdateRequest body: invalid character 'i' looking for beginning of value"}`,
		},
		{
			name:           "Failed to update sheet",
			sheetId:        "12345",
			body:           `{"updates": [{"range": "A1", "value": "test"}]}`,
			mockReturn:     nil,
			mockError:      fmt.Errorf("update error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"update error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.sheetId != "" {
				mockUseCase.EXPECT().
					BatchUpdate(gomock.Any(), tt.sheetId, []byte(tt.body)).
					Return(tt.mockReturn, tt.mockError)
			}

			req := httptest.NewRequest(http.MethodPost, "/batch-update?sheet_id="+tt.sheetId, strings.NewReader(tt.body))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
