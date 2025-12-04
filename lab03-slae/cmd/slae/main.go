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
	fmt.Println("==============================================================================")
	fmt.Println("РЕШЕНИЕ СИСТЕМЫ ЛИНЕЙНЫХ АЛГЕБРАИЧЕСКИХ УРАВНЕНИЙ")
	fmt.Println("==============================================================================")

	// Примеры матриц для LU, Якоби и Зейделя
	examples := []struct {
		name string
		A    [][]float64
		b    []float64
	}{
		{
			name: "Пример 1 (мой вариант 8)",
			A: [][]float64{
				{1.7, 10, -1.3, 2.1},
				{3.1, 1.7, -2.1, 5.4},
				{3.3, -7.7, 4.4, -5.1},
				{10, -20.1, 20.4, 1.7},
			},
			b: []float64{3.1, 2.1, 1.9, 1.8},
		},
		{
			name: "матрица гилберта",
			A: [][]float64{
				{1.0, float64(1) / 2, 0.3333},
				{0.5, 0.3333, 0.25},
				{0.3333, 0.25, 0.2},
			},
			b: []float64{1, 1, 1},
		},
		{
			name: "Пример 3 (диагонально преобладающая 5x5)",
			A: [][]float64{
				{8, 1, -1, 0, 0},
				{1, 9, 2, -1, 0},
				{-1, 2, 10, 1, -1},
				{0, -1, 1, 7, 2},
				{0, 0, -1, 2, 6},
			},
			b: []float64{8, 11, 11, 9, 7},
		},
		{
			name: "матрица гилберта",
			A: [][]float64{
				{0.0, 1.0, 2.0},
				{1.0, 0.0, 1.0},
				{2.0, 1.0, 0.0},
			},
			b: []float64{1, 1, 1},
		},
	}

	tol := 1e-6
	maxIter := 1000000

	for i, example := range examples {
		fmt.Printf("\n\n")
		fmt.Println("------------------------------------------------------------------------------")
		fmt.Printf("  %s\n", example.name)
		fmt.Println("------------------------------------------------------------------------------")

		matrix.PrintMatrix("Матрица A", example.A)
		vector.PrintVector("Вектор b", example.b)
		fmt.Println()

		// 1. Прямые методы
		fmt.Println("\n1. ПРЯМЫЕ МЕТОДЫ")
		fmt.Println("-----------------------------------")

		// LU-разложение
		fmt.Println("\n1.1. Метод LU-разложения:")
		xLU, err := direct.SolveLu(matrix.CopyMatrix(example.A), vector.CopyVector(example.b))
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			vector.PrintVector("Решение (LU)", xLU)
			fmt.Printf("Невязка: %.2e\n", util.CheckSolution(example.A, xLU, example.b))
		}

		// 2. Итерационные методы
		fmt.Println("\n2. ИТЕРАЦИОННЫЕ МЕТОДЫ")
		fmt.Println("-----------------------------------")

		isDiagDom := matrix.IsDiagonallyDominant(example.A)
		if isDiagDom {
			fmt.Println("\n✓ Матрица диагонально преобладающая - итерационные методы должны сойтись")
		} else {
			fmt.Println("\n✗ Матрица НЕ диагонально преобладающая - итерационные методы могут не сойтись")
		}

		// Метод Якоби
		fmt.Println("\n2.1. Метод Якоби:")
		xJacobi, iterJacobi := iterative.JacobiMethod(matrix.CopyMatrix(example.A), vector.CopyVector(example.b), tol, maxIter)
		if iterJacobi >= maxIter {
			fmt.Println("  ВНИМАНИЕ: Метод не сошелся за максимальное число итераций")
			fmt.Printf("  Итераций: %d (достигнут предел)\n", iterJacobi)
		} else {
			vector.PrintVector("Решение (Якоби)", xJacobi)
			fmt.Printf("Итераций: %d\n", iterJacobi)
			residual := util.CheckSolution(example.A, xJacobi, example.b)
			fmt.Printf("Невязка: %.2e\n", residual)
		}

		// Метод Зейделя
		fmt.Println("\n2.2. Метод Зейделя:")
		xSeidel, iterSeidel := iterative.SeidelMethod(matrix.CopyMatrix(example.A), vector.CopyVector(example.b), tol, maxIter)
		if iterSeidel >= maxIter {
			fmt.Println("  ВНИМАНИЕ: Метод не сошелся за максимальное число итераций")
			fmt.Printf("  Итераций: %d (достигнут предел)\n", iterSeidel)
		} else {
			vector.PrintVector("Решение (Зейделя)", xSeidel)
			fmt.Printf("Итераций: %d\n", iterSeidel)
			residual := util.CheckSolution(example.A, xSeidel, example.b)
			fmt.Printf("Невязка: %.2e\n", residual)
		}

		// Нормы и числа обусловленности (только для первого примера)
		if i == 0 {
			fmt.Println("\n3. НОРМЫ И ЧИСЛА ОБУСЛОВЛЕННОСТИ")
			fmt.Println("-----------------------------------")

			fmt.Printf("\nМатрица A:\n")
			fmt.Printf("  ||A||₁ = %.6f,   cond₁(A) = %.6f\n", matrix.Norm1(example.A), matrix.ConditionNumber(example.A, matrix.Norm1))
			fmt.Printf("  ||A||₂ = %.6f,   cond₂(A) = %.6f\n", matrix.Norm2(example.A), matrix.ConditionNumber(example.A, matrix.Norm2))
			fmt.Printf("  ||A||∞ = %.6f,   cond∞(A) = %.6f\n", matrix.NormInf(example.A), matrix.ConditionNumber(example.A, matrix.NormInf))

			L, U, _ := direct.LuDecomposition(matrix.CopyMatrix(example.A))

			fmt.Printf("\nМатрица L:\n")
			fmt.Printf("  ||L||₁ = %.6f,   cond₁(L) = %.6f\n", matrix.Norm1(L), matrix.ConditionNumber(L, matrix.Norm1))
			fmt.Printf("  ||L||₂ = %.6f,   cond₂(L) = %.6f\n", matrix.Norm2(L), matrix.ConditionNumber(L, matrix.Norm2))
			fmt.Printf("  ||L||∞ = %.6f,   cond∞(L) = %.6f\n", matrix.NormInf(L), matrix.ConditionNumber(L, matrix.NormInf))

			fmt.Printf("\nМатрица U:\n")
			fmt.Printf("  ||U||₁ = %.6f,   cond₁(U) = %.6f\n", matrix.Norm1(U), matrix.ConditionNumber(U, matrix.Norm1))
			fmt.Printf("  ||U||₂ = %.6f,   cond₂(U) = %.6f\n", matrix.Norm2(U), matrix.ConditionNumber(U, matrix.Norm2))
			fmt.Printf("  ||U||∞ = %.6f,   cond∞(U) = %.6f\n", matrix.NormInf(U), matrix.ConditionNumber(U, matrix.NormInf))
		}
	}

	// Метод прогонки
	fmt.Println("\n\n")
	fmt.Println("==============================================================================")
	fmt.Println("4. МЕТОД ПРОГОНКИ ДЛЯ ТРЕХДИАГОНАЛЬНЫХ МАТРИЦ")
	fmt.Println("==============================================================================")

	// Примеры трехдиагональных матриц
	tridiagExamples := []struct {
		name  string
		n     int
		a     []float64 // нижняя диагональ
		bDiag []float64 // главная диагональ
		c     []float64 // верхняя диагональ
		d     []float64 // правая часть
	}{
		{
			name:  "Пример 1 (5x5)",
			n:     5,
			a:     []float64{0, -1, -1, -1, -1},
			bDiag: []float64{4, 4, 4, 4, 4},
			c:     []float64{-1, -1, -1, -1, 0},
			d:     []float64{5, 5, 5, 5, 5},
		},
		{
			name:  "Пример 2 (4x4)",
			n:     4,
			a:     []float64{0, -2, -2, -2},
			bDiag: []float64{10, 10, 10, 10},
			c:     []float64{-1, -1, -1, 0},
			d:     []float64{7, 7, 7, 7},
		},
		{
			name:  "Пример 3 (6x6)",
			n:     6,
			a:     []float64{0, 1, 1, 1, 1, 1},
			bDiag: []float64{-5, -5, -5, -5, -5, -5},
			c:     []float64{1, 1, 1, 1, 1, 0},
			d:     []float64{-3, -3, -3, -3, -3, -3},
		},
	}

	for _, example := range tridiagExamples {
		fmt.Printf("\n\n")
		fmt.Println("------------------------------------------------------------------------------")
		fmt.Printf("  %s\n", example.name)
		fmt.Println("------------------------------------------------------------------------------")

		fmt.Println("\nТрехдиагональная матрица:")
		for i := 0; i < example.n; i++ {
			fmt.Print("  [")
			for j := 0; j < example.n; j++ {
				if i == j {
					fmt.Printf("%8.4f", example.bDiag[i])
				} else if j == i-1 && i > 0 {
					fmt.Printf("%8.4f", example.a[i])
				} else if j == i+1 && i < example.n-1 {
					fmt.Printf("%8.4f", example.c[i])
				} else {
					fmt.Printf("%8.4f", 0.0)
				}
			}
			fmt.Println("]")
		}

		vector.PrintVector("\nВектор правой части", example.d)

		xTridiag := direct.TridiagonalSolve(example.a, example.bDiag, example.c, example.d)
		vector.PrintVector("\nРешение (прогонка)", xTridiag)

		// Проверка для трехдиагональной матрицы
		ATridiag := make([][]float64, example.n)
		for i := 0; i < example.n; i++ {
			ATridiag[i] = make([]float64, example.n)
			ATridiag[i][i] = example.bDiag[i]
			if i > 0 {
				ATridiag[i][i-1] = example.a[i]
			}
			if i < example.n-1 {
				ATridiag[i][i+1] = example.c[i]
			}
		}
		fmt.Printf("Невязка: %.2e\n", util.CheckSolution(ATridiag, xTridiag, example.d))
	}
}
