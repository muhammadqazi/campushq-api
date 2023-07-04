package response

import (
	"encoding/json"
	"net/http"
)

func JSONMessageResponse(rw http.ResponseWriter, status int, message string) {
	rw.WriteHeader(status)
	var statusBool bool
	if status == http.StatusOK || status == http.StatusCreated {
		statusBool = true

	} else {
		statusBool = false
	}

	json.NewEncoder(rw).Encode(
		map[string]interface{}{
			"status":  statusBool,
			"message": message,
		},
	)
}
