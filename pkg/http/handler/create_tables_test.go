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

func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockusecase.NewMockSheetUseCase(ctrl)
	r := gin.Default()
	log := logrus.New()

	h := Handler{
		Log:      log,
		SheetUse: mockUseCase,
	}

	r.POST("/create", h.Create)

	tests := []struct {
		name           string
		body           string
		mockReturn     interface{}
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success",
			body:           `{"sheet_name": "Sheet1"}`,
			mockReturn:     &sheets.Spreadsheet{},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"ok","data":{}}`,
		},
		{
			name:           "Invalid request body",
			body:           `{invalid json}`,
			mockReturn:     nil,
			mockError:      fmt.Errorf("failed to unmarshal createRequest body: invalid character 'i' looking for beginning of value"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"failed to unmarshal createRequest body: invalid character 'i' looking for beginning of value"}`,
		},
		{
			name:           "Failed to create table",
			body:           `{"sheet_name": "Sheet1"}`,
			mockReturn:     nil,
			mockError:      fmt.Errorf("creation error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"creation error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase.EXPECT().
				Create(gomock.Any(), []byte(tt.body)).
				Return(tt.mockReturn, tt.mockError)

			req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(tt.body))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
