package pathfinders

import "lem-in/internal/entities"

type bfs struct{}

func NewBFS() Pathfinder {
	return &bfs{}
}

func (b *bfs) Find(colony *entities.Anthill) []entities.Path {
	return nil
}
