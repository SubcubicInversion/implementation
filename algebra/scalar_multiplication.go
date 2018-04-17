package algebra

import (
	"sync"
)

// ScalarMultiply multiplies a matrix by a scalar in the form:
// g[ a b ] = [ ga gb ]
//  [ c d ]   [ gc gd ]
func ScalarMultiply(scalar float32, matrix [][]float32) {
	if matrix == nil {
		return
	}
	if scalar == 1 {
		return
	}

	var wg sync.WaitGroup

	numRows := len(matrix)

	wg.Add(numRows)
	for row := range matrix {
		go scaleRow(row, scalar, matrix, &wg)
	}
	wg.Wait()
}

// Multiply a row of a matrix in a goroutine
func scaleRow(row int, scalar float32, matrix [][]float32, wg *sync.WaitGroup) {
	defer wg.Done()
	for col := range matrix[row] {
		matrix[row][col] = matrix[row][col] * scalar
	}
}
