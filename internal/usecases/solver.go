package usecases

import (
	"errors"
	"fmt"

	"lem-in/internal/entities"
	"lem-in/internal/usecases/pathfinders"
)

type Pathfinder interface {
	Find(*entities.Anthill) []*entities.Path
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
	// fmt.Println("started pathfinding")
	paths := s.pathMakers["dijkstra"].Find(colony)
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
