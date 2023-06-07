package usecases

import (
	"lem-in/internal/entities"
)

// Pathfinder is an interface for objects that can find paths from
// start room to end room
type Pathfinder interface {
	Find(*entities.Anthill) []*entities.Node
}

type simple struct {
	end *entities.Room
}

func NewPathfinder() Pathfinder {
	return &simple{}
}

func (s *simple) Find(colony *entities.Anthill) []*entities.Node {
	start := colony.GetStart()
	s.end = colony.GetEnd()
	start.Visited = true
	paths := make([]*entities.Node, 0)
	for _, afterStart := range start.Connections {
		if !afterStart.Visited {
			path := s.checkNeighbors(afterStart)
			if path != nil {
				last := path.GetLast()
				last.Next = &entities.Node{
					Current: start,
					Next:    nil,
				}

				paths = append(paths, path)
			}
		}
	}
	return paths
}

func (s *simple) checkNeighbors(current *entities.Room) *entities.Node {
	if current.Kind == entities.End {
		return &entities.Node{
			Current: s.end,
			Next:    nil,
		}
	}
	current.Visited = true
	for _, neighbor := range current.Connections {
		if !neighbor.Visited {
			route := s.checkNeighbors(neighbor)
			if route == nil {
				return nil
			} else {
				last := route.GetLast()
				last.Next = &entities.Node{
					Current: current,
					Next:    nil,
				}
				return route
			}
		}
	}
	return nil
}
