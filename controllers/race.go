package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/CoreBuilder/go-starcraftservice/models"
)

type raceController struct {
	raceIDPattern *regexp.Regexp
}

func (uc raceController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/races" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.raceIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *raceController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetRaces(), w)
}

func (uc *raceController) get(id int, w http.ResponseWriter) {
	u, err := models.GetRaceByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *raceController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Race object"))
		return
	}
	u, err = models.AddRace(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *raceController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Race object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted race must match ID in URL"))
		return
	}
	u, err = models.UpdateRace(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *raceController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveRaceByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *raceController) parseRequest(r *http.Request) (models.Race, error) {
	dec := json.NewDecoder(r.Body)
	var u models.Race
	err := dec.Decode(&u)
	if err != nil {
		return models.Race{}, err
	}
	return u, nil
}

func newUserController() *raceController {
	return &raceController{
		raceIDPattern: regexp.MustCompile(`^/races/(\d+)/?`),
	}
}
