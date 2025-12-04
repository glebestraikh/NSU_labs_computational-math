package main

import (
	"fmt"
	"math"
)

// Analytical solution: y(x) = -exp(-x)
func analyticalSolution(x float64) float64 {
	return -math.Exp(-x)
}

// f(x, y) = y' = -y
func f(x, y float64) float64 {
	return -y
}

// Euler method (explicit, O(h))
func eulerMethod(y0 float64, x0, xn float64, h float64) ([]float64, []float64) {
	n := int((xn-x0)/h) + 1
	x := make([]float64, n)
	y := make([]float64, n)
	x[0], y[0] = x0, y0

	for i := 0; i < n-1; i++ {
		y[i+1] = y[i] + h*f(x[i], y[i])
		x[i+1] = x[i] + h
	}
	return x, y
}

// Two-point scheme (implicit trapezoidal, O(h^2))
func trapezoidalMethod(y0 float64, x0, xn float64, h float64) ([]float64, []float64) {
	n := int((xn-x0)/h) + 1
	x := make([]float64, n)
	y := make([]float64, n)
	x[0], y[0] = x0, y0

	for i := 0; i < n-1; i++ {
		// y[i+1] = y[i] + h/2 * (-y[i] - y[i+1])
		// y[i+1] * (1 + h/2) = y[i] * (1 - h/2)
		y[i+1] = y[i] * (1 - h/2) / (1 + h/2)
		x[i+1] = x[i] + h
	}
	return x, y
}

// Four-point scheme (O(h^4)), adapted from user's Python code for y' = -y
func fourPointMethod(y0 float64, x0, xn float64, h float64) ([]float64, []float64) {
	n := int((xn-x0)/h) + 1
	x := make([]float64, n)
	y := make([]float64, n)
	x[0], y[0] = x0, y0

	if n > 1 {
		// Formula for y[1] from Python code `getY1(h)` adapted for y'=-y
		// The original python code was for y'=y, so we use -h.
		// y[1] = y0 * (6 - (-h)*(-h)) / (2 * (3 + 3*(-h) + (-h)*(-h)))
		y[1] = y0 * (6 - h*h) / (2 * (3 - 3*h + h*h))
		x[1] = x[0] + h
	}

	if n > 2 {
		// Formula for y[i+1] from Python code `appForK4(y1, y0, h)` adapted for y'=-y
		// The original python code was for y'=y, so we use -h.
		// y[i+1] = (-4*(-h)*y[i] - ((-h)-3)*y[i-1]) / ((-h)+3)
		// y[i+1] = (4*h*y[i] + (h+3)*y[i-1]) / (3-h)
		for i := 1; i < n-1; i++ {
			y[i+1] = (4*h*y[i] + (h+3)*y[i-1]) / (3 - h)
			x[i+1] = x[i] + h
		}
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

	// --- Euler Method ---
	xEuler1, yEuler1 := eulerMethod(y0, x0, xn, h1)
	xEuler2, yEuler2 := eulerMethod(y0, x0, xn, h2)
	errEuler1 := calculateMaxError(yEuler1, xEuler1)
	errEuler2 := calculateMaxError(yEuler2, xEuler2)
	pEuler := rungeRule(errEuler1, errEuler2, h1/h2)

	fmt.Println("--- Euler Method (k=1, O(h)) ---")
	fmt.Printf("Max error (h=%.2f): %e\n", h1, errEuler1)
	fmt.Printf("Max error (h=%.2f): %e\n", h2, errEuler2)
	fmt.Printf("Numerical convergence order (p): %.4f\n\n", pEuler)

	// --- Trapezoidal Method ---
	xTrap1, yTrap1 := trapezoidalMethod(y0, x0, xn, h1)
	xTrap2, yTrap2 := trapezoidalMethod(y0, x0, xn, h2)
	errTrap1 := calculateMaxError(yTrap1, xTrap1)
	errTrap2 := calculateMaxError(yTrap2, xTrap2)
	pTrap := rungeRule(errTrap1, errTrap2, h1/h2)

	fmt.Println("--- Trapezoidal Method (k=2, O(h^2)) ---")
	fmt.Printf("Max error (h=%.2f): %e\n", h1, errTrap1)
	fmt.Printf("Max error (h=%.2f): %e\n", h2, errTrap2)
	fmt.Printf("Numerical convergence order (p): %.4f\n\n", pTrap)

	// --- Four-Point Method ---
	x4P_1, y4P_1 := fourPointMethod(y0, x0, xn, h1)
	x4P_2, y4P_2 := fourPointMethod(y0, x0, xn, h2)
	err4P_1 := calculateMaxError(y4P_1, x4P_1)
	err4P_2 := calculateMaxError(y4P_2, x4P_2)
	p4P := rungeRule(err4P_1, err4P_2, h1/h2)

	fmt.Println("--- Four-Point Method (k=4, O(h^4)) ---")
	fmt.Printf("Max error (h=%.2f): %e\n", h1, err4P_1)
	fmt.Printf("Max error (h=%.2f): %e\n", h2, err4P_2)
	fmt.Printf("Numerical convergence order (p): %.4f\n\n", p4P)

	// Generate data for plotting with a smaller step
	plotH := 0.1
	xPlot, yEulerPlot := eulerMethod(y0, x0, xn, plotH)
	_, yTrapPlot := trapezoidalMethod(y0, x0, xn, plotH)
	_, y4P_Plot := fourPointMethod(y0, x0, xn, plotH)
	yAnalyticalPlot := make([]float64, len(xPlot))
	for i, x := range xPlot {
		yAnalyticalPlot[i] = analyticalSolution(x)
	}

	err := generatePlot(xPlot, yAnalyticalPlot, yEulerPlot, yTrapPlot, y4P_Plot)
	if err != nil {
		fmt.Println("Error generating plot:", err)
	} else {
		fmt.Println("Plot successfully generated in plot.html")
	}
}
