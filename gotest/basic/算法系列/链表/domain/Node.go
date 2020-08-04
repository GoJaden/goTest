package domain

type Node struct {
	Next *Node
	Val  int
}

func (root *Node) GetLastNNode(n int) *Node {
	for i := 0; i < n; i++ {
		root = root.Next
	}
	return root
}
