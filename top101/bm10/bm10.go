/*
author niewenlong
date 2023/1/10 20:41
description BM10 两个链表的第一个公共结点
leetCodeUrl https://www.nowcoder.com/practice/6ab1d9a29e88450685099d45c9e31e46?tpId=295&tqId=23257&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	var pHead1Size, pHead2Size = 0, 0
	var p11, p21 = pHead1, pHead2
	for p11 != nil {
		pHead1Size++
		p11 = p11.Next
	}
	for p21 != nil {
		pHead2Size++
		p21 = p21.Next
	}
	//找出长度长的那一个
	var p1, p2 = pHead1, pHead2
	var diff = 0
	if pHead1Size < pHead2Size {
		p1, p2 = pHead2, pHead1
		diff = pHead2Size - pHead1Size
	} else {
		diff = pHead1Size - pHead2Size
	}

	for i := 0; i < diff; i++ {
		p1 = p1.Next
	}

	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	// write code here
	return nil
}

func main() {
	FindFirstCommonNode(nil, nil)
}
