/*
author niewenlong
date 2023/1/2 21:32
description BM3 链表中的节点每k个一组翻转
leetCodeUrl https://www.nowcoder.com/practice/b49c3dc907814e9bbfa8437c251b028e?tpId=295&tqId=722&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var tail = head
	for i := 0; i < k; i++ {
		if tail == nil {
			//不够K个 不翻转
			return head
		}
		tail = tail.Next
	}
	//翻转
	var pre *ListNode
	cur := head
	for cur != tail {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	//递归下一组
	head.Next = reverseKGroup(tail, k)
	return pre
}

func main() {
	reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}, 2)
}
