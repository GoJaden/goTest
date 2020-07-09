package main

import (
	"fmt"
)

type NodeList struct {
	Val  int
	Next *NodeList
}

func main() {

	root := &NodeList{
		Val: 21,
		Next: &NodeList{
			Val: 22,
			Next: &NodeList{
				Val: 28,
				Next: &NodeList{
					Val: 67,
					Next: &NodeList{
						Val: 68,
						Next: &NodeList{
							Val: 78,
							Next: &NodeList{
								Val: 78,
								Next: &NodeList{
									Val: 99,
									Next: &NodeList{
										Val:  321,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	/*node := Quchong(root)*/

	node := ReverseNodeList(root)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

}

//反转链表
func ReverseNodeList(list *NodeList) *NodeList {
	if list == nil {
		return nil
	}
	cur := list
	var lastHead *NodeList
	for cur != nil {
		var next *NodeList
		if cur.Next != nil {
			next = cur.Next
			cur.Next = lastHead
			lastHead = cur
			cur = next
		} else {
			lastHead = &NodeList{
				Val:  cur.Val,
				Next: lastHead,
			}
			return lastHead
		}

	}
	return list
}

func Quchong(head *NodeList) *NodeList {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}
