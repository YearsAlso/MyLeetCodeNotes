package minimumTotal

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] = triangle[i][j] + triangle[i-1][j]
				continue
			}

			if j == len(triangle[i])-1 {
				triangle[i][j] = triangle[i][j] + triangle[i-1][j-1]
				continue
			}

			triangle[i][j] = triangle[i][j] + min(triangle[i-1][j], triangle[i-1][j-1])
		}
	}

	return triangle[len(triangle)-1][0]
}
