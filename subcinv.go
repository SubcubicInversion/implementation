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

	matrix1 := utils.GenerateSquareMatrix(512)

	fmt.Println(algebra.InvertMatrix(matrix1))

	fmt.Printf("Elapsed Time: %f seconds", timer.ElapsedTime(&executionTime).Seconds())
	fmt.Println("")
}
