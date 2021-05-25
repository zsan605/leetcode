package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	rowMax := len(matrix)
	colMax := len(matrix[0])
	mid := isSameLine(0, 0, matrix, rowMax, colMax)
	if !mid {
		return false
	}

	for i := 1; i < rowMax; i++ {
		if isSameLine(i, 0, matrix, rowMax, colMax) == false {
			return false
		}
	}

	for j := 1; j < colMax; j++ {
		if isSameLine(0, j, matrix, rowMax, colMax) == false {
			return false
		}
	}
	return true
}

func isSameLine(row, col int, matrix [][]int, rowMax, colMax int) bool {
	var temp = matrix[row][col]
	for i, j := row, col; i < rowMax && j < colMax; i, j = i+1, j+1 {

		if temp != matrix[i][j] {
			return false
		}
	}
	return true
}
