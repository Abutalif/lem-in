package usecases

import (
	"errors"

	"lem-in/internal/entities"
)

type Builder struct {
	anthill *entities.Anthill
}

func NewBuilder() Builder {
	return Builder{
		anthill: &entities.Anthill{
			Rooms: make(map[string]*entities.Room),
		},
	}
}

func (b *Builder) Anthill() *entities.Anthill {
	return b.anthill
}

func (b *Builder) SetAnts(num uint) {
	b.anthill.AntNum = num
}

func (b *Builder) CreateRoom(name string, x, y int, kind entities.RoomKind) error {
	if _, has := b.anthill.Rooms[name]; has {
		return errors.New("ERROR: repeated rooms")
	}
	newRoom := &entities.Room{
		Name:        name,
		Visited:     false,
		X:           x,
		Y:           y,
		Kind:        kind,
		Connections: make([]*entities.Room, 0),
	}
	b.anthill.Rooms[name] = newRoom
	return nil
}

func (b *Builder) CreateTunnel(roomNames []string) error {
	room1, has1 := b.anthill.Rooms[roomNames[0]]
	room2, has2 := b.anthill.Rooms[roomNames[1]]
	if !has1 || !has2 {
		return errors.New("ERROR: invalid tunnel info - tunnel to nonexisting room")
	}

	var existing bool
	if len(room1.Connections) > len(room2.Connections) {
		existing = room2.IsNeighbor(room1)
	} else {
		existing = room1.IsNeighbor(room2)
	}

	if existing {
		return nil
	}

	room1.Connections = append(room1.Connections, room2)
	room2.Connections = append(room2.Connections, room1)

	return nil
}
