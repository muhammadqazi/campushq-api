package response

import (
	"encoding/json"
	"net/http"
)

func JSONDataResponse(rw http.ResponseWriter, status int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)

	var statusBool bool

	if status >= 200 && status <= 299 {
		statusBool = true
	} else {
		statusBool = false
	}

	json.NewEncoder(rw).Encode(
		map[string]interface{}{
			"status": statusBool,
			"data":   data,
		},
	)
}
