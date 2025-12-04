package util

import "math"

func CheckSolution(A [][]float64, x, b []float64) float64 {
	n := len(b)
	residual := make([]float64, n)

	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += A[i][j] * x[j]
		}
		residual[i] = sum - b[i]
	}

	// Вычисляем норму невязки
	norm := 0.0
	for i := 0; i < n; i++ {
		norm += residual[i] * residual[i]
	}
	return math.Sqrt(norm)
}
