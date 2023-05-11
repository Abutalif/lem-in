package usecases

import "lem-in/internal/entities"

type Solver interface{
	Solve() entities.Solution
}

type naive struct {}

func NewNaiveSolver() Solver {
	return &naive{}
}

func (n *naive) Solve() entities.Solution {
	return entities.Solution{}
}