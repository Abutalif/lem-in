package usecases

import (
	"fmt"

	"lem-in/internal/entities"
)

type Solver interface {
	Solve(*entities.Anthill) []entities.Solution
}

type naive struct{}

func NewNaiveSolver() Solver {
	return &naive{}
}

// ***HERE***
func (n *naive) Solve(anthill *entities.Anthill) []entities.Solution {
	validPaths := make([]entities.Solution, 0)
	startRoom := anthill.GetStart()
	fmt.Println(startRoom) // placeholder to not my code RED.
	// shortestPath := n.visitNeighbors(startRoom)
	// validPaths = append(validPaths, shortestPath)
	return validPaths
}

// func (n *naive) visitNeighbors(room *entities.Room) entities.Solution {
// 	room.Visited = true
// 	for _, neighbor := range room.Connections {
// 		if !neighbor.Visited {
// 			if neighbor.Kind == entities.Regular {
// 				pathFrom Neighbor:=n.visit
// 				return n.visitNeighbors(neighbor)
// 			} else if neighbor.Kind == entities.End =
// 			n.visitNeighbors(neighbor)
// 			if
// 		}
// 	}

// 	return entities.Solution{}
// }

// func
