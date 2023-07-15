package usecases

import (
	"errors"
	"fmt"

	"lem-in/internal/entities"
	"lem-in/internal/usecases/pathfinders"
)

type Pathfinder interface {
	Find(entities.Anthill) []*entities.Path
}

type solveMachine struct {
	pathMakers map[string]Pathfinder
	scheduler  Scheduler
}

type Solver interface {
	Solve(*entities.Anthill) (entities.Queue, error)
}

func NewSolver() Solver {
	pathMakers := make(map[string]Pathfinder)
	pathMakers["simple"] = pathfinders.NewSimple()
	pathMakers["dijkstra"] = pathfinders.NewDikjstra()
	pathMakers["augDijkstra"] = pathfinders.NewAugDikjstra()
	return &solveMachine{
		pathMakers: pathMakers,
		scheduler:  NewScheduler(),
	}
}

func (s *solveMachine) Solve(colony *entities.Anthill) (entities.Queue, error) {
	paths := s.pathMakers["augDijkstra"].Find(*colony)
	fmt.Println("Used paths")
	for _, p := range paths {
		fmt.Println(p.Start.PrintList())
	}
	// fmt.Println("\n", colony.Show())
	if len(paths) < 1 {
		return entities.Queue{}, errors.New("no path found")
	}
	queue := s.scheduler.Schedule(paths, colony.AntNum)
	return queue, nil
}

// this function should count number of steps, which we will use to create array (not slice)
// that will be our queue
// func CountSteps(paths []entities.Path, totalAnts int) int {
// 	pathsNum := len(paths)
// 	if pathsNum == 1 {
// 		return totalAnts - 1 + paths[0].Len
// 	}
// 	SortPaths(paths)
// 	longestPathLen := paths[pathsNum-1].Len
// 	steps := 0
// 	for i := 0; i < pathsNum-1; i++ {
// 		diff := paths[i+1].Len - paths[i].Len
// 		steps += diff
// 		totalAnts -= longestPathLen - paths[i].Len
// 	}
// 	steps += totalAnts/pathsNum + longestPathLen - 1
// 	return steps
// }

// func LongestPath(paths []entities.Path) *entities.Path {
// 	if len(paths) == 0 {
// 		return nil
// 	}
// 	biggest := &paths[0]
// 	for i := range paths {
// 		if paths[i].Len > biggest.Len {
// 			biggest = &paths[i]
// 		}
// 	}
// 	return biggest
// }
