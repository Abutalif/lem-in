package usecases

import (
	"lem-in/internal/entities"
	pathfinder "lem-in/internal/usecases/pathfinders"
)

type Pathfinder interface {
	Find(*entities.Anthill) []*entities.Node
}

func NewPathfinder(name string) Pathfinder {
	switch name {
	default:
		return pathfinder.NewSimple()
	}
}

// func GetPathfinders() []string {
// 	// return list of names of pathfinders
// 	return nil
// }
