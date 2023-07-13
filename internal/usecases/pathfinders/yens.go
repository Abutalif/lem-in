package pathfinders

import (
	"fmt"
	"lem-in/internal/entities"
)

type yens struct {
	paths    []*entities.Path
	dijkstra Pathfinder
}

func NewYens() Pathfinder {
	return &yens{
		paths:    make([]*entities.Path, 0),
		dijkstra: NewDikjstra(),
	}
}

func (y *yens) Find(colony entities.Anthill) []*entities.Path {
	initialPaths := y.dijkstra.Find(colony)
	fmt.Println("initial paths")
	for _, p := range initialPaths {
		fmt.Println(p.Start.PrintList())
	}
	fmt.Println()
	for _, p := range initialPaths {
		start := p.Start // first step after start
		var updatedPaths []*entities.Path
		j := 0
		for start.Next != nil {
			thisRoom := start.Current
			nextRoom := start.Next.Current
			thisRoom.DeleteNeighbor(nextRoom)
			nextRoom.DeleteNeighbor(thisRoom)
			colony.UnvisitAll()
			updatedPaths = y.dijkstra.Find(colony)

			fmt.Println("Path update", j)
			for _, updated := range updatedPaths {
				fmt.Println(updated.Start.PrintList())
				found := false
				for _, initial := range initialPaths {
					if initial == updated {
						found = true
						break
					}
				}
				for _, existing := range y.paths {
					if existing == updated {
						found = true
						break
					}
				}
				if !found {
					y.paths = append(y.paths, updated)
				}
				fmt.Println()
			}
			thisRoom.Connections = append(thisRoom.Connections, nextRoom)
			nextRoom.Connections = append(nextRoom.Connections, thisRoom)
			start = start.Next
			j++
		}
	}
	return y.paths
}

// func (y *yens) GetPaths() []entities.Path {
// 	res := make([]entities.Path, 0)
// 	for i, p := range y.paths {
// 		if p {
// 			res = append(res, i)
// 		}
// 	}
// 	return res
// }
