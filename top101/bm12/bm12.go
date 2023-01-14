/*
author niewenlong
date 2023/1/13 17:40
description BM12 单链表的排序
leetCodeUrl https://www.nowcoder.com/practice/f23604257af94d939848729b1a5cda08?tpId=295&tqId=1008897&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortInList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//找到链表中间的那个节点 方便切割
	var dummy = &ListNode{Next: head}
	var slow, fast = dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//分治 切割链表成左右
	var right = slow.Next
	//切割
	slow.Next = nil
	//归并
	var left = head

	var leftNode = sortInList(left)
	var rightNode = sortInList(right)

	//合并两个有序链表 归并
	return Merge(leftNode, rightNode)
}
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	//left和right此时都可认为是有序的链表
	var dummy = &ListNode{}
	var cur = dummy
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		}
		cur = cur.Next
	}
	if pHead1 != nil {
		cur.Next = pHead1
	}
	if pHead2 != nil {
		cur.Next = pHead2
	}

	return dummy.Next
}

func main() {
	sortInList(&ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	})
}
