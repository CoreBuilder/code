package controllers

import "net/http"

// RegisterControllers ...
func RegisterControllers() {
	rc := newRaceController()

	http.Handle("/race", *rc)
	http.Handle("/races/", *rc)
}
