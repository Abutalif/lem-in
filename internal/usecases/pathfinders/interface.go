package pathfinders

import "lem-in/internal/entities"

type Pathfinder interface {
	Find(entities.Anthill) []*entities.Path
}

func SortPaths(paths []*entities.Path) {
	n := len(paths)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if paths[j].Len > paths[j+1].Len {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}
}
