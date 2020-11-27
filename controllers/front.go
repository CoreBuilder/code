package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers ...
func RegisterControllers() {
	rc := newUserController()

	http.Handle("/race", *rc)
	http.Handle("/races/", *rc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
