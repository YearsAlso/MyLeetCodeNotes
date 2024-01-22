package LCR020

/*
给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
*/
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	ans := 0
	for t := 0; t < n; t++ {
		for i := 0; i+t < n; i++ {
			j := i + t
			if t == 0 {
				dp[i][j] = true
			} else if t == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] {
				ans++
			}
		}
	}
	return ans
}
