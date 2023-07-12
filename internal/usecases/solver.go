package usecases

import (
	"errors"

	"lem-in/internal/entities"
	"lem-in/internal/usecases/pathfinders"
)

type Pathfinder interface {
	Find(entities.Anthill) []entities.Path
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
	return &solveMachine{
		pathMakers: pathMakers,
		scheduler:  NewScheduler(),
	}
}

func (s *solveMachine) Solve(colony *entities.Anthill) (entities.Queue, error) {
	paths := s.pathMakers["dijkstra"].Find(*colony)
	if len(paths) < 1 {
		return entities.Queue{}, errors.New("no path found")
	}
	queue := s.scheduler.Schedule(paths, colony.AntNum)

	// the idea is simple
	// we are doing only simple single threaded approach
	// no goroutines
	// I just ask pathfinder to give me paths,
	// then I ask scheduler to schedule ants it will say wich ant where will move
	// Then RunAnts will compose a queue of steps
	// which will be returned by this func
	return queue, nil
}

// this function should count number of steps, which we will use to create array (not slice)
// that will be our queue
func CountSteps(paths []entities.Path, totalAnts int) int {
	pathsNum := len(paths)
	if pathsNum == 1 {
		return totalAnts - 1 + paths[0].Len
	}
	SortPaths(paths)
	longestPathLen := paths[pathsNum-1].Len
	steps := 0
	for i := 0; i < pathsNum-1; i++ {
		diff := paths[i+1].Len - paths[i].Len
		steps += diff
		totalAnts -= longestPathLen - paths[i].Len
	}
	steps += totalAnts/pathsNum + longestPathLen - 1
	return steps
}

func LongestPath(paths []entities.Path) *entities.Path {
	if len(paths) == 0 {
		return nil
	}
	biggest := &paths[0]
	for i := range paths {
		if paths[i].Len > biggest.Len {
			biggest = &paths[i]
		}
	}
	return biggest
}

// ideally i dont want to sort paths.
// this is wasteful
func SortPaths(paths []entities.Path) {
	n := len(paths)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if paths[j].Len > paths[j+1].Len {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}
}
