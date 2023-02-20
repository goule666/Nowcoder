/*
author niewenlong
date 2023/2/20 19:56
description BM22 比较版本号
*/

package main

import (
	"strconv"
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 比较版本号
 * @param version1 string字符串
 * @param version2 string字符串
 * @return int整型
 */
func compare(version1 string, version2 string) int {
	versionSplit1 := strings.Split(version1, ".")
	versionSplit2 := strings.Split(version2, ".")
	var reverse bool
	if len(versionSplit1) < len(versionSplit2) {
		versionSplit1, versionSplit2 = versionSplit2, versionSplit1
		reverse = true
	}
	for index, s1 := range versionSplit1 {
		versionInt1, _ := strconv.ParseInt(s1, 10, 64)
		var versionInt2 = int64(0)
		if len(versionSplit2)-1 >= index {
			versionInt2, _ = strconv.ParseInt(versionSplit2[index], 10, 64)
		}
		if reverse {
			if versionInt1 > versionInt2 {
				return -1
			} else if versionInt1 < versionInt2 {
				return 1
			} else {
				continue
			}
		} else {
			if versionInt1 > versionInt2 {
				return 1
			} else if versionInt1 < versionInt2 {
				return -1
			} else {
				continue
			}
		}
	}
	return 0
}

func main() {
	println(compare("1.2.0.1", "1.2.3"))
}
