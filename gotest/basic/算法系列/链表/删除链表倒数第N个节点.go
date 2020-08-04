package main

import (
	"fmt"
	"goProject/gotest/basic/算法系列/链表/domain"
)

func main() {
	node := DelPoiNode2(&domain.Node{
		Next: &domain.Node{
			Next: &domain.Node{
				Next: &domain.Node{
					Next: &domain.Node{
						Next: &domain.Node{

							Val: 6,
						},
						Val: 5,
					},
					Val: 4,
				},
				Val: 3,
			},
			Val: 2,
		},
		Val: 1,
	}, 2)

	for ; node != nil; node = node.Next {
		fmt.Println(node.Val)

	}
}

/*删除链表倒数第N个节点
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
*/

//这个方式，是通过获取每次移动指定步长的方式，知道尾节点为空
func DelPoiNode(root *domain.Node, n int) *domain.Node {
	if root == nil {
		return nil
	}
	//记录上一个节点
	var lastNode *domain.Node
	var curNode = root
	var tarNode = root.GetLastNNode(n)
	for curNode != nil && tarNode != nil {
		lastNode = curNode
		curNode = curNode.Next
		tarNode = curNode.GetLastNNode(n)
	}
	lastNode.Next = curNode.Next

	return root
}

//双指针的方式，通过控制两个指针之间的间距，达到获取指定间隙间数据的办法
func DelPoiNode2(root *domain.Node, n int) *domain.Node {
	//哨兵节点
	var result = new(domain.Node)
	result.Next = root
	var curNode = result
	var preNode = root
	var lastNode *domain.Node
	var i int = 1
	for preNode != nil {
		if i >= n {
			lastNode = curNode
			curNode = curNode.Next
		}
		i++
		preNode = preNode.Next
	}
	lastNode.Next = lastNode.Next.Next

	return result.Next
}
