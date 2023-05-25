package usecases

import (
	"lem-in/internal/entities"
)

type Solver interface {
	Solve(*entities.Anthill) entities.Solution
}

type naive struct {
	end *entities.Room
}

func NewNaiveSolver() Solver {
	return &naive{}
}

// ****HERE****
func (n *naive) Solve(anthill *entities.Anthill) entities.Solution {
	solution := make(entities.Solution, 0)
	start := anthill.GetStart()
	n.end = anthill.GetEnd()
	for _, conn := range start.Connections {
		path := &entities.Path{
			Current: conn,
		}
		n.getPath(path)
		solution = append(solution, path)
	}
	return solution
}

func (n *naive) getPath(path *entities.Path) {
}
