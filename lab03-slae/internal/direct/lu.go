package direct

import (
	"fmt"
	"slae/internal/matrix"
)

func LuDecomposition(A [][]float64) ([][]float64, [][]float64, error) {
	n := len(A)
	L := make([][]float64, n)
	U := matrix.CopyMatrix(A)

	for i := 0; i < n; i++ {
		L[i] = make([]float64, n)
		L[i][i] = 1.0
	}

	for k := 0; k < n-1; k++ { // индекс пивота
		for i := k + 1; i < n; i++ {
			L[i][k] = U[i][k] / U[k][k] // коэффициент, который нужен для обнуления элементов ниже диагонали в столбце k
			for j := k; j < n; j++ {
				U[i][j] -= L[i][k] * U[k][j] // Обновляем строки матрицы U, вычитая соответствующую часть из верхней треугольной матрицы
			}
		}
	}

	return L, U, nil
}

func SolveLu(A [][]float64, b []float64) ([]float64, error) {
	L, U, err := LuDecomposition(A)
	if err != nil {
		return nil, err
	}

	fmt.Println("\n3.1. LU-разложение:")
	matrix.PrintMatrix("L", L)
	matrix.PrintMatrix("U", U)
	fmt.Println()

	n := len(b)

	// Прямая подстановка: Ly = b
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < i; j++ {
			sum += L[i][j] * y[j] // для каждой строки i находим вклад уже найденных элементов y[0..i-1]
		}
		y[i] = b[i] - sum
	}

	// Обратная подстановка: Ux = y
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += U[i][j] * x[j]
		}
		x[i] = (y[i] - sum) / U[i][i]
	}

	return x, nil
}
