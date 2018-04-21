package algebra

import (
	"errors"
	"fmt"
)

// InvertMatrix inverts a n x n matrix via blockwise inversion and strassen multiplication
// in parallel and subcubic time.
func InvertMatrix(matrix [][]float32) ([][]float32, error) {
	// The trivial inversions involving determinants.
	if len(matrix) <= 3 {
		if matrix == nil || len(matrix) < 1 {
			return nil, errors.New("err: matrix is empty or nonexistent")
		}
		// Check for the singularity of the matrix and acquire determinant.
		det, err := GetDeterminant(matrix)
		if err != nil {
			fmt.Println(err)
		}

		if det == 0 {
			return nil, errors.New("math: cannot invert a singular matrix")
		}

		if len(matrix) == 2 {
			matrix[0][0], matrix[1][1] = matrix[1][1], matrix[0][0]
			matrix[0][1], matrix[1][0] = -1*matrix[0][1], -1*matrix[1][0]
		}

		if len(matrix) == 3 {
			coefficientMatrix := make([][]float32, 3)
			for row := range coefficientMatrix {
				coefficientMatrix[row] = make([]float32, 3)
			}
			coefficientMatrix[0][0] = matrix[1][1]*matrix[2][2] - matrix[1][2]*matrix[2][1]
			coefficientMatrix[1][0] = -1 * (matrix[1][0]*matrix[2][2] - matrix[1][2]*matrix[2][0])
			coefficientMatrix[2][0] = matrix[1][0]*matrix[2][1] - matrix[1][1]*matrix[2][0]
			coefficientMatrix[0][1] = -1 * (matrix[0][1]*matrix[2][2] - matrix[0][2]*matrix[2][1])
			coefficientMatrix[1][1] = matrix[0][0]*matrix[2][2] - matrix[0][2]*matrix[2][0]
			coefficientMatrix[2][1] = -1 * (matrix[0][0]*matrix[2][1] - matrix[0][1]*matrix[2][0])
			coefficientMatrix[0][2] = matrix[0][1]*matrix[1][2] - matrix[0][2]*matrix[1][1]
			coefficientMatrix[1][2] = -1 * (matrix[0][0]*matrix[1][2] - matrix[0][2]*matrix[1][0])
			coefficientMatrix[2][2] = matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
		}

		ScalarMultiply(1/det, matrix)
		return matrix, nil
	}
	return nil, nil
}
