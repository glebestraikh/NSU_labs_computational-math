package iterative

import "math"

func JacobiMethod(A [][]float64, b []float64, tol float64, maxIter int) ([]float64, int) {
	n := len(b)
	x := make([]float64, n)
	xNew := make([]float64, n)

	for iter := 0; iter < maxIter; iter++ {
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if i != j {
					sum += A[i][j] * x[j]
				}
			}
			xNew[i] = (b[i] - sum) / A[i][i] // для метода Якоби все новые значения вычисляются исключительно из старого вектора x, поэтому нужен отдельный xNew
		}

		// Проверка сходимости
		diff := 0.0
		for i := 0; i < n; i++ {
			diff += math.Abs(xNew[i] - x[i])
		}

		copy(x, xNew)

		if diff < tol {
			return x, iter + 1
		}
	}

	return x, maxIter
}
