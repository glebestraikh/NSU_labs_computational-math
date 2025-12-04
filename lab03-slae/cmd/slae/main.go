package main

import (
	"fmt"

	"slae/internal/direct"
	"slae/internal/iterative"
	"slae/internal/matrix"
	"slae/internal/util"
	"slae/internal/vector"
)

func main() {
	A := [][]float64{
		{10, -1, 2, 0},
		{-1, 11, -1, 3},
		{2, -1, 10, -1},
		{0, 3, -1, 8},
	}

	b := []float64{6, 25, -11, 15}

	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("РЕШЕНИЕ СИСТЕМЫ ЛИНЕЙНЫХ АЛГЕБРАИЧЕСКИХ УРАВНЕНИЙ")
	fmt.Println("------------------------------------------------------------------------------")

	matrix.PrintMatrix("Матрица A", A)
	vector.PrintVector("Вектор b", b)
	fmt.Println()

	// 1. Прямые методы
	fmt.Println("\n1. ПРЯМЫЕ МЕТОДЫ")
	fmt.Println("------------------------------------------------------------------------------")

	// LU-разложение
	fmt.Println("\n1.1. Метод LU-разложения:")
	xLU, err := direct.SolveLu(matrix.CopyMatrix(A), vector.CopyVector(b))
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		vector.PrintVector("Решение (LU)", xLU)
		fmt.Printf("Невязка: %.2e\n", util.CheckSolution(A, xLU, b))
	}

	// 2. Итерационные методы
	fmt.Println("\n\n2. ИТЕРАЦИОННЫЕ МЕТОДЫ")
	fmt.Println("------------------------------------------------------------------------------")

	tol := 1e-6
	maxIter := 1000

	// Метод Якоби
	fmt.Println("\n2.1. Метод Якоби:")
	xJacobi, iterJacobi := iterative.JacobiMethod(matrix.CopyMatrix(A), vector.CopyVector(b), tol, maxIter)
	vector.PrintVector("Решение (Якоби)", xJacobi)
	fmt.Printf("Итераций: %d\n", iterJacobi)
	fmt.Printf("Невязка: %.2e\n", util.CheckSolution(A, xJacobi, b))

	// Метод Зейделя
	fmt.Println("\n2.2. Метод Зейделя:")
	xSeidel, iterSeidel := iterative.SeidelMethod(matrix.CopyMatrix(A), vector.CopyVector(b), tol, maxIter)
	vector.PrintVector("Решение (Зейделя)", xSeidel)
	fmt.Printf("Итераций: %d\n", iterSeidel)
	fmt.Printf("Невязка: %.2e\n", util.CheckSolution(A, xSeidel, b))

	// Нормы и числа обусловленности
	fmt.Println("\n\n3. НОРМЫ И ЧИСЛА ОБУСЛОВЛЕННОСТИ")
	fmt.Println("------------------------------------------------------------------------------")

	fmt.Printf("\nМатрица A:\n")
	fmt.Printf("  ||A||₁ = %.6f,   cond₁(A) = %.6f\n", matrix.Norm1(A), matrix.ConditionNumber(A, matrix.Norm1))
	fmt.Printf("  ||A||₂ = %.6f,   cond₂(A) = %.6f\n", matrix.Norm2(A), matrix.ConditionNumber(A, matrix.Norm2))
	fmt.Printf("  ||A||∞ = %.6f,   cond∞(A) = %.6f\n", matrix.NormInf(A), matrix.ConditionNumber(A, matrix.NormInf))

	L, U, _ := direct.LuDecomposition(matrix.CopyMatrix(A))

	fmt.Printf("\nМатрица L:\n")
	fmt.Printf("  ||L||₁ = %.6f,   cond₁(L) = %.6f\n", matrix.Norm1(L), matrix.ConditionNumber(L, matrix.Norm1))
	fmt.Printf("  ||L||₂ = %.6f,   cond₂(L) = %.6f\n", matrix.Norm2(L), matrix.ConditionNumber(L, matrix.Norm2))
	fmt.Printf("  ||L||∞ = %.6f,   cond∞(L) = %.6f\n", matrix.NormInf(L), matrix.ConditionNumber(L, matrix.NormInf))

	fmt.Printf("\nМатрица U:\n")
	fmt.Printf("  ||U||₁ = %.6f,   cond₁(U) = %.6f\n", matrix.Norm1(U), matrix.ConditionNumber(U, matrix.Norm1))
	fmt.Printf("  ||U||₂ = %.6f,   cond₂(U) = %.6f\n", matrix.Norm2(U), matrix.ConditionNumber(U, matrix.Norm2))
	fmt.Printf("  ||U||∞ = %.6f,   cond∞(U) = %.6f\n", matrix.NormInf(U), matrix.ConditionNumber(U, matrix.NormInf))

	// 6. Метод прогонки
	fmt.Println("\n\n4. МЕТОД ПРОГОНКИ ДЛЯ ТРЕХДИАГОНАЛЬНОЙ МАТРИЦЫ")
	fmt.Println("------------------------------------------------------------------------------")

	// Пример трехдиагональной матрицы
	n := 5
	a := []float64{0, -1, -1, -1, -1} // нижняя диагональ
	bDiag := []float64{4, 4, 4, 4, 4} // главная диагональ
	c := []float64{-1, -1, -1, -1, 0} // верхняя диагональ
	d := []float64{5, 5, 5, 5, 5}     // правая часть

	fmt.Println("Трехдиагональная матрица:")
	for i := 0; i < n; i++ {
		fmt.Print("  [")
		for j := 0; j < n; j++ {
			if i == j {
				fmt.Printf("%8.4f", bDiag[i])
			} else if j == i-1 && i > 0 {
				fmt.Printf("%8.4f", a[i])
			} else if j == i+1 && i < n-1 {
				fmt.Printf("%8.4f", c[i])
			} else {
				fmt.Printf("%8.4f", 0.0)
			}
		}
		fmt.Println("]")
	}

	xTridiag := direct.TridiagonalSolve(a, bDiag, c, d)
	vector.PrintVector("\nРешение (прогонка)", xTridiag)

	// Проверка для трехдиагональной матрицы
	ATridiag := make([][]float64, n)
	for i := 0; i < n; i++ {
		ATridiag[i] = make([]float64, n)
		ATridiag[i][i] = bDiag[i]
		if i > 0 {
			ATridiag[i][i-1] = a[i]
		}
		if i < n-1 {
			ATridiag[i][i+1] = c[i]
		}
	}
	fmt.Printf("Невязка: %.2e\n", util.CheckSolution(ATridiag, xTridiag, d))

	fmt.Println("------------------------------------------------------------------------------")
}
