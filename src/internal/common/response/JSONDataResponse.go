package response

import (
	"encoding/json"
	"net/http"
)

func JSONDataResponse(rw http.ResponseWriter, status int, data interface{}, rest ...interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)

	var statusBool bool

	if status >= 200 && status <= 299 {
		statusBool = true
	} else {
		statusBool = false
	}

	responseData := map[string]interface{}{
		"status": statusBool,
		"data":   data,
	}

	// Add rest arguments to the response
	for i := 0; i < len(rest); i += 2 {
		if i+1 < len(rest) {
			key, ok := rest[i].(string)
			if !ok {
				// Skip if the key is not a string
				continue
			}
			value := rest[i+1]
			responseData[key] = value
		}
	}

	json.NewEncoder(rw).Encode(responseData)
}
