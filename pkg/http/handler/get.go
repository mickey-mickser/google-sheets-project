package handler

import (
	responses "github.com/mickey-mickser/telegram-project/pkg/http/handler/response"
	"net/http"
)

func (h handler) Get(w http.ResponseWriter, r *http.Request) {

	table := r.URL.Query().Get("table")
	rangeParam := r.URL.Query().Get("range")
	if table == "" {
		http.Error(w, "Table parameter is required", http.StatusBadRequest)
		return
	}

	asd, err := h.sheetUse.Get(r.Context(), table, rangeParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responses.WriteOK(w, map[string]interface{}{
		"status": "ok",
		"data":   asd,
	})
}
