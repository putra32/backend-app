package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(rw http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(js)

	return nil
}
