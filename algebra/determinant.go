package algebra

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

// GetDeterminant returns the determinant of an n x n matrix
func GetDeterminant(matrix [][]float32) (float32, error) {
	if matrix == nil || len(matrix) < 1 {
		return 0, errors.New("err: matrix is empty or nonexistent")
	}

	if len(matrix) == 1 {
		return matrix[0][0], nil
	}

	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0], nil
	}

	if len(matrix) == 3 {
		return matrix[0][0]*matrix[1][1]*matrix[2][2] - matrix[0][0]*matrix[1][2]*matrix[2][1] - matrix[0][1]*matrix[1][0]*matrix[2][2] + matrix[0][1]*matrix[1][2]*matrix[2][0] + matrix[0][2]*matrix[1][0]*matrix[2][1] - matrix[0][2]*matrix[1][1]*matrix[2][0], nil
	}

	// Instantiate the goroutine pool.
	var wg sync.WaitGroup

	// Instantiate a matrix to return
	determinants := make([]float32, len(matrix))

	for i := range matrix {
		wg.Add(1)
		go func() {
			innerMatrix := extractInnerMatrixIgnoringGivenRowAndCol(0, i, matrix)
			det, err := GetDeterminant(innerMatrix)

			if err != nil {
				fmt.Println(err)
			} else {
				determinants[i] = det * float32(math.Pow(float64(-1), float64(i)))
			}
			wg.Done()
		}()
	}

	// Wait until all goroutines have finished before returning
	wg.Wait()

	var response float32

	for i := range determinants {
		response += determinants[i] * matrix[0][i]
	}

	return response, nil
}

func extractInnerMatrixIgnoringGivenRowAndCol(elementRow int, elementCol int, matrix [][]float32) [][]float32 {
	innerMatrix := make([][]float32, len(matrix)-1)

	for row := range matrix {
		if elementRow == row {
			continue
		}

		realRow := row
		if elementRow <= row {
			realRow = row - 1
		}

		innerMatrix[realRow] = make([]float32, len(matrix)-1)

		for col := range matrix[row] {
			if elementCol == col {
				continue
			}

			realCol := col
			if elementCol <= col {
				realCol = col - 1
			}

			innerMatrix[realRow][realCol] = matrix[row][col]
		}
	}

	return innerMatrix
}
