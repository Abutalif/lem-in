package usecases

import (
	"fmt"
	"lem-in/internal/entities"
	"lem-in/internal/usecases/pathfinders"
)

type Pathfinder interface {
	Find(entities.Anthill) []entities.Path
}

type solveMachine struct {
	pathMakers []Pathfinder
	scheduler  Scheduler
}

type Solver interface {
	Solve(*entities.Anthill) entities.Queue
}

func NewSolver() Solver {
	pathMakers := make([]Pathfinder, 2)
	pathMakers[0] = pathfinders.NewSimple()
	pathMakers[1] = pathfinders.NewDikjstra()
	return &solveMachine{
		pathMakers: pathMakers,
		scheduler:  NewScheduler(),
	}
}

func (s *solveMachine) Solve(colony *entities.Anthill) entities.Queue {
	// biggestStep := entities.Infinity
	// for _, pm := range s.pathMakers {
	// 	paths := pm.Find(*colony)
	// 	order := s.scheduler.Schedule(paths, colony.AntNum, colony.GetStart())
	// 	// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
	// 	longestPath(order)
	// }
	return entities.Queue{}
}

func NewPathfinder(name string) Pathfinder {
	switch name {
	case "dijkstra":
		return pathfinders.NewDikjstra()
	default:
		return pathfinders.NewSimple()
	}
}

func Find() []entities.Path {
	// go through goroutines each running one pathfinder

	return nil
}

func CountSteps(paths []entities.Path, totalAnts int) int {
	pathsNum := len(paths)
	if pathsNum == 1 {
		return totalAnts - 1 + paths[0].Len
	}
	fmt.Println("paths num:", pathsNum)
	fmt.Println("Ants before ops:", totalAnts)
	sortPaths(paths)
	longestPathLen := paths[pathsNum-1].Len
	steps := 0
	for i := 0; i < pathsNum-1; i++ {
		fmt.Printf("path %v has len %v\n", i, paths[i].Len)
		diff := paths[i+1].Len - paths[i].Len
		steps += diff
		totalAnts -= longestPathLen - paths[i].Len
	}
	fmt.Println("longest path has len:", longestPathLen)

	fmt.Println("Ants after ops:", totalAnts)
	fmt.Println("Steps before:", steps)
	fmt.Println()
	steps += totalAnts/pathsNum + longestPathLen - 1
	return steps
}

func longestPath(paths []entities.Path) *entities.Path {
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

func sortPaths(paths []entities.Path) {
	n := len(paths)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if paths[j].Len > paths[j+1].Len {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}
}

func OldCountStep(paths []entities.Path, totalAnts int) int {
	fmt.Println("Num of paths:", len(paths))
	pathsNum := len(paths)
	offset := 0
	for _, p := range paths {
		offset += p.Len
	}
	biggest := longestPath(paths)
	return biggest.Len + (totalAnts-offset)/pathsNum + (totalAnts-offset)%pathsNum
}
