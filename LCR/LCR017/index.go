package LCR017

import "fmt"

/*
给定两个字符串 s 和 t 。返回 s 中包含 t 的所有字符的最短子字符串。如果 s 中不存在符合条件的子字符串，则返回空字符串 "" 。
如果 s 中存在多个符合条件的子字符串，返回任意一个。
注意： 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
*/
func myMinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	result := s

	start := 0
	end := 0

	strMap := make([]int, 256)
	indexMap := make([]int, 256)
	sum := len(t)

	for i := 0; i < len(t); i++ {
		strMap[t[i]-'A']++
		indexMap[t[i]-'A']++
	}
	// 再将最大的递归找到最小的
	for i := 0; i < len(s); i++ {
		if sum == 0 && (end-start) < len(result) {
			result = s[start : end+1]
			fmt.Println(result)
			// 重新装填，继续寻找
			if end != len(s) {
				// start 开始向前，直到找到第一个t中的字符
				start++
				for indexMap[s[start]-'A'] == 0 {
					start++
				}
				strMap[s[start]-'A']++
				sum++
			}
		}

		c := s[i]
		if strMap[c-'A'] > 0 {
			fmt.Print(strMap[c-'A'])
			fmt.Print(";")
			fmt.Println(c - 'A')
			// 说明刚发现第一个字符
			if sum == len(t) {
				start = i
				end = start + 1
			}
			strMap[c-'A']--
			sum--
			end = i
		}
	}

	return result

}

func myMinWindowPlus(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	result := s

	start := 0
	end := 0

	strMap := make(map[rune]int)
	indexMap := make(map[rune]int)
	sum := len(t)

	for _, c := range t {
		strMap[c]++
		indexMap[c]++
	}

	for i := range s {
		c := rune(s[i])
		if _, ok := strMap[c]; ok {
			if sum == len(t) {
				start = i
				end = start + 1
			}
			strMap[c]--
			sum--
			end = i
		}

		for sum == 0 {
			if end-start+1 < len(result) {
				result = s[start : end+1]
			}
			d := rune(s[start])
			if _, ok := indexMap[d]; ok {
				strMap[d]++
				if strMap[d] > 0 {
					sum++
				}
			}
			start++
		}
	}

	if len(result) > len(s) {
		return ""
	}
	return result
}

func minWindow(s string, t string) string {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range t {
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	start, length := 0, len(s)+1

	for right < len(s) {
		c := rune(s[right])
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			d := rune(s[left])
			left++

			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == len(s)+1 {
		return ""
	}
	return s[start : start+length]
}
