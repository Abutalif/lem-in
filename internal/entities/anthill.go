package entities

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
