package entities

import "fmt"

type Anthill struct {
	Rooms  map[string]*Room
	AntNum uint
}

type Room struct {
	Name        string
	X           int
	Y           int
	Visited     bool
	Kind        RoomKind
	Connections []*Room
}

type RoomKind uint8

const (
	Unknown RoomKind = iota
	Regular
	Start
	End
)

func (a *Anthill) GetStart() *Room {
	for _, room := range a.Rooms {
		if room.Kind == Start {
			return room
		}
	}
	return nil
}

func (a *Anthill) GetEnd() *Room {
	for _, room := range a.Rooms {
		if room.Kind == End {
			return room
		}
	}
	return nil
}

func (a *Anthill) Show() {
	fmt.Println("AntNum:", a.AntNum)
	fmt.Println("Rooms:")
	for _, val := range a.Rooms {
		var kind string
		switch val.Kind {
		case Start:
			kind = "start"
		case Regular:
			kind = "regular"
		case End:
			kind = "end"
		}
		fmt.Printf("%v - %v - visited: %v, connected to:\n", val.Name, kind, val.Visited)
		for _, cons := range val.Connections {
			fmt.Printf("%v, ", cons.Name)
		}
		fmt.Printf("\n\n")
	}
}

func (r *Room) IsNeighbor(room *Room) bool {
	for _, neighbor := range r.Connections {
		if neighbor == room {
			return true
		}
	}

	return false
}
