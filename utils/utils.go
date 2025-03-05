package utils

import (
	"encoding/json"
	"net/http"
)

// returns status followed by the data, e.g book/loan
// can be further strengthened in future to return book/loan data type instead of simply interface.
func SendResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
