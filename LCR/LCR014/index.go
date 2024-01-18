package LCR014

import (
	"fmt"
	"sort"
)

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。

换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// 将s1 切片并排序
	s1 = sortString(s1)

	for i := 0; i < 1+(len(s2)-len(s1)); i++ {
		tmpStr := sortString(s2[i : i+len(s1)])
		fmt.Print(tmpStr + ";\t")
		if tmpStr == s1 {
			return true
		}
	}

	return false
}

func sortString(s string) string {
	// 将字符串转换为字符切片
	charSlice := []rune(s)

	// 对字符切片进行排序
	sort.Slice(charSlice, func(i, j int) bool {
		return charSlice[i] < charSlice[j]
	})

	// 将排序后的字符切片转换回字符串
	return string(charSlice)
}
