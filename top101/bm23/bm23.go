/*
author niewenlong
date 2023/2/21 18:59
description BM23 二叉树的前序遍历
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param root TreeNode类
 * @return int整型一维数组
 */
func preorderTraversal(root *TreeNode) []int {
	var result []int
	t(root, &result)
	return result
}

func t(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	t(node.Left, result)
	t(node.Right, result)
}

func main() {
	preorderTraversal(nil)
}
