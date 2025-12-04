package iterative

import (
	"math"

	"slae/internal/vector"
)

func SeidelMethod(A [][]float64, b []float64, tol float64, maxIter int) ([]float64, int) {
	n := len(b)
	x := make([]float64, n)

	for iter := 0; iter < maxIter; iter++ {
		xOld := vector.CopyVector(x) // Для вычисления нового x_i используются только значения из предыдущей итерации (x[j])

		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if i != j {
					sum += A[i][j] * x[j]
				}
			}
			x[i] = (b[i] - sum) / A[i][i]
		}

		// Проверка сходимости
		diff := 0.0
		for i := 0; i < n; i++ {
			diff += math.Abs(x[i] - xOld[i])
		}

		if diff < tol {
			return x, iter + 1
		}
	}

	return x, maxIter
}
