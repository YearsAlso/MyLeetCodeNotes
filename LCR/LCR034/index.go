package LCR034

/*
某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。

给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回 false。
*/
func isAlienSorted(words []string, order string) bool {

	indexMap := make(map[byte]int)

	for i := 0; i < len(order); i++ {
		indexMap[order[i]] = i
	}

	for i := 0; i < len(words); i++ {
		word := words[i]
		chartIndex := indexMap[word[0]]
		for j := 0; j < len(word); j++ {
			if chartIndex < indexMap[word[j]] {
				return false
			} else if chartIndex == indexMap[word[j]] {
				continue
			} else {
				chartIndex = indexMap[word[j]]
			}
		}
	}

	return true
}
