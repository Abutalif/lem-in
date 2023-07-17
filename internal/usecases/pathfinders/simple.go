package pathfinders

import "lem-in/internal/entities"

type simple struct {
	end *entities.Room
}

func NewSimple() Pathfinder {
	return &simple{}
}

func (s *simple) Find(colony *entities.Anthill) []*entities.Path {
	start := colony.GetStart()
	s.end = colony.GetEnd()
	start.Visited = true
	paths := make([]*entities.Path, 0)
	for afterStart := range start.Connections {
		if !afterStart.Visited {
			route := s.checkNeighbors(afterStart)
			if route != nil {
				len := route.Len()
				route = route.Reverse()
				route = route.ChangeFirst(start)
				path := entities.Path{
					Start: route,
					Len:   len,
					Ants:  0,
				}
				paths = append(paths, &path)
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
	for neighbor := range current.Connections {
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
