package matrix

import (
	"fmt"
	"math"
)

func CopyMatrix(A [][]float64) [][]float64 {
	n := len(A)
	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, len(A[i]))
		copy(result[i], A[i])
	}
	return result
}

func PrintMatrix(name string, A [][]float64) {
	fmt.Printf("%s =\n", name)
	for i := 0; i < len(A); i++ {
		fmt.Print("  [")
		for j := 0; j < len(A[i]); j++ {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%8.4f", A[i][j])
		}
		fmt.Println("]")
	}
}

// столбцовая норма
func Norm1(A [][]float64) float64 {
	n := len(A)
	m := len(A[0])
	maxSum := 0.0

	for j := 0; j < m; j++ {
		sum := 0.0
		for i := 0; i < n; i++ {
			sum += math.Abs(A[i][j])
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

// норма фробениуса
func Norm2(A [][]float64) float64 {
	sum := 0.0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			sum += A[i][j] * A[i][j]
		}
	}
	return math.Sqrt(sum)
}

// строковая норма
func NormInf(A [][]float64) float64 {
	maxSum := 0.0
	for i := 0; i < len(A); i++ {
		sum := 0.0
		for j := 0; j < len(A[i]); j++ {
			sum += math.Abs(A[i][j])
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

// Обращение матрицы
func invertMatrix(A [][]float64) [][]float64 {
	n := len(A)
	aug := make([][]float64, n) // это матрица размера n×2n, где слева исходная матрица A, а справа единичная матрица I
	for i := 0; i < n; i++ {
		aug[i] = make([]float64, 2*n)
		for j := 0; j < n; j++ {
			aug[i][j] = A[i][j]
		}
		aug[i][n+i] = 1.0
	}

	// Прямой ход
	// После этой процедуры слева от разделителя будет единичная матрица, а справа — обратная матрица
	for i := 0; i < n; i++ {
		pivot := aug[i][i]
		for j := 0; j < 2*n; j++ {
			aug[i][j] /= pivot // Делим всю строку на pivot
		}

		for k := 0; k < n; k++ {
			if k != i {
				factor := aug[k][i]
				for j := 0; j < 2*n; j++ {
					aug[k][j] -= factor * aug[i][j]
				}
			}
		}
	}

	// Извлекаем обратную матрицу
	inv := make([][]float64, n)
	for i := 0; i < n; i++ {
		inv[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			inv[i][j] = aug[i][n+j]
		}
	}

	return inv
}

func ConditionNumber(A [][]float64, norm func([][]float64) float64) float64 {
	invA := invertMatrix(A)
	return norm(A) * norm(invA)
}

// IsDiagonallyDominant проверяет, является ли матрица диагонально доминантной
func IsDiagonallyDominant(A [][]float64) bool {
	n := len(A)
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			if i != j {
				sum += math.Abs(A[i][j])
			}
		}
		if math.Abs(A[i][i]) <= sum {
			return false
		}
	}
	return true
}
