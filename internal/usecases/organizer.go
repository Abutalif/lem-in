package usecases

import (
	"lem-in/internal/entities"
)

type Organizer interface {
	Schedule([]entities.Path, int, *entities.Room) entities.Queue
}

type output struct{}

func NewOrganizer() Organizer {
	return &output{}
}

func (o *output) sortPaths(paths []entities.Path) {
	n := len(paths)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if paths[j].Len > paths[j+1].Len {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}
}

func (o *output) Schedule(paths []entities.Path, totalAnts int, start *entities.Room) entities.Queue {
	o.sortPaths(paths) // might be not needed
	queue := make(entities.Queue, 0)

	order := make([]*entities.Node, totalAnts)
	for i := range order {
		if len(paths) == 1 {
			order[i] = paths[0].Start
		} else {
			for j := 0; j < len(paths)-1; j++ {
				if paths[j].Len+paths[j].Ants > paths[j+1].Len+paths[j+1].Ants {
					order[i] = paths[j+1].Start
					paths[j+1].Ants++
				} else {
					order[i] = paths[j].Start
					paths[j].Ants++
				}
				// order[i].PrintList()
			}
		}
	}

	// FIXME: does not what was intendet
	var nilCounter int

	for {
		var usedPath []*entities.Node
		nilCounter = 0
		step := make(entities.Step, 0)
	Mid:
		for i := 0; i < totalAnts; i++ {
			if order[i] == nil {
				continue
			}
			for _, used := range usedPath {
				if used == order[i] {
					break Mid
				}
			}
			move := entities.Move{
				Ant:         i + 1,
				Destination: order[i].Current.Name,
			}
			usedPath = append(usedPath, order[i])
			order[i] = order[i].Next
			step = append(step, move)
		}

		queue = append(queue, step)
		for _, j := range order {
			if j == nil {
				nilCounter++
			}
		}
		if nilCounter == totalAnts {
			break
		}
	}

	return queue
}
