package utils

import (
	"math"
	"math/rand"
	"sync"
)

// GenerateSquaredMatrix generates an n x n square matrix with floats as elements.
func GenerateSquaredMatrix(n int) [][]float32 {

	dp := make([][]float32, n, n)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = rand.Float32() * float32(rand.Intn(100))
		}
	}
	return dp
}

// PadMatrix will pad a matrix with zeroes until its dimension is 2^n x 2^n.
func PadMatrix(matrix [][]float32) [][]float32 {
	currentDim := len(matrix)
	newDim := int(math.Ceil(math.Log2(float64(currentDim))))

	newMatrix := make([][]float32, newDim, newDim)
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
