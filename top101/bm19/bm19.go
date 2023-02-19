/*
author niewenlong
date 2023/2/19 21:08
description BM19 寻找峰值
leetCodeUrl https://www.nowcoder.com/practice/fcf87540c4f347bcb4cf720b5b350c76?tpId=295&tqId=2227748&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */
func findPeakElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	//0 1 2 3 4
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	findPeakElement(nil)
}
