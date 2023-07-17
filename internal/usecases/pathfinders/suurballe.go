package pathfinders

import "lem-in/internal/entities"

type suurballe struct {
	d dijkstra
}

func NewSuurbale() Pathfinder {
	return &suurballe{
		d: dijkstra{},
	}
}

func (s *suurballe) Find(colony *entities.Anthill) []*entities.Path {
	initialPaths := s.d.Find(colony)
	// sortPaths(initialPaths)
	// extend for several paths.
	// if len(initialPaths) == maxPaths(colony.GetStart(), colony.GetEnd()) {
	// 	return initialPaths
	// }
	if len(initialPaths) != 1 {
		return initialPaths
	}
	// path1 := initialPaths[0]
	// TODO modify weights and run dijkstra again.

	return nil
}
