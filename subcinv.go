package main

import (
	"fmt"

	"github.com/SubcubicInversion/implementation/utils"
)

func main() {
	fmt.Println(utils.IsPowerOfTwo(1))
	a := utils.GenerateSquaredMatrix(8)
	fmt.Println(a)
}
