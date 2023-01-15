/*
author niewenlong
date 2023/1/15 13:37
description BM16 删除有序链表中重复的元素-II
leetCodeUrl https://www.nowcoder.com/practice/71cef9f8b5564579bf7ed93fbe0b2024?tpId=295&tqId=663&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var dummy = &ListNode{Next: head}
	var cur = dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == temp {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func main() {
	deleteDuplicates(nil)
}
