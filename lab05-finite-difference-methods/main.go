package main

import (
	"fmt"
	"math"
)

// Analytical solution: y(x) = -exp(-x)
func analyticalSolution(x float64) float64 {
	return -math.Exp(-x)
}

// firstOrderDifferenceOperatorForm (explicit, O(h))
//
// АНАЛИТИЧЕСКИЙ АНАЛИЗ:
// Разностная схема: y[i+1] = y[i] * (1 - h)
//
// Порядок аппроксимации: O(h)
// Разложение точного решения в ряд Тейлора:
//
//	y(x+h) = y(x) - h*y(x) + h²/2*y(x) + O(h³)
//	y(x+h) = y(x)*(1 - h + h²/2) + O(h³)
//
// Невязка: ψ = y(x+h) - y[i+1] = y(x)*h²/2 + O(h³)
// Главный член невязки: ψ₀ = -exp(-x)*h²/2
//
// Порядок сходимости: O(h) (первый порядок)
// Устойчивость: схема устойчива при h ≤ 2 (|1-h| ≤ 1)
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

// secondOrderDifferenceOperatorForm (O(h^2))
//
// АНАЛИТИЧЕСКИЙ АНАЛИЗ:
// Разностная схема: y[i+1] = y[i-1] - 2*h*y[i]
//
// Порядок аппроксимации: O(h²)
// Центральная разностная аппроксимация производной:
//
//	(y[i+1] - y[i-1])/(2h) = y'[i] + O(h²)
//
// Для y' = -y получаем:
//
//	y[i+1] = y[i-1] - 2h*y[i]
//
// Разложение в ряд Тейлора:
//
//	y(x+h) - y(x-h) = 2h*y'(x) + h³/3*y'''(x) + O(h⁵)
//	y(x+h) = y(x-h) - 2h*y(x) + h³/3*y(x) + O(h⁵)
//
// Невязка: ψ = h³/3*y(x) + O(h⁵)
// Главный член невязки: ψ₀ = -h³*exp(-x)/3
//
// Порядок сходимости: O(h²) (второй порядок)
// Устойчивость: условно устойчива (λ² + 2hλ - 1 = 0)
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

// forthOrderDifferenceOperatorForm (O(h^4))
//
// АНАЛИТИЧЕСКИЙ АНАЛИЗ:
// Разностная схема: y[i] = (3/(3+h)) * ((-4/3)*h*y[i-1] + (1-h/3)*y[i-2])
//
// Порядок аппроксимации: O(h⁴)
// Схема получена методом неопределенных коэффициентов для достижения
// четвертого порядка аппроксимации.
//
// Разложение точного решения:
//
//	y(x+h) = y(x)*(1 - h + h²/2 - h³/6 + h⁴/24) + O(h⁵)
//
// где используется, что для y'=-y:
//
//	y'=-y, y''=y, y'''=-y, y⁽⁴⁾=y
//
// Коэффициенты схемы подобраны так, чтобы совпадали члены до h³ включительно.
//
// Невязка: ψ = C*h⁵*y(x) + O(h⁶)
// Главный член невязки: ψ₀ = C*h⁵*(-exp(-x))
//
// Порядок сходимости: O(h⁴) (четвертый порядок)
// Устойчивость: проверяется численно
func forthOrderDifferenceOperatorForm(y0 float64, x0, xn float64, h float64) ([]float64, []float64) {
	n := int((xn-x0)/h) + 1
	x := make([]float64, n)
	y := make([]float64, n)
	x[0], y[0] = x0, y0

	if n > 1 {
		// Специальная формула для y[1], обеспечивающая O(h⁴)
		// y[1] = (h² - 6)/(2*h² + 6*h + 6)
		y[1] = (h*h - 6) / (2*h*h + 6*h + 6)
		x[1] = x[0] + h
	}

	// Основная итерационная формула
	// y[i] = (3/(3 + h)) * ((-4/3)*h*y[i-1] + (1-h/3)*y[i-2])
	for i := 2; i < n; i++ {
		y[i] = (3.0 / (3.0 + h)) * ((-4.0/3.0)*h*y[i-1] + (1.0-h/3.0)*y[i-2])
		x[i] = x[i-1] + h
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

func main() {
	y0 := -1.0
	x0 := 0.0
	xn := 2.0
	h1 := 0.1
	h2 := h1 / 2.0

	fmt.Println("Differential equation: y' = -y, y(0) = -1")
	fmt.Printf("Interval: [%.1f, %.1f]\n\n", x0, xn)

	// --- secondOrderDifferenceOperatorForm ---
	xEuler1, yEuler1 := firstOrderDifferenceOperatorForm(y0, x0, xn, h1)
	xEuler2, yEuler2 := firstOrderDifferenceOperatorForm(y0, x0, xn, h2)
	errEuler1 := calculateMaxError(yEuler1, xEuler1)
	errEuler2 := calculateMaxError(yEuler2, xEuler2)
	pEuler := rungeRule(errEuler1, errEuler2, h1/h2)

	fmt.Println("--- Euler Method (k=1, O(h)) ---")
	fmt.Printf("Max error (h=%.2f): %e\n", h1, errEuler1)
	fmt.Printf("Max error (h=%.2f): %e\n", h2, errEuler2)
	fmt.Printf("Numerical convergence order (p): %.4f\n\n", pEuler)

	// --- secondOrderDifferenceOperatorForm ---
	xTrap1, yTrap1 := secondOrderDifferenceOperatorForm(y0, x0, xn, h1)
	xTrap2, yTrap2 := secondOrderDifferenceOperatorForm(y0, x0, xn, h2)
	errTrap1 := calculateMaxError(yTrap1, xTrap1)
	errTrap2 := calculateMaxError(yTrap2, xTrap2)
	pTrap := rungeRule(errTrap1, errTrap2, h1/h2)

	fmt.Println("--- Trapezoidal Method (k=2, O(h^2)) ---")
	fmt.Printf("Max error (h=%.2f): %e\n", h1, errTrap1)
	fmt.Printf("Max error (h=%.2f): %e\n", h2, errTrap2)
	fmt.Printf("Numerical convergence order (p): %.4f\n\n", pTrap)

	// --- Four-Point Method ---
	x4P1, y4P1 := forthOrderDifferenceOperatorForm(y0, x0, xn, h1)
	x4P2, y4P2 := forthOrderDifferenceOperatorForm(y0, x0, xn, h2)
	err4P1 := calculateMaxError(y4P1, x4P1)
	err4P2 := calculateMaxError(y4P2, x4P2)
	p4P := rungeRule(err4P1, err4P2, h1/h2)

	fmt.Println("--- Four-Point Method (k=4, O(h^4)) ---")
	fmt.Printf("Max error (h=%.2f): %e\n", h1, err4P1)
	fmt.Printf("Max error (h=%.2f): %e\n", h2, err4P2)
	fmt.Printf("Numerical convergence order (p): %.4f\n\n", p4P)

	// Generate data for plotting with a smaller step
	plotH := 0.1
	xPlot, yEulerPlot := firstOrderDifferenceOperatorForm(y0, x0, xn, plotH)
	_, yTrapPlot := secondOrderDifferenceOperatorForm(y0, x0, xn, plotH)
	_, y4PPlot := forthOrderDifferenceOperatorForm(y0, x0, xn, plotH)
	yAnalyticalPlot := make([]float64, len(xPlot))
	for i, x := range xPlot {
		yAnalyticalPlot[i] = analyticalSolution(x)
	}

	err := generatePlot(xPlot, yAnalyticalPlot, yEulerPlot, yTrapPlot, y4PPlot)
	if err != nil {
		fmt.Println("Error generating plot:", err)
	} else {
		fmt.Println("Plot successfully generated in plot.html")
	}
}
