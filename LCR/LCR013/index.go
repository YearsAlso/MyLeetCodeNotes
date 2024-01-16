package LCR013

/*
给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回左上角 (row1, col1) 、右下角 (row2, col2) 的子矩阵的元素总和。
*/

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {

	yLength := len(matrix)
	if yLength == 0 {
		return NumMatrix{}
	}

	xLength := len(matrix[0])

	sumMatrix := make([][]int, yLength+1)
	sumMatrix[0] = make([]int, xLength+1)
	for i := 0; i < yLength; i++ {
		sumMatrix[i+1] = make([]int, xLength+1)
		for j := 0; j < xLength; j++ {
			sumMatrix[i+1][j+1] = sumMatrix[i+1][j] + sumMatrix[i][j+1] - sumMatrix[i][j] + matrix[i][j]
		}
	}

	return NumMatrix{matrix: sumMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.matrix[row2+1][col2+1] - this.matrix[row1][col2+1] - this.matrix[row2+1][col1] + this.matrix[row1][col1]
}
