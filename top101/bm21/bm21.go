/*
author niewenlong
date 2023/2/20 19:33
description BM21 旋转数组的最小数字
leetCodeUrl https://www.nowcoder.com/practice/9f3231a991af4f55b95579b44b7a01ba?tpId=295&tqId=23269&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
*/

package main

/**
 *
 * @param rotateArray int整型一维数组
 * @return int整型
 */
func minNumberInRotateArray(rotateArray []int) int {
	l, r := 0, len(rotateArray)-1
	for l < r {
		mid := (r-l)>>1 + l
		if rotateArray[mid] > rotateArray[r] {
			l = mid + 1
		} else if rotateArray[mid] < rotateArray[r] {
			r = mid
		} else {
			r--
		}
	}
	return rotateArray[l]
}

func main() {
	minNumberInRotateArray([]int{3, 4, 5, 1, 2})
}
