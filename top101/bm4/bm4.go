/*
author niewenlong
date 2023/1/3 21:13
description BM4 合并两个排序的链表
leetCodeUrl https://www.nowcoder.com/practice/d8b6b4358f774294a89de2a6ac4d9337?tpId=295&tqId=23267&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	var dummy = &ListNode{}
	var cur = dummy
	cur1 := pHead1
	cur2 := pHead2

	for cur1 != nil && cur2 != nil {
		if cur1.Val > cur2.Val {
			cur.Next = cur2
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur1 = cur1.Next
		}
		cur = cur.Next
	}

	if cur1 != nil {
		cur.Next = cur1
	}
	if cur2 != nil {
		cur.Next = cur2
	}

	return dummy.Next
}

func main() {
	Merge(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{Val: 5},
		},
	}, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{Val: 6},
		},
	})
}
