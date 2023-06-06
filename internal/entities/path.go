package entities

import "fmt"

type Path struct {
	Start *Node
	Len   uint
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

func (n *Node) PrintList() {
	start := n
	for start.Next != nil {
		fmt.Printf("%v <- ", start.Current.Name)
		start = start.Next
	}
	fmt.Printf("%v\n", start.Current.Name)
}
