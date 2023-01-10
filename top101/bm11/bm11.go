/*
author niewenlong
date 2023/1/10 21:00
description BM11 链表相加(二)
leetCodeUrl https://www.nowcoder.com/practice/c56f6c70fb3f4849bc56e33ff2a50b6b?tpId=295&tqId=1008772&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	head1 = reverse(head1)
	head2 = reverse(head2)
	var a = 0

	var head *ListNode
	var cur = head

	for head1 != nil || head2 != nil {
		a1, a2 := 0, 0
		if head1 != nil {
			a1 = head1.Val
			head1 = head1.Next
		}
		if head2 != nil {
			a2 = head2.Val
			head2 = head2.Next
		}
		b := a1 + a2 + a
		if b >= 10 {
			a = 1
			b = b - 10
		} else {
			a = 0
		}
		if cur == nil {
			cur = &ListNode{Val: b}
			head = cur
		} else {
			cur.Next = &ListNode{Val: b}
			cur = cur.Next
		}

	}
	if a != 0 {
		cur.Next = &ListNode{Val: a}
	}
	// write code here
	return reverse(head)
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	var cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	res := addInList(&ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 0}}},
		},
	}, &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{Val: 6}}},
		},
	})
	fmt.Println(res)
}
