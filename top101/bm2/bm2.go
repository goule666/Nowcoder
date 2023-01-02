/*
author niewenlong
date 2023/1/2 18:55
description BM2 链表内指定区间反转
leetCodeUrl https://www.nowcoder.com/practice/b58434e200a648c589ca2063f1faf58c?tpId=295&tqId=654&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var dummy = &ListNode{Next: head}
	preStart, start := dummy, head
	for i := 0; i < m-1; i++ {
		preStart = start
		start = start.Next
	}
	for i := 0; i < n-m; i++ {
		temp := start.Next
		start.Next = temp.Next
		temp.Next = preStart.Next
		preStart.Next = temp
	}
	return dummy.Next
}

func main() {
	reverseBetween(&ListNode{
		Val:  3,
		Next: &ListNode{Val: 5},
	}, 1, 2)
}
