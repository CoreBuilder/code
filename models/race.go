package models

import (
	"errors"
	"fmt"
)

// Race is a playable specie in starcraft
type Race struct {
	ID       int
	Name     string
	Religion string
}

var (
	races  []*Race
	nextID = 1
)

// GetRaces ...
func GetRaces() []*Race {
	return races
}

// AddRace ...
func AddRace(r Race) (Race, error) {
	if r.ID != 0 {
		return Race{}, errors.New("the new race cannot have an id")
	}
	r.ID = nextID
	nextID++
	races = append(races, &r)
	return r, nil
}

// GetRaceByID ...
func GetRaceByID(id int) (Race, error) {
	for _, v := range races {
		if v.ID == id {
			return *v, nil
		}
	}
	return Race{}, fmt.Errorf("Race with ID '%v' not found", id)
}

// UpdateRace ...
func UpdateRace(r Race) (Race, error) {
	for i, candidate := range races {
		if candidate.ID == r.ID {
			races[i] = &r
			return r, nil
		}
	}
	return Race{}, fmt.Errorf("Race with ID '%v' not found", r.ID)
}

// RemoveRaceByID ...
func RemoveRaceByID(id int) error {
	for i, r := range races {
		if r.ID == id {
			races = append(races[:i], races[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Race with ID '%v' not found", id)
}
