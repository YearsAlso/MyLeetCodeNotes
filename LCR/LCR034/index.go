package LCR034

import "fmt"

func isAlienSorted(words []string, order string) bool {

	indexMap := make(map[byte]int)

	l := len(order)
	for i := 0; i < l; i++ {
		indexMap[order[i]] = l - i
	}

	// 找到最大长度的单词
	maxLen := 0
	for _, word := range words {
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	for i := 0; i < maxLen; i++ {
		for j := 0; j < len(words)-1; j++ {
			if i >= len(words[j]) {
				continue
			}

			if indexMap[words[j][i]] < indexMap[words[j+1][i]] {
				fmt.Println(string(words[j][i]), string(words[j+1][i]))
				return false
			}
		}
	}

	return true

}
