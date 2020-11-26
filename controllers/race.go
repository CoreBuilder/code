package controllers

import (
	"net/http"
	"regexp"
)

type raceController struct {
	raceIDPattern *regexp.Regexp
}

func (rc raceController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("We are the protoss. Children of ancient gods!"))
}

func newRaceController() *raceController {
	return &raceController{
		raceIDPattern: regexp.MustCompile(`^/races/(\d+)/?`),
	}
}
