package usecases

import "lem-in/internal/entities"

type Builder struct {
	anthill entities.Anthill
}

func NewBuilder() Builder {
	return Builder{anthill: entities.Anthill{}}
}
