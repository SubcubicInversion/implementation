package algebra

import (
	"sync"
)

// Transpose transposes a matrix
func Transpose(matrix [][]float32) [][]float32 {
	if matrix == nil {
		return nil
	}

	var wg sync.WaitGroup

	length := len(matrix)
	transpose := make([][]float32, length)
	for i := range transpose {
		transpose[i] = make([]float32, length)
	}

	wg.Add(length)
	for row := range matrix {
		go flip(row, transpose, matrix, &wg)
	}
	wg.Wait()

	return transpose
}

// flip flips the rows and performs the transpose in parallel
func flip(row int, transpose [][]float32, matrix [][]float32, wg *sync.WaitGroup) {
	for col := range matrix {
		transpose[col][row] = matrix[row][col]
	}
	wg.Done()
}
