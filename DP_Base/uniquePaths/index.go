package uniquePaths

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？
*/

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}
