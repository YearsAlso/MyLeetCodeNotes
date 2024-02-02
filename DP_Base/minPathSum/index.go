package minPathSum

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
func minPathSum(grid [][]int) int {
	return minPathSumCore(grid, 0, 0, 0)
}

func minPathSumCore(grid [][]int, x int, y int, sum int) int {
	sum += grid[y][x]
	if x == (len(grid[0])-1) && y == (len(grid)-1) {
		return sum
	}

	if x == len(grid[0])-1 {
		return minPathSumCore(grid, x, y+1, sum)
	}

	if y == len(grid)-1 {
		return minPathSumCore(grid, x+1, y, sum)
	}

	return min(minPathSumCore(grid, x+1, y, sum), minPathSumCore(grid, x, y+1, sum))
}
