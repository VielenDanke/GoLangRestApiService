package apiserver

import (
	"encoding/json"
	"net/http"
)

// JSONResponse ...
type JSONResponse struct {
	http.ResponseWriter
	code int
}

// CreateJSONResponse ...
func (jr *JSONResponse) CreateJSONResponse(status int, body interface{}) {
	jr.ResponseWriter.Header().Set("Content-Type", "application/json")
	jr.ResponseWriter.WriteHeader(status)
	json.NewEncoder(jr).Encode(body)
}
