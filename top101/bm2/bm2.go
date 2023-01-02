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
	if head.Next == nil {
		return head
	}
	index := 0
	var a, b, prev *ListNode
	cur := head
	for cur != nil {
		index++ //1 //2
		if index+1 == m {
			a = cur //a->1
		}
		if index == m {
			b = cur                   //b->2
			for i := m; i <= n; i++ { //2 //3 //4
				//开始反转
				temp := cur.Next
				cur.Next = prev
				prev = cur
				cur = temp
				//反转结束 prev->4 cur ->5
			}
			break
		}
		cur = cur.Next
	}
	if a == nil {
		head = prev
	} else {
		a.Next = prev
	}
	b.Next = cur
	return head
}

func main() {
	reverseBetween(&ListNode{
		Val:  3,
		Next: &ListNode{Val: 5},
	}, 1, 2)
}
