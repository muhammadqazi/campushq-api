package response

import (
	"encoding/json"
	"net/http"
)

func JSONDataResponse(rw http.ResponseWriter, status int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)

	json.NewEncoder(rw).Encode(data)
}
