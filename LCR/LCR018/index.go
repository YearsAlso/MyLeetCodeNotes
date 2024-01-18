package LCR018

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	// 字符串提取字母
	start := 0
	end := len(s) - 1
	for i := 0; i < len(s); i++ {
		startC := s[start]
		if startC >= 'A' && startC <= 'Z' {
			startC = startC - 'A' + 'a'
		}

		if !(startC >= '0' && startC <= '9') && !(startC >= 'a' && startC <= 'z') {
			start++
			continue
		}

		endC := s[end]
		if endC >= 'A' && endC <= 'Z' {
			endC = endC - 'A' + 'a'
		}

		if !(endC >= '0' && endC <= '9') && !(endC >= 'a' && endC <= 'z') {
			end--
			continue
		}
		if startC != endC {
			return false
		}
		start++
		end--
		if start >= end {
			return true
		}
	}

	if start >= end {
		return true
	}

	return false
}
