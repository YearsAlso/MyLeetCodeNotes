package LCR014

import (
	"sort"
	"strings"
)

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。

换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// 将字符串切片
	s1Slice := strings.Split(s1, "")
	s2Slice := strings.Split(s2, "")

	// 对两个数组进行排序
	sort.Strings(s1Slice)
	sort.Strings(s2Slice)

	s1Index := 0

	for i := 0; i < len(s2Slice); i++ {
		if s2Slice[i] != s1Slice[s1Index] {
			s1Index = 0
		} else {
			s1Index += 1
		}

		if s1Index == len(s1Slice)-1 {
			return true
		}
	}

	return false
}
