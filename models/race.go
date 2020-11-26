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

func GetRaces() []*Race {
	return races
}

func AddRace(r Race) (Race, error) {
	r.ID = nextID
	nextID++
	races = append(races, &r)
	return r, nil
}
