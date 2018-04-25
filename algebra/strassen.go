package algebra

import (
	"github.com/SubcubicInversion/implementation/utils"
)

// StrassenMultiply will use Strassen's algorithm to multiply two matrices in subcubic time (O(n^log_2(7)), approximately O(n^2.807)).
func StrassenMultiply(matrixA [][]float32, matrixB [][]float32) [][]float32 {
	if !utils.IsPowerOfTwo(len(matrixA)) {
		matrixA = utils.PadMatrix(matrixA)
	}
	if !utils.IsPowerOfTwo(len(matrixB)) {
		matrixB = utils.PadMatrix(matrixB)
	}

	a11, a12, a21, a22 := utils.GetMatrixSubBlocks(matrixA)
	b11, b12, b21, b22 := utils.GetMatrixSubBlocks(matrixB)

	return nil
}
