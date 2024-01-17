package LCR015

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 变位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
func findAnagrams(s string, p string) []int {
	var res []int
	left, right, match := 0, 0, 0
	needs, window := make(map[rune]int), make(map[rune]int)

	for _, char := range p {
		needs[char]++
	}

	for right < len(s) {
		c := rune(s[right])
		if _, ok := needs[c]; ok {
			window[c]++
			if window[c] == needs[c] {
				match++
			}
		}
		right++

		for match == len(needs) {
			if right-left == len(p) {
				res = append(res, left)
			}
			c := rune(s[left])
			if _, ok := needs[c]; ok {
				window[c]--
				if window[c] < needs[c] {
					match--
				}
			}
			left++
		}
	}
	return res
}
