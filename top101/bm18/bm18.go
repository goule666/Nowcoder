/*
author niewenlong
date 2023/2/1 19:10
description BM18 二维数组中的查找
leetCodeUrl https://www.nowcoder.com/practice/abc3fe2ce8e146608e868a70efebf62e?tpId=295&tqId=23256&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */
func find(target int, array [][]int) bool {
	if array == nil || len(array) == 0 {
		return false
	}
	x, y := len(array[0]), len(array)
	for i := 0; i < x; i++ {
		leftIndex, rightIndex := 0, y-1
		for leftIndex <= rightIndex {
			mid := (leftIndex + rightIndex) / 2
			if target == array[mid][i] {
				return true
			} else if target > array[mid][i] {
				leftIndex = mid + 1
			} else {
				rightIndex = mid - 1
			}
		}
	}
	return false
}

func main() {
	println(find(19, [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}))
}
