package main

import (
	"fmt"

	"github.com/SubcubicInversion/implementation/timer"

	"github.com/SubcubicInversion/implementation/algebra"

	"github.com/SubcubicInversion/implementation/utils"
)

func main() {
	var executionTime timer.Timer
	timer.StartTimer(&executionTime, "Execution Test")

	fmt.Println(utils.IsPowerOfTwo(1))

	a := utils.GenerateSquareMatrix(5)
	matrix := [][]float32{
		{2, 7, 3, 5},
		{5, 1, 5, 3},
		{9, 2, 4, 7},
		{1, 4, 6, 2},
	}

	fmt.Println(algebra.InvertMatrix(a))
	fmt.Println(a)

	fmt.Printf("Elapsed Time: %f seconds", timer.ElapsedTime(&executionTime).Seconds())
	fmt.Println("")
}
