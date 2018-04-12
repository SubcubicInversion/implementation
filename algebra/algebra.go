package algebra

import (
	"sync"
)

func addMatrices(matrixA [][]float32, matrixB [][]float32) [][]float32 {
	if matrixA == nil || matrixB == nil {
		return nil
	}

	aNumRows := len(matrixA)
	aNumCols := len(matrixA[0])
	bNumRows := len(matrixB)
	bNumCols := len(matrixB[0])

	if aNumCols != bNumCols || aNumRows != bNumRows {
		return nil
	}

	var wg sync.WaitGroup
	result := make([][]float32, aNumRows, aNumCols)

	wg.Add(aNumRows)
	for i := range matrixA {
		go addRow(i, matrixA, matrixB, result, &wg)
	}
	wg.Wait()

	return result

}

func addRow(irow int, matrixA [][]float32, matrixB [][]float32, result [][]float32, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range matrixA[irow] {
		result[irow][i] = matrixA[irow][i] + matrixB[irow][i]
	}
}

func scalarMultiply(scalar float32, matrix [][]float32) {
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
}

func scaleRow(irow int, scalar float32, matrix [][]float32, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range matrix[irow] {
		matrix[irow][i] = matrix[irow][i] * scalar
	}
}
