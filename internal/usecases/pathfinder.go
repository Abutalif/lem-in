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
	// paths []entities.Node
	end *entities.Room
}

func NewPathfinder() Pathfinder {
	return &simple{} //paths: make([]entities.Node, 0)
}

// finishing
func (s *simple) Find(colony *entities.Anthill) []*entities.Node {
	start := colony.GetStart()
	start.Visited = true
	s.end = colony.GetEnd()
	paths := make([]*entities.Node, 0)
	for _, afterStart := range start.Connections {
		if !afterStart.Visited {
			path := s.checkNeighbors(afterStart)
			if path != nil {
				last := path.GetLast()
				last.Next = &entities.Node{
					Current: afterStart,
					Next: &entities.Node{
						Current: start,
						Next:    nil,
					},
				}
				paths = append(paths, path)
			}
		}
	}
	return paths
}

func (s *simple) checkNeighbors(current *entities.Room) *entities.Node {
	current.Visited = true
	// fmt.Printf("visiting: %v\n", current.Name)
	for _, neighbor := range current.Connections {
		if !neighbor.Visited {
			if neighbor.Kind == entities.End {
				return &entities.Node{
					Current: s.end,
					Next:    nil,
				}
			}
			if neighbor.Kind == entities.Regular {
				route := s.checkNeighbors(neighbor)
				if route == nil {
					return nil
				} else {
					last := route.GetLast()
					last.Next = &entities.Node{
						Current: neighbor,
						Next:    nil,
					}
					return route
				}
			}
		}
	}
	return nil
}
