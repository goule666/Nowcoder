/*
author niewenlong
date 2023/1/14 21:21
description BM13 判断一个链表是否为回文结构
leetCodeUrl https://www.nowcoder.com/practice/3fed228444e740c8be66232ce8b87c2f?tpId=295&tqId=1008769&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPail(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//判断奇偶性 如果fast不等于nil 则为奇数
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	queue := list.New()
	for head != nil {
		if fast != nil && head == slow {
			//如果是奇数的话 跳过中间那个值
			head = head.Next
			continue
		}
		node := queue.Back()
		if node != nil {
			if node.Value == head.Val {
				//move
				queue.Remove(node)
			} else {
				//add
				queue.PushBack(head.Val)
			}
		} else {
			queue.PushBack(head.Val)
		}
		head = head.Next
	}
	return queue.Len() == 0
}

func main() {
	isPail(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 1},
			},
		},
	})
}
