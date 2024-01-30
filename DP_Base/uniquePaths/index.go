package uniquePaths

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？
*/

func uniquePaths1(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	return uniquePaths1(m-1, n) + uniquePaths1(m, n-1)
}

// 这个方法会占用较多的内存，可以使用一行数组进行保存
// 递归转遍历，可以将递归的内容写出公式，然后进行推演，将所有的条件放到数组中求解，就能得到最终的结果
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
