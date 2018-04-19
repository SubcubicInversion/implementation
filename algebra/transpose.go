package algebra

import (
	"sync"
)

//transpose an inserted matrix
func transpose(matrixA[][]float32) [][]float32 {
	
	var wg sync.WaitGroup

	length := len(matrixA)

	matrixB := make([][]float32, length)

	for i := range matrixB {
		matrixB[i] = make([]float32, length)
	}

	wg.Add(len(matrixB))
	for row := range matrixA {
		go flip(row, matrixB, matrixA, &wg)
	}

	wg.Wait()
	
	return  matrixB
}

//flip the values
func flip(row int, matrixB[][]float32, matrixA[][]float32, wg *sync.WaitGroup) {
	defer wg.Done()
	for col := range matrixA {
		matrixB[col][row] = matrixA[row][col]
	}
}
