/*
author niewenlong
date 2023/1/10 20:32
description BM9 删除链表的倒数第n个节点
leetCodeUrl https://www.nowcoder.com/practice/f95dcdafbde44b22a6d741baf71653f6?tpId=295&tqId=727&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//find n node
	var dummy = &ListNode{Next: head}
	var slow, fast = dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		if fast.Next == nil {
			if slow.Next != nil {
				slow.Next = slow.Next.Next
			}
		}
		fast = fast.Next
		slow = slow.Next
	}
	// write code here
	return dummy.Next
}

func main() {
	removeNthFromEnd(nil, 10)
}
