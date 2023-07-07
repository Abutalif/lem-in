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

func (o *output) Schedule(paths []entities.Path, totalAnts int, start *entities.Room) entities.Queue {
	// for i := range paths {
	// 	fmt.Println(paths[i].Start.PrintList())
	// }
	order := make([]*entities.Node, totalAnts)
	smallest := &paths[0]
	for i := len(order) - 1; i >= 0; i-- {
		for j := range paths {
			if smallest.Ants+smallest.Len > paths[j].Len+paths[j].Ants && smallest != &paths[j] {
				smallest = &paths[j]
			}
		}
		// smallest = o.leastBusyPath(paths)
		order[i] = smallest.Start.Next
		smallest.Ants++
	}
	queue := o.runAnts(order, totalAnts)

	return queue
}

func (o *output) runAnts(order []*entities.Node, totalAnts int) entities.Queue {
	queue := make(entities.Queue, 0)

	for {
		var usedPath []*entities.Node
		nilCounter := 0
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
