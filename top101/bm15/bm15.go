/*
author niewenlong
date 2023/1/15 13:20
description BM15 删除有序链表中重复的元素-I
leetCodeUrl https://www.nowcoder.com/practice/c087914fae584da886a0091e877f2c79?tpId=295&tqId=664&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
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
	var pre *ListNode
	var cur = head
	for cur != nil {
		if pre == nil {
			pre = cur
		} else {
			//compare
			if cur.Val == pre.Val {
				pre.Next = cur.Next
			} else {
				pre = pre.Next
			}
		}
		cur = cur.Next
	}
	// write code here
	return head
}

func main() {
	deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	})
}
