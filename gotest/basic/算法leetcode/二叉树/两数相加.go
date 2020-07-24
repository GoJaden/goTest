package main

import "fmt"

/*
//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
// Related Topics 链表 数学

*/

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}
	//l1  1234
	//l2  45678
	//result
	result := addTwoNumbers(l1, l2)

	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*递归法对每一个对应的位置累加，如果越界需要对下一位+1  。主要是需要判断清楚边界*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		if l1.Val >= 10 {
			//为了进行进位运算
			l2 = &ListNode{
				Val:  0,
				Next: nil,
			}
		} else {
			return l1
		}
	}
	temp := l1.Val + l2.Val
	//模值赋值给当前l1的值
	l1.Val = temp % 10
	if temp >= 10 {
		if l1.Next != nil {
			l1.Next.Val++
		} else {
			l1.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
	}
	l1.Next = addTwoNumbers(l1.Next, l2.Next)

	return l1
}
