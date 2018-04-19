package main

import (
	"fmt"

	"github.com/SubcubicInversion/implementation/utils"
	"github.com/SubcubicInversion/implementation/timer"
	"github.com/SubcubicInversion/implementation/algebra"
)

func main() {
	var executionTime timer.Timer
	timer.StartTimer(&executionTime, "Execution Test")

	fmt.Println(utils.IsPowerOfTwo(1))
	a := utils.GenerateSquaredMatrix(5)
	matrix := [][]float32{
		{2, 7, 3, 5},
		{5, 1, 5, 3},
		{9, 2, 4, 7},
		{1, 4, 6, 2},
	}

	det, err := algebra.GetDeterminant(matrix)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(det)
	}


	fmt.Println(a)

	fmt.Printf("Elapsed Time: %f seconds", timer.ElapsedTime(&executionTime).Seconds())
	fmt.Println("")
}
