package usecases_lol

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
	// queue := o.runAnts(order, len(paths))

	return queue
}

func (o *output) runAnts(order []*entities.Node, totalAnts int) entities.Queue {
	queue := make(entities.Queue, 0)
	// repetition := 0
	var nilCounter int
	for nilCounter != totalAnts {
		usedNode := make(map[*entities.Node]bool)
		nilCounter = totalAnts
		step := make(entities.Step, 0)
	Mid:
		for i := 0; i < totalAnts; i++ {
			if order[i] == nil {
				continue
			}
			if usedNode[order[i]] {
				break Mid
			}
			move := entities.Move{
				Ant:         i + 1,
				Destination: order[i].Current.Name,
			}
			usedNode[order[i]] = true
			order[i] = order[i].Next
			step = append(step, move)
			nilCounter--
			// repetition++
		}
		queue = append(queue, step)
	}
	// fmt.Println("queue length (outer loop):", len(queue))
	// fmt.Println("repetitions:", repetition)
	return queue
}

// func (o *output) runAntsGOroutine(order []*entities.Node, pathsNum int) entities.Queue {
// 	totalAnts := len(order)
// 	queue := make(entities.Queue, totalAnts)
// 	repetition := 0
// 	for i := 0; i < totalAnts; i++ {
// 		go func(i int) {
// 			// steps := make(entities.Step, order[i].Len())
// 			// steps := make(entities.Step, pathsNum)
// 			steps := make(entities.Step, 0)
// 			j := 0
// 			for order[i] != nil {
// 				move := entities.Move{
// 					Ant:         i + 1,
// 					Destination: order[i].Current.Name,
// 				}
// 				order[i] = order[i].Next
// 				// steps[j] = move
// 				steps = append(steps, move)
// 				j++
// 				repetition++
// 			}
// 			if i < len(queue) {
// 				queue[i] = steps
// 			} else {
// 				queue = append(queue, steps)
// 			}
// 		}(i)
// 	}
// 	// fmt.Println("queue length (outer loop):", len(queue))
// 	// fmt.Println("repetitions:", repetition)
// 	return queue
// }

// func (o *output) oldRunAnts(order []*entities.Node, totalAnts int) entities.Queue {
// 	queue := make(entities.Queue, 0)

// 	for {
// 		var usedPath []*entities.Node
// 		nilCounter := 0
// 		step := make(entities.Step, 0)
// 	Mid:
// 		for i := 0; i < totalAnts; i++ {
// 			if order[i] == nil {
// 				continue
// 			}
// 			for _, used := range usedPath {
// 				if used == order[i] {
// 					break Mid
// 				}
// 			}
// 			move := entities.Move{
// 				Ant:         i + 1,
// 				Destination: order[i].Current.Name,
// 			}
// 			usedPath = append(usedPath, order[i])
// 			order[i] = order[i].Next
// 			step = append(step, move)
// 		}

// 		queue = append(queue, step)
// 		for _, j := range order {
// 			if j == nil {
// 				nilCounter++
// 			}
// 		}
// 		if nilCounter == totalAnts {
// 			break
// 		}
// 	}
// 	return queue
// }
