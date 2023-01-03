/*
author niewenlong
date 2023/1/3 21:48
description BM5 合并k个已排序的链表
leetCodeUrl https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6?tpId=295&tqId=724&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	var arrSize = len(lists)
	if arrSize == 0 {
		return nil
	}
	if arrSize == 1 {
		return lists[0]
	}
	mid := arrSize >> 1                                              //分治
	return merge(mergeKLists(lists[:mid]), mergeKLists(lists[mid:])) //归并
}

func merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	var dummy = &ListNode{}
	var cur = dummy
	cur1 := pHead1
	cur2 := pHead2

	for cur1 != nil && cur2 != nil {
		if cur1.Val > cur2.Val {
			cur.Next = cur2
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur1 = cur1.Next
		}
		cur = cur.Next
	}

	if cur1 != nil {
		cur.Next = cur1
	}
	if cur2 != nil {
		cur.Next = cur2
	}

	return dummy.Next
}

func main() {
	mergeKLists(nil)
}
