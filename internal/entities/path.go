package entities

type Path struct {
	Start *Node
	Len   int
	Ants  int
}

type Node struct {
	Current *Room
	Next    *Node
}

func (n *Node) GetLast() *Node {
	if n.Next == nil {
		return n
	}
	return n.Next.GetLast()
}

func (n *Node) PrintList() string {
	res := ""
	start := n
	for start != nil {
		res += start.Current.Name + "->"
		start = start.Next
	}
	res += "nil"
	return res
}

func (n *Node) ChangeFirst(toAdd *Room) *Node {
	return &Node{
		Current: toAdd,
		Next:    n,
	}
}

func (n *Node) Reverse() *Node {
	if n == nil || n.Next == nil {
		return n
	}
	newHead := n.Next.Reverse()
	n.Next.Next = n
	n.Next = nil
	return newHead
}

func (n *Node) Len() int {
	count := 0
	current := n
	for current != nil {
		current = current.Next
		count++
	}
	return count
}
