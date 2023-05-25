package entities

type Solution []*Path

type Path struct {
	Current *Room
	Next    *Path
	// HasAnt  bool
}
