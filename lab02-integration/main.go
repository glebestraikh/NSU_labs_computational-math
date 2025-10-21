package main

import (
	"fmt"
	"math"
)

const (
	reset = "\033[0m"
	green = "\033[32m"
)

const (
	scale = 0.01
)

type Segment struct {
	a, b float64
}

type Example struct {
	name             string
	f                func(float64) float64
	value            float64
	secondDerivative func(float64) float64
	fourthDerivative func(float64) float64
	segment          Segment
}

type Answer struct {
	value         float64
	errorAnalytic float64
	errorRunge    float64
}

// Пример функции
func f(x float64) float64 {
	return math.Log10(x+2) / x
}

func secondDerivative(x float64) float64 {
	numerator := 2*math.Pow(x+2, 2)*math.Log(x+2) - x*(3*x+4)
	denominator := math.Pow(x, 3) * math.Pow(x+2, 2) * math.Log(10)
	return numerator / denominator
}

func fourthDerivative(x float64) float64 {
	numerator := 24*math.Pow(x+2, 4)*math.Log(x+2) - 2*x*(25*math.Pow(x, 3)+104*math.Pow(x, 2)+168*x+96)
	denominator := math.Pow(x, 5) * math.Pow(x+2, 4) * math.Log(10)
	return numerator / denominator
}

// Составная квадратурная формула трапеции
func trapezoidalRule(ex Example, intervals int, calculateErrors bool) Answer {
	a, b := ex.segment.a, ex.segment.b
	f := ex.f
	step := (b - a) / float64(intervals)
	sum := 0.0

	for i := 0; i < intervals; i++ {
		x0 := a + float64(i)*step
		x1 := a + float64(i+1)*step
		sum += (f(x0) + f(x1)) * step / 2.0
	}

	ans := Answer{value: sum}

	if calculateErrors {
		// Аналитическая погрешность
		maxSecondDer := math.Abs(ex.secondDerivative(a))
		for xi := a + scale; xi <= b; xi += scale {
			secondDer := math.Abs(ex.secondDerivative(xi))
			if secondDer > maxSecondDer {
				maxSecondDer = secondDer
			}
		}
		ans.errorAnalytic = maxSecondDer * (b - a) * step * step / 24

		// Погрешность Рунге
		ansRunge := trapezoidalRule(ex, intervals*2, false)
		ans.errorRunge = (4 * (math.Abs(ansRunge.value - ans.value))) / 3
	}

	return ans
}

// Составная квадратурная формула Симпсона
func simpsonRule(ex Example, intervals int, calculateErrors bool) Answer {
	a, b := ex.segment.a, ex.segment.b
	f := ex.f
	step := (b - a) / float64(intervals)
	sum := 0.0

	x0 := a
	for i := 0; i < intervals; i++ {
		x1 := x0 + step
		xm := (x0 + x1) / 2
		sum += (f(x0) + 4*f(xm) + f(x1)) * step / 6.0
		x0 = x1
	}

	ans := Answer{value: sum}

	if calculateErrors {
		// Аналитическая погрешность
		maxFourthDer := math.Abs(ex.fourthDerivative(a))
		for xi := a + scale; xi <= b; xi += scale {
			fourthDer := math.Abs(ex.fourthDerivative(xi))
			if fourthDer > maxFourthDer {
				maxFourthDer = fourthDer
			}
		}
		ans.errorAnalytic = maxFourthDer * (b - a) * math.Pow(step, 4) / 2880

		// Погрешность Рунге
		ans2 := simpsonRule(ex, intervals*2, false)
		ans.errorRunge = (16 * (math.Abs(ans2.value - ans.value))) / 15
	}

	return ans
}

func main() {
	examples := []Example{
		{
			name:             "log10(x + 2) / x на [1.2, 2]",
			f:                f,
			value:            0.281613,
			secondDerivative: secondDerivative,
			fourthDerivative: fourthDerivative,
			segment:          Segment{1.2, 2},
		},
	}

	ex := examples[0]

	fmt.Println("═══════════════════════════════════════════════════════════════════════════════")
	fmt.Printf("Численное интегрирование: %s\n", ex.name)
	fmt.Println("═══════════════════════════════════════════════════════════════════════════════")

	intervals := []int{10, 20, 40, 50}
	for i := 0; i < len(intervals); i++ {
		ansTrapezoidal := trapezoidalRule(ex, intervals[i], true)
		ansSimpson := simpsonRule(ex, intervals[i], true)

		fmt.Println("------------------------------------------------------------------------------")
		fmt.Printf("Итерация %d | Интервалов: %d\n", i+1, intervals[i])
		fmt.Println("------------------------------------------------------------------------------")

		fmt.Printf("Трапеция | Вычисленное значение: %.8f\n", ansTrapezoidal.value)
		fmt.Printf("Трапеция | Разность квадратурной формулы и wolframalpha: %.8f\n", math.Abs(ansTrapezoidal.value-ex.value))
		fmt.Printf("Трапеция | Аналитическая погрешность: %.8f\n", ansTrapezoidal.errorAnalytic)
		fmt.Printf("Трапеция | Погрешность по правилу Рунге: %.8f\n\n", ansTrapezoidal.errorRunge)

		fmt.Printf(green+"Симпсон | Вычисленное значение: %.8f"+reset+"\n", ansSimpson.value)
		fmt.Printf(green+"Симпсон | Разность квадратурной формулы и wolframalpha: %.8f"+reset+"\n", math.Abs(ansSimpson.value-ex.value))
		fmt.Printf(green+"Симпсон | Аналитическая погрешность: %.8f"+reset+"\n", ansSimpson.errorAnalytic)
		fmt.Printf(green+"Симпсон | Погрешность по правилу Рунге: %.8f"+reset+"\n", ansSimpson.errorRunge)
	}

	fmt.Println("==============================================================================")
	fmt.Println("Вычисления завершены")
	fmt.Println("==============================================================================")
}
