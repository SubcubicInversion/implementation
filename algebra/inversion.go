package algebra

import (
	"errors"
	"fmt"
	"sync"

	"github.com/SubcubicInversion/implementation/utils"
)

// InvertMatrix inverts a n x n matrix via blockwise inversion and strassen multiplication
// in parallel and subcubic time.
func InvertMatrix(matrix [][]float32) ([][]float32, error) {
	// The trivial inversions involving determinants.
	dim := len(matrix)
	if dim <= 3 {
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

		if dim == 2 {
			matrix[0][0], matrix[1][1] = matrix[1][1], matrix[0][0]
			matrix[0][1], matrix[1][0] = -1*matrix[0][1], -1*matrix[1][0]
		}

		if dim == 3 {
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

			matrix = coefficientMatrix
		}

		return ScalarMultiply(1/det, matrix), nil
	}

	// Create four matrix sub-blocks
	a := utils.SliceMatrix(matrix, 0, dim/2, 0, dim/2)
	b := utils.SliceMatrix(matrix, 0, dim/2, dim/2, dim)
	c := utils.SliceMatrix(matrix, dim/2, dim, 0, dim/2)
	d := utils.SliceMatrix(matrix, dim/2, dim, dim/2, dim)

	aInv, err := InvertMatrix(a)
	if err != nil {
		fmt.Println(err)
	}

	caInv := StrassenMultiply(c, aInv)
	dInv, err := InvertMatrix(AddMatrices(d, StrassenMultiply(ScalarMultiply(-1, caInv), b)))
	abdInv := StrassenMultiply(aInv, StrassenMultiply(b, dInv))

	var wg sync.WaitGroup

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// TODO: Make this actually work because Go isn't this nice.
		a = AddMatrices(aInv, StrassenMultiply(abdInv, caInv))
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		c = ScalarMultiply(-1, StrassenMultiply(dInv, caInv))
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		d = ScalarMultiply(-1, abdInv)
	}(&wg)

	return matrix, nil
}
