package usecases

import (
	"lem-in/internal/entities"
	pathfinder "lem-in/internal/usecases/pathfinders"
)

type Pathfinder interface {
	Find(*entities.Anthill) []entities.Path
}

func NewPathfinder(name string) Pathfinder {
	switch name {
	default:
		return pathfinder.NewSimple()
	}
}

// TODO
func GetPathfinders() []Pathfinder { // or can return []string
	// return list of avialable of pathfinders
	return nil
}
