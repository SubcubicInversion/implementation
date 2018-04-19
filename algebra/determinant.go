package algebra

import "errors"

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

	//TODO: The non-trivial, recursive, parallelized determinant.
	return 0, errors.New("dev: this return is dummy and has not been finalized")
}
