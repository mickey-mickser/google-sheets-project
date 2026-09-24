package handler_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/handler"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Capabilities(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("OK", func(t *testing.T) {
		// Подготавливаем временную директорию и файл capabilities.json
		tmp := t.TempDir()
		oldWD, _ := os.Getwd()
		defer os.Chdir(oldWD)
		_ = os.Chdir(tmp)

		okJSON := []byte(`{"status":"success","features":["share","get","batchUpdate"]}`)
		err := os.WriteFile(filepath.Join(tmp, "capabilities.json"), okJSON, 0o644)
		assert.NoError(t, err)

		h := &handler.Handler{}
		r := gin.New()
		r.GET("/capabilities", h.Capabilities)

		req := httptest.NewRequest(http.MethodGet, "/capabilities", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"status":"success"`)
	})

	t.Run("Not found -> 404", func(t *testing.T) {
		tmp := t.TempDir()
		oldWD, _ := os.Getwd()
		defer os.Chdir(oldWD)
		_ = os.Chdir(tmp) // Файл не создаём

		h := &handler.Handler{}
		r := gin.New()
		r.GET("/capabilities", h.Capabilities)

		req := httptest.NewRequest(http.MethodGet, "/capabilities", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), `file not found`)
	})

	t.Run("Invalid JSON -> 400", func(t *testing.T) {
		tmp := t.TempDir()
		oldWD, _ := os.Getwd()
		defer os.Chdir(oldWD)
		_ = os.Chdir(tmp)

		// Кладём битый JSON
		err := os.WriteFile(filepath.Join(tmp, "capabilities.json"), []byte(`{not-json`), 0o644)
		assert.NoError(t, err)

		h := &handler.Handler{}
		r := gin.New()
		r.GET("/capabilities", h.Capabilities)

		req := httptest.NewRequest(http.MethodGet, "/capabilities", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `file is not valid JSON`)
	})
}
