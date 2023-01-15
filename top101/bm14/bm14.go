/*
author niewenlong
date 2023/1/15 14:09
description BM14 链表的奇偶重排
leetCodeUrl https://www.nowcoder.com/practice/02bf49ea45cd486daa031614f9bd6fc3?tpId=295&tqId=1073463&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//a是奇数 b是偶数
	var a, b = head, head.Next
	var bTemp = b
	for b != nil && b.Next != nil {
		a.Next = b.Next
		a = a.Next
		b.Next = a.Next
		b = b.Next
	}
	a.Next = bTemp
	// write code here
	return head
}

func main() {
	oddEvenList(nil)
}
