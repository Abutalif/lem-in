package entities

import (
	"strconv"
)

type Anthill struct {
	Rooms  map[string]*Room
	AntNum int
}

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

func (a *Anthill) UnvisitAll() {
	for _, v := range a.Rooms {
		v.Visited = false
	}
}

func (a *Anthill) Show() string {
	res := ""
	res += "AntNum: " + strconv.Itoa(a.AntNum) + "\n"
	res += "Rooms:\n"
	for _, val := range a.Rooms {
		var kind string
		var visited string
		if val.Visited {
			visited = " - was visted"
		} else {
			visited = " - was not visited"
		}
		switch val.Kind {
		case Start:
			kind = "start"
		case Regular:
			kind = "regular"
		case End:
			kind = "end"
		}
		res += val.Name + " - " + kind + " - startDist:" + strconv.Itoa(int(val.StartDist)) + visited + "\n"
		res += "Connections:\n"
		i := 0
		for room := range val.Connections {
			res += room.Name
			if i < len(val.Connections) {
				res += ", "
			}
			i++
		}
		res += "\n\n"
	}
	return res
}
