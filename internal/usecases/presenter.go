package usecases

import "lem-in/internal/entities"

type Presenter interface{}

type output struct{}

func NewPresenter(out output) Presenter {
	return []output{} // I don't like it. Should be more conteinerized.
}

func (o *output) MovementQueue(routes []entities.Path) {
}
