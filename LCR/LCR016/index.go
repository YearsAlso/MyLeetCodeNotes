package LCR016

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。
func lengthOfLongestSubstring(s string) int {

	start := 0
	end := 0

	strInt := make([]int, 256)

	maxLen := 0

	for i := 0; i < len(s); i++ {
		strInt[s[i]-'a']++
		end++
		// 如果当前字符出现次数大于1，那么就需要移动start指针，说明有重复字符
		if strInt[s[i]-'a'] > 1 {
			for strInt[s[i]-'a'] > 1 {
				strInt[s[start]-'a']--
				start++
			}
		}
		maxLen = max(maxLen, end-start)
	}

	return maxLen
}
