package usecases

import "lem-in/internal/entities"

type Builder struct {
	anthill *entities.Anthill
}

func NewBuilder() Builder {
	return Builder{anthill: &entities.Anthill{}}
}

func (b *Builder) SetAnts(num uint) {
	b.anthill.AntNum = num
}

// TODO
func (b *Builder) CreateRoom(line string, kind entities.Kind) error {
	return nil
}

// roomRgx := regexp.MustCompile("")
// tunnelRgx := regexp.MustCompile("")
// commentRgx := regexp.MustCompile("")
