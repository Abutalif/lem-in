package pathfinders

import (
	"lem-in/internal/entities"
)

type augmentedDijkstra struct {
	d dijkstra
}

func NewAugDikjstra() Pathfinder {
	return &augmentedDijkstra{
		d: dijkstra{},
	}
}

func (a *augmentedDijkstra) Find(colony entities.Anthill) []*entities.Path {
	paths := a.d.Find(colony)
	if len(paths) == a.MaxPaths(*colony.GetStart(), *colony.GetEnd()) {
		return paths
	}

	var newPaths []*entities.Path
	node := paths[0].Start.Reverse()
	colony.UnvisitAll()
	for node.Next != nil {
		currentRoom := node.Current
		nextRoom := node.Next.Current
		for _, room := range currentRoom.Connections {
			if room == nextRoom {
				continue
			}
			route := a.d.getRoute(room)
			if route == nil {
				continue
			}
			len := route.Len()
			// route = route.ChangeFirst(end) i have to append only nodes that came before before
			path := entities.Path{
				Start: route.Reverse(),
				Len:   len,
				Ants:  0,
			}

			paths = append(paths, &path)
		}
		node = node.Next

	}

	// checking condintion
	// if len(colony.GetEnd().Connections) == len(paths) {
	// 	return paths
	// }
	// var newPaths []*entities.Path
	// for i := range paths {
	// 	fmt.Println("\njudjing paths")
	// 	colony.UnvisitAll()
	// 	// unvisit all paths
	// 	// break all connections from a paths
	// 	// go dijkstra
	// 	// if found better, return
	// 	route := paths[i].Start
	// 	broken := make(map[*entities.Room]*entities.Room)
	// 	for route.Next != nil {
	// 		fmt.Println("cheking route:room:", route.Current.Name)
	// 		this := route.Current
	// 		next := route.Next.Current
	// 		broken[this] = next
	// 		this.DeleteNeighbor(next)
	// 		next.DeleteNeighbor(this)
	// 		route = route.Next
	// 	}
	// 	// newPaths := a.dijkstra.Find(colony)
	// 	// return a.betterPaths(paths, newPaths)
	// 	// for k, v := range broken {
	// 	// 	k.Connections = append(k.Connections, v)
	// 	// 	v.Connections = append(v.Connections, k)
	// 	// 	delete(broken, k)
	// 	// }
	// }

	return newPaths
}

func (a *augmentedDijkstra) MaxPaths(start, end entities.Room) int {
	startConns := len(start.Connections)
	endConns := len(end.Connections)
	if startConns < endConns {
		return startConns
	}
	return endConns
}
