package main

import "math/rand"

func generateSquaredMatrix(n int) ([][]float32) {

	dp := make([][]float32, n)
	for i := range dp {
		dp[i] = make([]float32, n)

		for j := range dp[i] {
			dp[i][j] = rand.Float32() * float32(rand.Intn(100))
		}
	}

	return dp
}
