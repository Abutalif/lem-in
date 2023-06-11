package entities

import "fmt"

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

func (n *Node) PrintList() {
	start := n
	for start.Next != nil {
		fmt.Printf("%v -> ", start.Current.Name)
		start = start.Next
	}
	fmt.Printf("%v\n", start.Current.Name)
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
	count := 1
	current := n
	for current != nil {
		current = current.Next
		count++
	}
	return count
}
