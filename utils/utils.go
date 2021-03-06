package utils

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

// GenerateSquareMatrix generates an n x n square matrix with floats as elements.
// Returns a slice of a 2D array of n x n dimensions.
func GenerateSquareMatrix(n int) [][]float32 {
	squareMatrix := make([][]float32, n)
	for row := range squareMatrix {
		squareMatrix[row] = make([]float32, n)
		for col := range squareMatrix[row] {
			squareMatrix[row][col] = rand.Float32() * float32(rand.Intn(100))
		}
	}
	return squareMatrix
}

// MakeMatrix makes returns a numRows x numCols matrix slice
func MakeMatrix(numRows int, numCols int) [][]float32 {
	matrix := make([][]float32, numRows)
	for row := range matrix {
		matrix[row] = make([]float32, numCols)
	}
	return matrix
}

// PadMatrix will pad a matrix with zeroes until its dimension is 2^n x 2^n.
func PadMatrix(matrix [][]float32) [][]float32 {
	currentDim := len(matrix)
	fmt.Println("CurrentDim:", currentDim)
	newDim := int(math.Pow(2, math.Ceil(math.Log2(float64(currentDim)))))
	fmt.Println("NewDim:", newDim)

	newMatrix := MakeMatrix(newDim, newDim)

	CopyMatrix(newMatrix, matrix)

	return newMatrix
}

// GetMatrixSubBlocks will cut a matrix into four sub-blocks.
func GetMatrixSubBlocks(matrix [][]float32) ([][]float32, [][]float32, [][]float32, [][]float32) {

	dim := len(matrix)

	a := sliceMatrix(matrix, 0, dim/2, 0, dim/2)
	b := sliceMatrix(matrix, 0, dim/2, dim/2, dim)
	c := sliceMatrix(matrix, dim/2, dim, 0, dim/2)
	d := sliceMatrix(matrix, dim/2, dim, dim/2, dim)

	return a, b, c, d
}

// sliceMatrix will do what Go should do by default and slice a two-dimensional matrix into given dimensions.
func sliceMatrix(matrix [][]float32, begRow int, endRow int, begCol int, endCol int) [][]float32 {
	if matrix == nil {
		return nil
	}
	newMatrix := make([][]float32, endRow-begRow)
	copy(newMatrix, matrix[begRow:endRow])

	for row := range newMatrix {
		newMatrix[row] = newMatrix[row][begCol:endCol]
	}

	return newMatrix
}

// CopyMatrix will copy the elements of one matrix onto another matrix.
func CopyMatrix(newMatrix [][]float32, oldMatrix [][]float32) {

	var wg sync.WaitGroup

	if len(newMatrix) < len(oldMatrix) {
		fmt.Println("new:", newMatrix)
		fmt.Println("old:", oldMatrix)

		newMatrix = oldMatrix
		return
	}

	wg.Add(len(oldMatrix))
	for row := range oldMatrix {
		go copyRow(row, newMatrix, oldMatrix, &wg)
	}
	wg.Wait()
}

// copyRow will copy the elements of a row in a matrix to the same row in another matrix, in parallel.
func copyRow(row int, newMatrix [][]float32, oldMatrix [][]float32, wg *sync.WaitGroup) {
	for col := range oldMatrix[row] {
		newMatrix[row][col] = oldMatrix[row][col]
	}
	wg.Done()
}

// IsPowerOfTwo checks if an integer is a power of two.
func IsPowerOfTwo(x int) bool {
	return (x != 0) && ((x & (x - 1)) == 0)
}
