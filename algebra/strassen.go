package algebra

import (
	"sync"

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

	matrixC := utils.MakeMatrix(len(matrixA), len(matrixA))

	a11, a12, a21, a22 := utils.GetMatrixSubBlocks(matrixA)
	b11, b12, b21, b22 := utils.GetMatrixSubBlocks(matrixB)

	c11, c12, c21, c22 := utils.GetMatrixSubBlocks(matrixC)

	// Parallelize the computation of the seven intermediate matrices
	var wg sync.WaitGroup

	var m1, m2, m3, m4, m5, m6, m7 [][]float32
	wg.Add(7)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		m1 = StrassenMultiply(AddMatrices(a11, a22), AddMatrices(b11, b22))
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		m2 = StrassenMultiply(AddMatrices(a21, a22), b11)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		m3 = StrassenMultiply(a11, AddMatrices(b12, ScalarMultiply(-1, b22)))
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		m4 = StrassenMultiply(a22, AddMatrices(b21, ScalarMultiply(-1, b11)))
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		m5 = StrassenMultiply(AddMatrices(a11, a22), b22)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		m6 = StrassenMultiply(AddMatrices(a21, ScalarMultiply(-1, a11)), AddMatrices(b11, b12))
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		m7 = StrassenMultiply(AddMatrices(a12, ScalarMultiply(-1, a22)), AddMatrices(b21, b22))
	}(&wg)
	wg.Wait()

	wg.Add(4)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// TODO: Here, I am pointing these things to new slices. This will not
		// provide a solution. Rather, I need to fill these slices with the result.
		c11 = AddMatrices(AddMatrices(AddMatrices(m1, m4), ScalarMultiply(-1, m5)), m7)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		c12 = AddMatrices(m3, m5)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		c21 = AddMatrices(m2, m4)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		c22 = AddMatrices(AddMatrices(AddMatrices(m1, ScalarMultiply(-1, m2)), m3), m6)
	}(&wg)
	wg.Wait()

	// m1 := StrassenMultiply(AddMatrices(a11, a22), AddMatrices(b11, b22))
	// m2 := StrassenMultiply(AddMatrices(a21, a22), b11)
	// m3 := StrassenMultiply(a11, AddMatrices(b12, ScalarMultiply(-1, b22)))
	// m4 := StrassenMultiply(a22, AddMatrices(b21, ScalarMultiply(-1, b11)))
	// m5 := StrassenMultiply(AddMatrices(a11, a22), b22)
	// m6 := StrassenMultiply(AddMatrices(a21, ScalarMultiply(-1, a11)), AddMatrices(b11, b12))
	// m7 := StrassenMultiply(AddMatrices(a12, ScalarMultiply(-1, a22)), AddMatrices(b21, b22))

	// Less parallelized
	// c11 = StrassenMultiply(AddMatrices(a11, a22), AddMatrices(b11, b22)) + StrassenMultiply(a22, AddMatrices(b21, ScalarMultiply(-1, b11))) - StrassenMultiply(AddMatrices(a11, a22), b22) + StrassenMultiply(AddMatrices(a12, ScalarMultiply(-1, a22)), AddMatrices(b21, b22))
	// c12 = StrassenMultiply(a11, AddMatrices(b12, ScalarMultiply(-1, b22)))
	// c21 =

	return matrixC
}
