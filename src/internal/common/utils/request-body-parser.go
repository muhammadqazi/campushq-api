package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RequestBodyParser(r *http.Request, data interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &data)
	if err != nil {
		if parseErr, ok := err.(*json.UnmarshalTypeError); ok {
			return fmt.Errorf("%v: must be a %v", parseErr.Field, parseErr.Type.Name())
		}
		return err
	}

	return nil
}
