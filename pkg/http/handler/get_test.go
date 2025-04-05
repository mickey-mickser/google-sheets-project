package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockusecase "github.com/mickey-mickser/google-sheets-project/pkg/usecase/sheets/mocks"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mockusecase.NewMockSheetUseCase(ctrl)
	r := gin.Default()
	log := logrus.New()

	h := Handler{
		Log:      log,
		SheetUse: mockUseCase,
	}

	r.GET("/get", h.Get)

	t.Run("Success", func(t *testing.T) {
		mockUseCase.EXPECT().Get(gomock.Any(), "12345", "A1:B2").Return([][]interface{}{{"data"}}, nil)

		req := httptest.NewRequest(http.MethodGet, "/get?sheet_id=12345&range=A1:B2", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "ok")
	})

	t.Run("Missing sheet_id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Sheet ID parameter is required")
	})

	t.Run("Google Sheets error", func(t *testing.T) {
		mockUseCase.EXPECT().Get(gomock.Any(), "12345", "A1:B2").Return(nil, errors.New("service error"))

		req := httptest.NewRequest(http.MethodGet, "/get?sheet_id=12345&range=A1:B2", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "service error")
	})
}
