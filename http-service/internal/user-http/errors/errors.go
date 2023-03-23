package errors

import (
	"encoding/json"
	"net/http"
	httpModels "wbHTTPServer/http-service/internal/user-http/models"
)

// проброс ошибок во внешку, т.к. тест задание, а не продакшен

func ErrStatusInternal(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&httpModels.ErrorJSON{Error: err.Error()})
}

func ErrStatusBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&httpModels.ErrorJSON{Error: err.Error()})
}
