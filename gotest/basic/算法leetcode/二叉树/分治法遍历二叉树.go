package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func main1() {

	r11 := &TreeNode{
		Data:  32,
		Left:  nil,
		Right: nil,
	}
	l12 := &TreeNode{
		Data:  123,
		Left:  nil,
		Right: nil,
	}
	r1 := &TreeNode{
		Data:  88,
		Left:  l12,
		Right: r11,
	}
	l1 := &TreeNode{
		Data:  321,
		Left:  nil,
		Right: nil,
	}

	root := &TreeNode{
		Data:  10,
		Left:  l1,
		Right: r1,
	}

	fmt.Println(DivideAndConquer(root))

	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	// 返回条件处理
	if root == nil {
		return 0
	}
	// divide：分左右子树分别计算
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	bts, _ := json.Marshal(root)
	// conquer：合并左右子树结果
	if left > right {
		fmt.Printf("该节点%s的深度为:%v\n", string(bts), left+1)
		return left + 1
	}
	fmt.Printf("该节点%s的深度为:%v\n", string(bts), right+1)
	return right + 1
}

//分而治之的方式
func DivideAndConquer(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	left := DivideAndConquer(root.Left)
	right := DivideAndConquer(root.Right)
	res = append(res, root.Data)
	res = append(res, left...)
	res = append(res, right...)
	return res
}
