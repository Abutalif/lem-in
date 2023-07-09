package usecases

import (
	"lem-in/internal/entities"
	"lem-in/internal/usecases/pathfinders"
)

type Pathfinder interface {
	Find(*entities.Anthill) []entities.Path
}
type PathMaker interface {
	Find(*entities.Anthill) []entities.Path
}

func NewPathfinder(name string) Pathfinder {
	switch name {
	case "dijkstra":
		return pathfinders.NewDikjstra()
	default:
		return pathfinders.NewSimple()
	}
}

func FindPath() []entities.Path {
	// go through goroutines each running one pathfinder
	return nil
}
