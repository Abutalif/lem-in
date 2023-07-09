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

// FIXME: unoptimised func
func (o *output) runAnts(order []*entities.Node, totalAnts int) entities.Queue {
	queue := make(entities.Queue, 0)
	nilCounter := 0
	for nilCounter != totalAnts {
		usedNodes := make(map[*entities.Node]bool)
		nilCounter = 0
		step := make(entities.Step, 0)
		for i := 0; i < totalAnts; i++ {
			if order[i] == nil {
				nilCounter++
				continue
			}
			if usedNodes[order[i]] {
				break
			}
			move := entities.Move{
				Ant:         i + 1,
				Destination: order[i].Current.Name,
			}
			usedNodes[order[i]] = true
			order[i] = order[i].Next
			step = append(step, move)
		}
		if len(step) != 0 {
			queue = append(queue, step)
		}
	}
	// return entities.Queue{}
	return queue
}
