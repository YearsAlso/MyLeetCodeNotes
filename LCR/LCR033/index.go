package LCR033

/*
给定一个字符串数组 strs ，将 变位词 组合在一起。 可以按任意顺序返回结果列表。

注意：若两个字符串中每个字符出现的次数都相同，则称它们互为变位词。
*/
func groupAnagrams(strs []string) [][]string {

	var result [][]string

	filterMap := make([]bool, len(strs))

	anagramMap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		if filterMap[i] {
			continue
		}
		anagramMap[strs[i]] = append(anagramMap[strs[i]], strs[i])
		for j := i; j < len(strs); j++ {
			if filterMap[j] {
				continue
			}
			if isAnagram(strs[i], strs[j]) {
				anagramMap[strs[i]] = append(anagramMap[strs[i]], strs[j])
				filterMap[j] = true
			}
		}
	}

	for _, v := range anagramMap {
		result = append(result, v)
	}

	return result
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return false
	}
	wordNumberMap := make([]int, 26)

	for i := 0; i < len(s); i++ {
		wordNumberMap[s[i]-'a']++
		wordNumberMap[t[i]-'a']--
	}

	for _, v := range wordNumberMap {
		if v != 0 {
			return false
		}
	}

	return true
}
