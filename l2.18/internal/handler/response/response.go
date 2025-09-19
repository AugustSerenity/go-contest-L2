package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

type SuccessResponce struct {
	Result interface{} `json:"result,omitempty"`
}

func SendError(w http.ResponseWriter, code int, err string) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err})
}

func SendSuccess(w http.ResponseWriter, code int, result interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(SuccessResponce{Result: result})
}
