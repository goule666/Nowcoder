/*
author niewenlong
date 2023/1/8 18:29
description BM7 链表中环的入口结点
leetCodeUrl https://www.nowcoder.com/practice/253d2c59ec3e4bc68da16833f79a38e4?tpId=295&tqId=23449&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil {
		return pHead
	}
	//find meet node
	var start = pHead
	var meet = &ListNode{}
	var fast = pHead
	var slow = pHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			//find meet
			meet = slow
			for start != nil && start.Next != nil {
				if start == meet {
					//first
					return start
				}
				start = start.Next
				meet = meet.Next
			}
			return nil
		}
	}
	return nil
}

func main() {
	EntryNodeOfLoop(nil)
}
