package main

import (
	"fmt"
	"math"
)

// Analytical solution: y(x) = -exp(-x)
func analyticalSolution(x float64) float64 {
	return -math.Exp(-x)
}

func firstOrderDifferenceOperatorForm(y0 float64, x0, xn float64, h float64) ([]float64, []float64) {
	n := int((xn-x0)/h) + 1
	x := make([]float64, n)
	y := make([]float64, n)
	x[0], y[0] = x0, y0

	// y[i+1] = y[i] * (1 - h)
	for i := 0; i < n-1; i++ {
		y[i+1] = y[i] * (1 - h)
		x[i+1] = x[i] + h
	}
	return x, y
}

func secondOrderDifferenceOperatorForm(y0 float64, x0, xn float64, h float64) ([]float64, []float64) {
	n := int((xn-x0)/h) + 1
	x := make([]float64, n)
	y := make([]float64, n)
	x[0], y[0] = x0, y0

	if n > 1 {
		y[1] = y[0] * (1 - h)
		x[1] = x[0] + h
	}

	// y[i+1] = y[i-1] - 2*h*y[i]
	for i := 1; i < n-1; i++ {
		y[i+1] = y[i-1] - 2*h*y[i]
		x[i+1] = x[i] + h
	}
	return x, y
}

// Calculate max error between numerical and analytical solution
func calculateMaxError(yNum, xNum []float64) float64 {
	maxErr := 0.0
	for i := range yNum {
		err := math.Abs(yNum[i] - analyticalSolution(xNum[i]))
		if err > maxErr {
			maxErr = err
		}
	}
	return maxErr
}

// Runge's rule for convergence order
func rungeRule(err1, err2, r float64) float64 {
	return math.Log(err1/err2) / math.Log(r)
}

// Анализ устойчивости схемы первого порядка (Эйлер)
func analyzeStabilityFirstOrder(h float64) {
	fmt.Println("=== АНАЛИЗ УСТОЙЧИВОСТИ СХЕМЫ k=1 (Эйлер) ===")
	fmt.Println("Разностная схема: y[i+1] = y[i] * (1 - h)")
	fmt.Println()
	fmt.Println("Коэффициент усиления: λ = 1 - h")
	fmt.Printf("При h = %.4f: λ = %.4f\n", h, 1-h)
	fmt.Println()
	fmt.Println("Условие устойчивости по теореме Лакса: |λ| ≤ 1")
	fmt.Println("  |1 - h| ≤ 1")
	fmt.Println("  -1 ≤ 1 - h ≤ 1")
	fmt.Println("  0 ≤ h ≤ 2")
	fmt.Println()
	if h > 0 && h <= 2 {
		fmt.Printf("✓ Схема УСТОЙЧИВА при h = %.4f\n", h)
	} else {
		fmt.Printf("✗ Схема НЕУСТОЙЧИВА при h = %.4f\n", h)
	}
	fmt.Println()
}

// Анализ устойчивости схемы второго порядка
func analyzeStabilitySecondOrder(h float64) {
	fmt.Println("=== АНАЛИЗ УСТОЙЧИВОСТИ СХЕМЫ k=2 ===")
	fmt.Println("Разностная схема: y[i+1] = y[i-1] - 2*h*y[i]")
	fmt.Println()
	fmt.Println("Характеристическое уравнение: λ² + 2h*λ - 1 = 0")
	fmt.Println()

	// Решаем характеристическое уравнение
	discriminant := 4*h*h + 4
	lambda1 := (-2*h + math.Sqrt(discriminant)) / 2
	lambda2 := (-2*h - math.Sqrt(discriminant)) / 2

	fmt.Printf("Корни характеристического уравнения при h = %.4f:\n", h)
	fmt.Printf("  λ₁ = %.6f, |λ₁| = %.6f\n", lambda1, math.Abs(lambda1))
	fmt.Printf("  λ₂ = %.6f, |λ₂| = %.6f\n", lambda2, math.Abs(lambda2))
	fmt.Println()
	fmt.Println("Условие устойчивости: |λᵢ| ≤ 1 для всех i")

	stable := math.Abs(lambda1) <= 1.0001 && math.Abs(lambda2) <= 1.0001 // небольшая погрешность
	if stable {
		fmt.Println("✓ Схема УСТОЙЧИВА (оба корня удовлетворяют условию)")
	} else {
		fmt.Println("✗ Схема УСЛОВНО УСТОЙЧИВА или НЕУСТОЙЧИВА")
		fmt.Println("  (один из корней по модулю больше 1)")
	}
	fmt.Println()
}

// Анализ устойчивости схемы четвертого порядка
func analyzeStabilityFourthOrder() {

}

// Анализ порядка аппроксимации и невязки для схемы первого порядка
func analyzeApproximationFirstOrder() {
	fmt.Println("=== АНАЛИЗ ПОРЯДКА АППРОКСИМАЦИИ И НЕВЯЗКИ k=1 ===")
	fmt.Println()
	fmt.Println("Разностная схема: y[i+1] = y[i] * (1 - h)")
	fmt.Println()
	fmt.Println("Разложение точного решения в ряд Тейлора:")
	fmt.Println("  y(x+h) = y(x) + h*y'(x) + (h²/2)*y''(x) + O(h³)")
	fmt.Println()
	fmt.Println("Для уравнения y' = -y:")
	fmt.Println("  y'(x) = -y(x)")
	fmt.Println("  y''(x) = y(x)")
	fmt.Println()
	fmt.Println("Подставляем:")
	fmt.Println("  y(x+h) = y(x) - h*y(x) + (h²/2)*y(x) + O(h³)")
	fmt.Println("  y(x+h) = y(x) * (1 - h + h²/2) + O(h³)")
	fmt.Println()
	fmt.Println("Невязка ψ = y(x+h) - y[i+1]:")
	fmt.Println("  ψ = y(x) * (1 - h + h²/2) - y(x) * (1 - h) + O(h³)")
	fmt.Println("  ψ = y(x) * h²/2 + O(h³)")
	fmt.Println()
	fmt.Println("ГЛАВНЫЙ ЧЛЕН НЕВЯЗКИ: ψ₀ = -exp(-x) * h²/2")
	fmt.Println("ПОРЯДОК АППРОКСИМАЦИИ: O(h) - первый порядок")
	fmt.Println()
}

// Анализ порядка аппроксимации и невязки для схемы второго порядка
func analyzeApproximationSecondOrder() {
	fmt.Println("=== АНАЛИЗ ПОРЯДКА АППРОКСИМАЦИИ И НЕВЯЗКИ k=2 ===")
	fmt.Println()
	fmt.Println("Разностная схема: y[i+1] = y[i-1] - 2*h*y[i]")
	fmt.Println()
	fmt.Println("Центральная разностная аппроксимация производной:")
	fmt.Println("  (y[i+1] - y[i-1])/(2h) = y'[i] + O(h²)")
	fmt.Println()
	fmt.Println("Для y' = -y:")
	fmt.Println("  (y[i+1] - y[i-1])/(2h) = -y[i]")
	fmt.Println("  y[i+1] = y[i-1] - 2h*y[i]")
	fmt.Println()
	fmt.Println("Разложение в ряд Тейлора:")
	fmt.Println("  y(x+h) = y(x) + h*y'(x) + (h²/2)*y''(x) + (h³/6)*y'''(x) + O(h⁴)")
	fmt.Println("  y(x-h) = y(x) - h*y'(x) + (h²/2)*y''(x) - (h³/6)*y'''(x) + O(h⁴)")
	fmt.Println()
	fmt.Println("Вычитаем:")
	fmt.Println("  y(x+h) - y(x-h) = 2h*y'(x) + (h³/3)*y'''(x) + O(h⁵)")
	fmt.Println()
	fmt.Println("Для y' = -y, y''' = y:")
	fmt.Println("  y(x+h) = y(x-h) - 2h*y(x) + (h³/3)*y(x) + O(h⁵)")
	fmt.Println()
	fmt.Println("Невязка ψ = y(x+h) - [y(x-h) - 2h*y(x)]:")
	fmt.Println("  ψ = (h³/3)*y(x) + O(h⁵)")
	fmt.Println()
	fmt.Println("ГЛАВНЫЙ ЧЛЕН НЕВЯЗКИ: ψ₀ = -(h³/3) * exp(-x)")
	fmt.Println("ПОРЯДОК АППРОКСИМАЦИИ: O(h²) - второй порядок")
	fmt.Println()
}

// Анализ порядка аппроксимации и невязки для схемы четвертого порядка
func analyzeApproximationFourthOrder() {
}

func main() {
	y0 := -1.0
	x0 := 0.0
	xn := 2.0
	h1 := 0.1
	h2 := h1 / 2.0

	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║  ЛАБОРАТОРНАЯ РАБОТА №5: МЕТОДЫ КОНЕЧНЫХ РАЗНОСТЕЙ                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Дифференциальное уравнение: y' = -y, y(0) = -1")
	fmt.Printf("Интервал: [%.1f, %.1f]\n", x0, xn)
	fmt.Printf("Шаги: h₁ = %.4f, h₂ = %.4f\n", h1, h2)
	fmt.Println()
	fmt.Println("════════════════════════════════════════════════════════════════════")
	fmt.Println()

	// ========== ТЕОРЕТИЧЕСКИЙ АНАЛИЗ ==========
	// 1. Анализ порядка аппроксимации и невязки
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println("1. ПОРЯДОК АППРОКСИМАЦИИ И ГЛАВНЫЕ ЧЛЕНЫ НЕВЯЗКИ")
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()

	analyzeApproximationFirstOrder()
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()

	analyzeApproximationSecondOrder()
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()

	analyzeApproximationFourthOrder()
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()

	// 2. Анализ устойчивости
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println("2. ПРОВЕРКА УСТОЙЧИВОСТИ ПО ТЕОРЕМЕ ЛАКСА")
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()

	analyzeStabilityFirstOrder(h1)
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()

	analyzeStabilitySecondOrder(h1)
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()

	analyzeStabilityFourthOrder()
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()

	// 3. Анализ порядка сходимости
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println("3. ПОРЯДОК СХОДИМОСТИ РАЗНОСТНЫХ СХЕМ")
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()

	fmt.Println("════════════════════════════════════════════════════════════════════")
	fmt.Println()

	fmt.Println()
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println("4. ПРАВИЛО РУНГЕ - ЧИСЛЕННАЯ ПРОВЕРКА ПОРЯДКА СХОДИМОСТИ")
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()

	// --- firstOrderDifferenceOperatorForm (k=1) ---
	xEuler1, yEuler1 := firstOrderDifferenceOperatorForm(y0, x0, xn, h1)
	xEuler2, yEuler2 := firstOrderDifferenceOperatorForm(y0, x0, xn, h2)
	errEuler1 := calculateMaxError(yEuler1, xEuler1)
	errEuler2 := calculateMaxError(yEuler2, xEuler2)
	pEuler := rungeRule(errEuler1, errEuler2, h1/h2)

	fmt.Println("┌─ Метод Эйлера (k=1, O(h)) ─────────────────────────────────────┐")
	fmt.Printf("│ Максимальная ошибка при h = %.4f:  %e\n", h1, errEuler1)
	fmt.Printf("│ Максимальная ошибка при h = %.4f:  %e\n", h2, errEuler2)
	fmt.Printf("│ Численный порядок сходимости p:         %.4f\n", pEuler)
	fmt.Println("│")
	if math.Abs(pEuler-1.0) < 0.2 {
		fmt.Println("│ ✓ Численный порядок сходимости соответствует теоретическому O(h)")
	} else {
		fmt.Println("│ ⚠ Численный порядок сходимости отличается от теоретического")
	}
	fmt.Println("└────────────────────────────────────────────────────────────────┘")
	fmt.Println()

	// --- secondOrderDifferenceOperatorForm (k=2) ---
	xTrap1, yTrap1 := secondOrderDifferenceOperatorForm(y0, x0, xn, h1)
	xTrap2, yTrap2 := secondOrderDifferenceOperatorForm(y0, x0, xn, h2)
	errTrap1 := calculateMaxError(yTrap1, xTrap1)
	errTrap2 := calculateMaxError(yTrap2, xTrap2)
	pTrap := rungeRule(errTrap1, errTrap2, h1/h2)

	fmt.Println("┌─ Метод второго порядка (k=2, O(h²)) ───────────────────────────┐")
	fmt.Printf("│ Максимальная ошибка при h = %.4f:  %e\n", h1, errTrap1)
	fmt.Printf("│ Максимальная ошибка при h = %.4f:  %e\n", h2, errTrap2)
	fmt.Printf("│ Численный порядок сходимости p:         %.4f\n", pTrap)
	fmt.Println("│")
	if math.Abs(pTrap-2.0) < 0.2 {
		fmt.Println("│ ✓ Численный порядок сходимости соответствует теоретическому O(h²)")
	} else {
		fmt.Println("│ ⚠ Численный порядок сходимости отличается от теоретического")
	}
	fmt.Println("└────────────────────────────────────────────────────────────────┘")
	fmt.Println()

	// --- forthOrderDifferenceOperatorForm (k=4) ---

	fmt.Println("════════════════════════════════════════════════════════════════════")
	fmt.Println()

	// Generate data for plotting with a smaller step
	plotH := 0.1
	xPlot, yEulerPlot := firstOrderDifferenceOperatorForm(y0, x0, xn, plotH)
	_, ySecondPlot := secondOrderDifferenceOperatorForm(y0, x0, xn, plotH)
	yAnalyticalPlot := make([]float64, len(xPlot))
	for i, x := range xPlot {
		yAnalyticalPlot[i] = analyticalSolution(x)
	}

	err := generatePlot(xPlot, yAnalyticalPlot, yEulerPlot, ySecondPlot, ySecondPlot)
	if err != nil {
		fmt.Println("Error generating plot:", err)
	} else {
		fmt.Println("Plot successfully generated in plot.html")
	}
}
