package entities

type Path struct {
	Start *Node
	Len   int
	Ants  int
}

func (p *Path) Unvisit() {
	start := p.Start
	for start != nil {
		start.Current.Visited = false
		start = start.Next
	}
}

func (p *Path) HasRoom(room *Room) bool {
	start := p.Start
	for start != nil {
		if start.Current == room {
			return true
		}
		start = start.Next
	}

	return false
}
