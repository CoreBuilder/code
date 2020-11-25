package models

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
