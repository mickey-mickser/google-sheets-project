package handler_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/drive/v3"

	mock_sheetUsecase "github.com/mickey-mickser/google-sheets-project/pkg/mocks"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/handler"
)

func setupShare(t *testing.T) (*gin.Engine, *mock_sheetUsecase.MockSheetUseCase) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockUC := mock_sheetUsecase.NewMockSheetUseCase(ctrl)

	h := &handler.Handler{
		Log:      logrus.New(),
		SheetUse: mockUC,
	}

	r := gin.New()
	r.POST("/share", h.Share)

	return r, mockUC
}

func TestHandler_Share(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		r, mockUC := setupShare(t)

		perm := &drive.Permission{Id: "perm-123"}
		mockUC.EXPECT().
			SharePermission(gomock.Any(), gomock.AssignableToTypeOf(models.ShareRequest{})).
			DoAndReturn(func(_ context.Context, req models.ShareRequest) (*drive.Permission, error) {
				if req.Email != "user@test.com" {
					t.Fatalf("unexpected email: %s", req.Email)
				}
				// sendEmail=false допускается без emailMessage
				return perm, nil
			})

		body := []byte(`{"email":"user@test.com","sendEmail":false}`)
		req := httptest.NewRequest(http.MethodPost, "/share", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code, w.Body.String())
		assert.Contains(t, w.Body.String(), `"status":"ok"`)
		assert.Contains(t, w.Body.String(), `"id":"perm-123"`)
	})

	t.Run("Missing email -> 400", func(t *testing.T) {
		r, _ := setupShare(t)

		body := []byte(`{"sendEmail":false}`)
		req := httptest.NewRequest(http.MethodPost, "/share", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "email")
	})

	t.Run("Invalid email -> 400", func(t *testing.T) {
		r, _ := setupShare(t)

		body := []byte(`{"email":"not-an-email"}`)
		req := httptest.NewRequest(http.MethodPost, "/share", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "invalid email format")
	})

	t.Run("sendEmail=true без emailMessage -> 400", func(t *testing.T) {
		r, _ := setupShare(t)

		body := []byte(`{"email":"user@test.com","sendEmail":true}`)
		req := httptest.NewRequest(http.MethodPost, "/share", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "emailMessage")
	})

	t.Run("UseCase error -> 500", func(t *testing.T) {
		r, mockUC := setupShare(t)

		mockUC.EXPECT().
			SharePermission(gomock.Any(), gomock.AssignableToTypeOf(models.ShareRequest{})).
			Return(nil, errors.New("service error"))

		body := []byte(`{"email":"user@test.com"}`)
		req := httptest.NewRequest(http.MethodPost, "/share", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
