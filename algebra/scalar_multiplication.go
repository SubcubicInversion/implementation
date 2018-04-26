package algebra

import (
	"sync"

	"github.com/SubcubicInversion/implementation/utils"
)

// ScalarMultiply multiplies a matrix by a scalar in the form:
// g[ a b ] = [ ga gb ]
//  [ c d ]   [ gc gd ]
func ScalarMultiply(scalar float32, matrix [][]float32) [][]float32 {
	if matrix == nil {
		return nil
	}
	if scalar == 1 {
		return nil
	}

	numRows := len(matrix)
	numCols := len(matrix[0])

	newMatrix := utils.MakeMatrix(numRows, numCols)

	utils.CopyMatrix(newMatrix, matrix)

	var wg sync.WaitGroup

	wg.Add(numRows)
	for row := range newMatrix {
		go scaleRow(row, scalar, newMatrix, &wg)
	}
	wg.Wait()

	return newMatrix
}

// Multiply a row of a matrix in a goroutine
func scaleRow(row int, scalar float32, matrix [][]float32, wg *sync.WaitGroup) {
	defer wg.Done()
	for col := range matrix[row] {
		matrix[row][col] = matrix[row][col] * scalar
	}
}
