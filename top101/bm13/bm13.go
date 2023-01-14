/*
author niewenlong
date 2023/1/14 21:21
description BM13 判断一个链表是否为回文结构
leetCodeUrl https://www.nowcoder.com/practice/3fed228444e740c8be66232ce8b87c2f?tpId=295&tqId=1008769&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPail(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var leftStart, rightStart = head, &ListNode{}
	if fast == nil {
		//偶数
		rightStart = slow
	} else {
		//奇数
		rightStart = slow.Next
	}
	rightStart = reverse(rightStart)

	for leftStart != nil && rightStart != nil {
		if leftStart.Val != rightStart.Val {
			return false
		}
		leftStart = leftStart.Next
		rightStart = rightStart.Next
	}
	return true
}

func reverse(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	var pre *ListNode
	var cur = node
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func main() {
	isPail(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 1},
			},
		},
	})
}
