package main

import "goProject/gotest/basic/算法系列/链表/domain"

func main() {

}

//将两个有序的链表合并
func MergeSortedList(a *domain.Node, b *domain.Node) *domain.Node {
	if a == nil && b == nil {
		return nil
	}
	var lastNode = &domain.Node{
		Next: nil,
	}

	var result = lastNode
	for a != nil && b != nil {
		if a.Val > b.Val {
			lastNode.Next = b
			b = b.Next
		} else {
			lastNode.Next = a
			a = a.Next
		}
		lastNode = lastNode.Next
	}
	if a != nil {
		lastNode.Next = a
	}
	if b != nil {
		lastNode.Next = b
	}

	return result
}
