/*
author niewenlong
date 2023/1/9 20:45
description BM8 链表中倒数最后k个结点
leetCodeUrl https://www.nowcoder.com/practice/886370fe658f41b498d40fb34ae76ff9?tpId=295&tqId=1377477&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	if pHead == nil || k < 1 {
		return nil
	}
	// write code here
	var slow = pHead
	var fast = pHead
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		if fast.Next == nil {
			return slow
		}
		fast = fast.Next
		slow = slow.Next
	}
	return nil
}

func main() {
	FindKthToTail(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}, 2)
}
