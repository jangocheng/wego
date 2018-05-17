package net

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, obj interface{}) error {
	w.WriteHeader(status)
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	w.Write(jsonBytes)
	return nil
}