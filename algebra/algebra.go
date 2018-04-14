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
	result := make([][]float32, aNumRows, aNumCols)

	// Add the number of columns to the goroutine pool
	wg.Add(aNumRows)
	for i := range matrixA {
		go addRow(i, matrixA, matrixB, result, &wg)
	}
	// Wait until all goroutines have finished before returning
	wg.Wait()

	return result

}

// Add two rows of two matrices together in a goroutine
func addRow(irow int, matrixA [][]float32, matrixB [][]float32, result [][]float32, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range matrixA[irow] {
		result[irow][i] = matrixA[irow][i] + matrixB[irow][i]
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
	for i := range matrix {
		go scaleRow(i, scalar, matrix, &wg)
	}
	wg.Wait()
}

// Multiply a row of a matrix in a goroutine
func scaleRow(irow int, scalar float32, matrix [][]float32, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range matrix[irow] {
		matrix[irow][i] = matrix[irow][i] * scalar
	}
}
