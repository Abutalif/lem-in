package pathfinders

import (
	"lem-in/internal/entities"
)

type dijkstra struct{}

func NewDikjstra() Pathfinder {
	return &dijkstra{}
}

func (d *dijkstra) Find(colony *entities.Anthill) []*entities.Path {
	paths := make([]*entities.Path, 0)
	start := colony.GetStart()
	start.StartDist = 0
	d.setDistances(start)

	end := colony.GetEnd()
	end.Visited = true
	sortedNeighbors := end.SortConnByDist()
	for _, neighbor := range sortedNeighbors {
		route := d.getRoute(neighbor)
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

		paths = append(paths, &path)
	}

	return paths
}

func (d *dijkstra) setDistances(room *entities.Room) {
	for neighbor, cost := range room.Connections {
		if neighbor.StartDist > room.StartDist+cost {
			neighbor.StartDist = room.StartDist + cost
			d.setDistances(neighbor)
		}
	}
}

func (d *dijkstra) getRoute(current *entities.Room) *entities.Node {
	if current.StartDist == 0 {
		return &entities.Node{
			Current: current,
			Next:    nil,
		}
	}
	// current.SortConnByDist()
	current.Visited = true
	sortedNeighbors := current.SortConnByDist()
	for _, neighbor := range sortedNeighbors {
		if neighbor.Visited || neighbor.StartDist > current.StartDist {
			continue
		}
		route := d.getRoute(neighbor)
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
