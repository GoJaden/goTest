package main

import (
	"fmt"
	"strings"
	"time"
)

type TreeNode1 struct {
	Data  int
	Left  *TreeNode1
	Right *TreeNode1
}

func main() {

	fmt.Println(strings.Repeat("a", 2))

	t := time.Time{}
	fmt.Println(t)

	r11 := &TreeNode1{
		Data:  32,
		Left:  nil,
		Right: nil,
	}
	l12 := &TreeNode1{
		Data:  123,
		Left:  nil,
		Right: nil,
	}
	r1 := &TreeNode1{
		Data:  88,
		Left:  l12,
		Right: r11,
	}
	l1 := &TreeNode1{
		Data: 321,
		Left: &TreeNode1{
			Data: 12,
			Left: &TreeNode1{
				Data:  323,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode1{
				Data:  122,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}

	root := &TreeNode1{
		Data:  10,
		Left:  l1,
		Right: r1,
	}

	fmt.Println(root)

	fmt.Println(ReverseTreeNode1(root))
}

//递归方法翻转二叉树
//递归条件 ，遍历到根节点
//翻转内容，左右节点交换
func ReverseTreeNode1(root *TreeNode1) *TreeNode1 {
	if root == nil {
		return nil
	}
	left := ReverseTreeNode1(root.Left)
	right := ReverseTreeNode1(root.Right)
	root.Left = right
	root.Right = left
	return root
}
