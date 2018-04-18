package algebra

import (
	"sync"
)

// AddMatrices adds two matrices in the form
// [ a b ] + [ e f ] = [ a+e b+f ]
// [ c d ]   [ g h ]   [ c+g d+h ]
// It does so parallelized row-by-row
func AddMatrices(matrixA [][]float32, matrixB [][]float32) [][]float32 {
	if matrixA == nil || matrixB == nil {
		return nil
	}

	aNumRows := len(matrixA)
	aNumCols := len(matrixA[0])
	bNumRows := len(matrixB)
	bNumCols := len(matrixB[0])

	// Make sure the matrices are of equal dimensions
	if aNumCols != bNumCols || aNumRows != bNumRows {
		return nil
	}

	// Instantiate the goroutine pool.
	var wg sync.WaitGroup

	// Instantiate a matrix to return
	result := make([][]float32, aNumRows)

	// Add the number of rows to the goroutine pool
	wg.Add(aNumRows)
	for row := range matrixA {
		result[row] = make([]float32, aNumCols)
		go addRow(row, matrixA, matrixB, result, &wg)
	}
	// Wait until all goroutines have finished before returning
	wg.Wait()

	return result

}

// Add two rows of two matrices together in a goroutine
func addRow(row int, matrixA [][]float32, matrixB [][]float32, result [][]float32, wg *sync.WaitGroup) {
	defer wg.Done()
	for col := range matrixA[row] {
		result[row][col] = matrixA[row][col] + matrixB[row][col]
	}
}

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
