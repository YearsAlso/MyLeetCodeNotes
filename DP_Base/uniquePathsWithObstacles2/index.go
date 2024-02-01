package uniquePathsWithObstacles2

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 && j == 0 && obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			obstacleGrid[i][j] = getVal(obstacleGrid, i-1, j) + getVal(obstacleGrid, i, j-1)
		}
	}

	return obstacleGrid[height-1][width-1]
}

func getVal(obstacleGrid [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}

	return obstacleGrid[i][j]
}
