package entities

import (
	"strconv"
)

type Anthill struct {
	Rooms  map[string]*Room
	AntNum int
}

type Room struct {
	Name        string
	X           int
	Y           int
	Visited     bool
	Kind        RoomKind
	StartDist   uint
	Connections []*Room
}

type RoomKind uint8

const (
	Unknown RoomKind = iota
	Regular
	Start
	End
)

const Infinity = ^uint(0)

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

func (a *Anthill) Show() string {
	res := ""
	res += "AntNum: " + strconv.Itoa(a.AntNum) + "\n"
	res += "Rooms:\n"
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
		res += val.Name + " - " + kind + " - startDist:" + strconv.Itoa(int(val.StartDist)) + "\n"
		res += "Connections:\n"
		for i, cons := range val.Connections {
			res += cons.Name
			if i < len(val.Connections)-1 {
				res += ", "
			}

		}
		res += "\n\n"
	}
	return res
}

func (r *Room) IsNeighbor(room *Room) bool {
	for _, neighbor := range r.Connections {
		if neighbor == room {
			return true
		}
	}

	return false
}

func (r *Room) SortConnByDist() {
	length := len(r.Connections)
	if length <= 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if r.Connections[j].StartDist > r.Connections[j+1].StartDist {
				r.Connections[j], r.Connections[j+1] = r.Connections[j+1], r.Connections[j]
			}
		}
	}
}
