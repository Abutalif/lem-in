package entities

const (
	Unknown RoomKind = iota
	Regular
	Start
	End
)

const Infinity = ^uint(0)

type RoomKind uint8

type Room struct {
	Name        string
	X           int
	Y           int
	Visited     bool
	Kind        RoomKind
	StartDist   uint
	Connections []*Room
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

func (r *Room) DeleteNeighbor(neighbor *Room) {
	for i, room := range r.Connections {
		if room == neighbor {
			r.Connections = append(r.Connections[:i], r.Connections[i+1:]...)
		}
	}

}
