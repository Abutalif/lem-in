package usecases

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

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

func (b *Builder) SetAnts(num uint) {
	b.anthill.AntNum = num
}

func (b *Builder) CreateRoom(line string, kind entities.RoomKind) error {
	name, x, y, err := b.checkData(line)
	if err != nil {
		return err
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

func (b *Builder) checkData(line string) (string, int, int, error) {
	roomData := strings.Split(line, " ")
	if len(roomData) != 3 {
		return "", 0, 0, errors.New("ERROR: invalid room format - wrong number of entires")
	}

	x, err := strconv.Atoi(roomData[1])
	if err != nil {
		return "", 0, 0, errors.New("ERROR: invalid room format - incorrect x coord")
	}

	y, err := strconv.Atoi(roomData[1])
	if err != nil {
		return "", 0, 0, errors.New("ERROR: invalid room format - incorrect y coord")
	}

	return roomData[0], x, y, nil
}

func (b *Builder) ShowAnthill() {
	fmt.Println("AntNum:", b.anthill.AntNum)
	fmt.Println("Rooms:")
	for _, val := range b.anthill.Rooms {
		var kind string
		switch val.Kind {
		case entities.Start:
			kind = "start"
		case entities.Regular:
			kind = "regular"
		case entities.End:
			kind = "end"
		}
		fmt.Printf("%v - %v, connected to:\n", val.Name, kind)
		for _, cons := range val.Connections {
			fmt.Printf("%v, ", cons.Name)
		}
		fmt.Printf("\n\n")
	}
}

func (b *Builder) CreateTunnel(line string) error {
	rooms := strings.Split(line, "-")
	if len(rooms) != 2 {
		return errors.New("ERROR: incorrect tunnel info - not 2 rooms")
	}
	room1, has1 := b.anthill.Rooms[rooms[0]]
	room2, has2 := b.anthill.Rooms[rooms[1]]
	if !has1 || !has2 {
		return errors.New("ERROR: invalid tunnel info - tunnel to nonexisting room")
	}
	// warning: biderectional tunnel might have error when ants will move from neighboring rooms
	room1.Connections = append(room1.Connections, room2)
	room2.Connections = append(room2.Connections, room1)

	return nil
}
