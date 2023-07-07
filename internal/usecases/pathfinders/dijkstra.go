package pathfinders

import (
	"lem-in/internal/entities"
)

type dijkstra struct{}

func NewDikjstra() Pathfinder {
	return &dijkstra{}
}

func (d *dijkstra) Find(colony *entities.Anthill) []entities.Path {
	paths := make([]entities.Path, 0)
	start := colony.GetStart()
	start.StartDist = 0
	d.setDistances(start)

	end := colony.GetEnd()
	end.SortConnByDist()
	for _, neighbor := range end.Connections {
		route := getRoute(neighbor)
		if route == nil {
			continue
		}
		len := route.Len()
		route = route.ChangeFirst(end)
		path := entities.Path{
			Start: route.Reverse(),
			Len:   len,
			Ants:  0,
		}

		paths = append(paths, path)

	}

	return paths
}

func (d *dijkstra) setDistances(room *entities.Room) {
	for _, neighbor := range room.Connections {
		if neighbor.StartDist > room.StartDist {
			neighbor.StartDist = room.StartDist + 1
			d.setDistances(neighbor)
		}
	}
}

func getRoute(current *entities.Room) *entities.Node {
	if current.StartDist == 0 {
		return &entities.Node{
			Current: current,
			Next:    nil,
		}
	}
	current.SortConnByDist()
	current.Visited = true
	for _, neighbor := range current.Connections {
		if neighbor.Visited || neighbor.StartDist > current.StartDist {
			continue
		}
		route := getRoute(neighbor)
		if route == nil {
			continue
		}
		return &entities.Node{
			Current: current,
			Next:    route,
		}
	}
	return nil
}
