package pathfinders

import "lem-in/internal/entities"

type Pathfinder interface {
	Find(entities.Anthill) []entities.Path
}
