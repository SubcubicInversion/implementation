package main

import (
	"fmt"

	"github.com/SubcubicInversion/implementation/timer"
	"github.com/SubcubicInversion/implementation/utils"
)

func main() {
	var executionTime timer.Timer
	timer.StartTimer(&executionTime, "Execution Test")

	fmt.Println(utils.IsPowerOfTwo(1))
	a := utils.GenerateSquareMatrix(8)
	fmt.Println(a)

	timer.StopTimer(&executionTime, "Execution Test")
	fmt.Printf("Elapsed Time: %f seconds", timer.ElapsedTime(&executionTime).Seconds())
	fmt.Println("")
}
