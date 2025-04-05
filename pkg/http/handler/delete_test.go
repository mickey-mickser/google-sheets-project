package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockusecase "github.com/mickey-mickser/google-sheets-project/pkg/usecase/sheets/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/sheets/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSheetUseCase := mockusecase.NewMockSheetUseCase(ctrl)

	h := &Handler{
		SheetUse: mockSheetUseCase,
		Log:      logrus.New(),
	}

	tests := []struct {
		name           string
		sheetID        string
		rangeParam     string
		body           []byte
		mockReturn     *sheets.ClearValuesResponse
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful Delete",
			sheetID:        "123",
			rangeParam:     "Sheet1!A1:B10",
			body:           []byte(`{"example": "data"}`),
			mockReturn:     &sheets.ClearValuesResponse{},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"ok","data":{}}`,
		},
		{
			name:           "Bad Request - No SheetID",
			sheetID:        "",
			rangeParam:     "Sheet1!A1:B10",
			body:           []byte(`{"example": "data"}`),
			mockReturn:     nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Sheet ID parameter is required"}`,
		},
		{
			name:           "Internal Server Error",
			sheetID:        "123",
			rangeParam:     "Sheet1!A1:B10",
			body:           []byte(`{"example": "data"}`),
			mockReturn:     nil,
			mockError:      fmt.Errorf("Internal error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Internal error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, "/delete?sheet_id="+tt.sheetID+"&range="+tt.rangeParam, bytes.NewReader(tt.body))
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			if tt.sheetID != "" {
				mockSheetUseCase.EXPECT().
					Delete(gomock.Any(), tt.sheetID, tt.rangeParam, tt.body).
					Return(tt.mockReturn, tt.mockError)
			}

			h.Delete(c)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}

}
