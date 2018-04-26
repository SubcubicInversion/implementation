package main

import (
	"fmt"
	"time"

	"github.com/SubcubicInversion/implementation/utils"

	"github.com/SubcubicInversion/implementation/algebra"
)

func main() {
	start := time.Now()

	matrix1 := utils.GenerateSquareMatrix(512)

	fmt.Println(algebra.InvertMatrix(matrix1))

	elapsed := time.Since(start)
	fmt.Printf("Elapsed Time: %s", elapsed)
}
