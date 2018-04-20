package utils

import (
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

// PadMatrix will pad a matrix with zeroes until its dimension is 2^n x 2^n.
func PadMatrix(matrix [][]float32) [][]float32 {
	currentDim := len(matrix)
	newDim := int(math.Ceil(math.Log2(float64(currentDim))))

	newMatrix := make([][]float32, newDim)
	for row := range newMatrix {
		newMatrix[row] = make([]float32, newDim)
	}
	copyMatrix(newMatrix, matrix)

	return newMatrix
}

// copyMatrix will copy the elements of one matrix onto another matrix.
func copyMatrix(newMatrix [][]float32, oldMatrix [][]float32) {

	var wg sync.WaitGroup

	wg.Add(len(oldMatrix))
	for row := range oldMatrix {
		go copyRow(row, newMatrix, oldMatrix, &wg)
	}
	wg.Wait()
}

// copyRow will copy the elements of a row in a matrix to the same row in another matrix, in parallel.
func copyRow(row int, newMatrix [][]float32, oldMatrix [][]float32, wg *sync.WaitGroup) {
	defer wg.Done()
	for col := range oldMatrix[row] {
		newMatrix[row][col] = oldMatrix[row][col]
	}
}

// IsPowerOfTwo checks if an integer is a power of two.
func IsPowerOfTwo(x int) bool {
	return (x != 0) && ((x & (x - 1)) == 0)
}
