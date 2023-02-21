/*
author niewenlong
date 2023/2/21 19:43
description BM24 二叉树的中序遍历
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
	*result = append(*result, node.Val)
	t(node.Right, result)
}

func main() {
	postorderTraversal(nil)
}
