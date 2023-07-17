package entities

import "sort"

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
	Connections map[*Room]uint
}

func (r *Room) IsNeighbor(room *Room) bool {
	_, ok := r.Connections[room]
	return ok
}

// I wanted to remove sort function
// func (r *Room) SortConnByDist() {
// 	length := len(r.Connections)
// 	if length <= 1 {
// 		return
// 	}

// 	for i := 0; i < length-1; i++ {
// 		for j := 0; j < length-i-1; j++ {
// 			if r.Connections[j].StartDist > r.Connections[j+1].StartDist {
// 				r.Connections[j], r.Connections[j+1] = r.Connections[j+1], r.Connections[j]
// 			}
// 		}
// 	}
// }

func (r *Room) SortConnByDist() []*Room {
	sortedRooms := make([]*Room, 0, len(r.Connections))
	for room := range r.Connections {
		sortedRooms = append(sortedRooms, room)
	}
	sort.Slice(sortedRooms, func(i, j int) bool {
		return sortedRooms[i].StartDist < sortedRooms[j].StartDist
	})
	return sortedRooms
}

func (r *Room) DeleteNeighbor(neighbor *Room) {
	delete(r.Connections, neighbor)
}
