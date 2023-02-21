/*
author niewenlong
date 2023/2/21 19:43
description BM25 二叉树的后序遍历
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	t(root, &result)
	return result
}

func t(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	t(node.Left, result)
	t(node.Right, result)
	*result = append(*result, node.Val)
}

func main() {
	postorderTraversal(nil)
}
