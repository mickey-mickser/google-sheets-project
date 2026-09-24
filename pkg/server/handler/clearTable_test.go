package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockusecase "github.com/mickey-mickser/google-sheets-project/pkg/mocks"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/sheets/v4"
)

func TestDeleteHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUC := mockusecase.NewMockSheetUseCase(ctrl)
	h := &Handler{
		SheetUse: mockUC,
		Log:      logrus.New(),
	}

	// We configure the router exactly the same way as in the real code
	router := gin.New()
	router.POST("/clearTable", h.ClearTable)

	tests := []struct {
		name           string
		sheetID        string
		rangeParam     string
		reqBody        models.DeleteRequest
		mockReturn     *sheets.ClearValuesResponse
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful Delete",
			sheetID:        "123",
			rangeParam:     "Sheet1!A1:B10",
			reqBody:        models.DeleteRequest{Range: "Sheet1!A1:B10"},
			mockReturn:     &sheets.ClearValuesResponse{},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"ok","data":{}}`,
		},
		{
			name:           "Missing sheet_id",
			sheetID:        "",
			rangeParam:     "",
			reqBody:        models.DeleteRequest{},
			mockReturn:     nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			// BadRequest → responses.NewMissingParamError("sheet_id")
			expectedBody: `{"detail":"sheet_id is required","code":"005"}`,
		},
		{
			name:           "Internal error from usecase",
			sheetID:        "123",
			rangeParam:     "A1:Z1",
			reqBody:        models.DeleteRequest{Range: "A1:Z1"},
			mockReturn:     nil,
			mockError:      fmt.Errorf("some error"),
			expectedStatus: http.StatusInternalServerError,
			// InternalError → responses.InternalErrorResponse (detail="internal error", code="000")
			expectedBody: `{"detail":"Something bad happened","code":"000"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.sheetID != "" {
				if tt.mockError != nil {
					mockUC.
						EXPECT().
						Delete(gomock.Any(), tt.sheetID, tt.reqBody.Range).
						Return(nil, tt.mockError)
				} else {
					mockUC.
						EXPECT().
						Delete(gomock.Any(), tt.sheetID, tt.reqBody.Range).
						Return(tt.mockReturn, nil)
				}
			}

			bodyBytes, _ := json.Marshal(tt.reqBody)
			url := "/clearTable"
			if tt.sheetID != "" {
				url += "?sheet_id=" + tt.sheetID
			}
			if tt.rangeParam != "" {
				url += "&range=" + tt.rangeParam
			}

			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}
