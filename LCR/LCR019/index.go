package LCR019

import "fmt"

// 给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。
func validPalindrome(s string) bool {

	haveDel := false

	start := 0
	end := len(s) - 1

	for {
		c := s[start]

		if c != s[end] {
			if haveDel {
				return false
			}
			haveDel = true

			// 如果只有一个没有匹配的，那么就是可以的
			if end-start == 1 {
				return true
			}

			if s[start] == s[end-1] && s[start+1] == s[end] {
				return loopFind(s[start:end]) || loopFind(s[start+1:end+1])
			}

			if s[start] == s[end-1] {
				end--
				printPr(s, start, end)
				fmt.Println("end++")
				continue
			}

			if s[start+1] == s[end] {
				start++
				printPr(s, start, end)
				fmt.Println("start--")
				continue
			}

			return false
		}

		if end < start {
			return true
		}

		end--
		start++
	}

	return true
}

func loopFind(s string) bool {
	start := 0
	end := len(s) - 1

	fmt.Println(s)

	for {
		c := s[start]
		if c != s[end] {
			return false
		}

		if end <= start {
			return true
		}
		end--
		start++
	}

	return true
}

func printPr(s string, start int, end int) {
	fmt.Println(s)
	fmt.Printf("%*s", start, "")
	fmt.Print("^")
	fmt.Printf("%*s", end-start-1, "")
	fmt.Print("^\n")
}
