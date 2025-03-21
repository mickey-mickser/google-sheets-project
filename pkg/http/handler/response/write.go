package responses

import (
	"encoding/json"
	"net/http"
)

func WriteBadRequest(w http.ResponseWriter, err ResponseError) {
	write(w, http.StatusBadRequest, err)
}

func WritePaymentRequiredRequest(w http.ResponseWriter, err ResponseError) {
	write(w, http.StatusPaymentRequired, err)
}

func WriteInternalErr(w http.ResponseWriter) {
	write(w, http.StatusInternalServerError, InternalErrorResponse)
}

func WriteUnauthorizedErr(w http.ResponseWriter, err ResponseError) {
	write(w, http.StatusUnauthorized, err)
}

func WriteNotFound(w http.ResponseWriter, body interface{}) {
	write(w, http.StatusNotFound, body)
}

func WriteConflict(w http.ResponseWriter, body interface{}) {
	write(w, http.StatusConflict, body)
}

func WriteOK(w http.ResponseWriter, body interface{}) {
	write(w, http.StatusOK, body)
}

func WriteNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func write(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
