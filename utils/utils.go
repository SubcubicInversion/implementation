package utils

import "math/rand"

// GenerateSquaredMatrix generates an n x n square matrix with floats as elements
func GenerateSquaredMatrix(n int) [][]float32 {

	dp := make([][]float32, n)
	for i := range dp {
		dp[i] = make([]float32, n)

		for j := range dp[i] {
			dp[i][j] = rand.Float32() * float32(rand.Intn(100))
		}
	}
	return dp
}
