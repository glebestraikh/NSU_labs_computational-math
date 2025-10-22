package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type rootInfo struct {
	method     string
	value      float64
	iterations int
	fValue     float64
}

type result struct {
	id    int
	name  string
	f     func(float64) float64
	a, b  float64
	roots []rootInfo
}

// корень из wolframalpha:
// x ≈ ± 0.792357147111781...
func f1(x float64) float64 {
	return math.Atan(x) - 1/(3*math.Pow(x, 3))
}

// корень из wolframalpha:
// x ≈ -3.6828
// x ≈ 0.016625
// x ≈ 8.1662
func f2(x float64) float64 {
	return 2*math.Pow(x, 3) - 9*math.Pow(x, 2) - 60*x + 1
}

// корень из wolframalpha:
// x ≈ -2.6983
// x ≈ -0.607815
func f3(x float64) float64 {
	if -x <= 0 {
		fmt.Println("f3: логарифм от неположительного числа")
		fmt.Println(-x)
		os.Exit(1)
	}
	return math.Log2(-x)*(x+2) + 1
}

// корень из wolframalpha:
// x ≈ 1.35204
func f4(x float64) float64 {
	return math.Sin(x+math.Pi/3) - 0.5*x
}

// Производные для метода Ньютона
func df1(x float64) float64 {
	return 1/(1+x*x) + 1/(math.Pow(x, 4))
}

func df2(x float64) float64 {
	return 6*math.Pow(x, 2) - 18*x - 60
}

func df3(x float64) float64 {
	if -x <= 0 {
		fmt.Println("df3: логарифм от неположительного числа")
		fmt.Println(-x)
		os.Exit(1)
	}
	return (x + x*math.Log(-x) + 2) / (x * math.Log(2))
}

func df4(x float64) float64 {
	return math.Sin(math.Pi/6-x) - 0.5
}

// Отделение корней
func separateRoots(f func(float64) float64, a, b, step float64) [][2]float64 {
	var intervals [][2]float64
	x1 := a
	y1 := f(x1)

	for x2 := a + step; x2 <= b; x2 += step {
		y2 := f(x2)
		if !math.IsNaN(y1) && !math.IsNaN(y2) && y1*y2 <= 0 {
			intervals = append(intervals, [2]float64{x1, x2})
		}
		x1 = x2
		y1 = y2
	}
	return intervals
}

// Метод бисекции
func bisection(f func(float64) float64, a, b, eps float64) (float64, int) {
	iter := 0
	for math.Abs(b-a) > eps {
		c := (a + b) / 2
		fc := f(c)
		if math.Abs(fc) < eps {
			return c, iter
		}
		if f(a)*fc < 0 {
			b = c
		} else {
			a = c
		}
		iter++
	}
	return (a + b) / 2, iter
}

// Метод Ньютона
func newton(f, df func(float64) float64, x0, eps float64, maxIter int) (float64, int, bool) {
	x := x0
	for i := 0; i < maxIter; i++ {
		fx := f(x)
		if math.Abs(fx) < eps {
			return x, i, true
		}
		dfx := df(x)
		x = x - fx/dfx
	}
	return x, maxIter, false
}

// Метод секущих
func secant(f func(float64) float64, x0, x1, eps float64, maxIter int) (float64, int, bool) {
	for i := 0; i < maxIter; i++ {
		fx0 := f(x0)
		fx1 := f(x1)
		if math.Abs(fx1) < eps {
			return x1, i, true
		}
		x2 := x1 - fx1*(x1-x0)/(fx1-fx0)
		x0 = x1
		x1 = x2
	}
	return x1, maxIter, false
}

func main() {
	eps := 1e-8
	maxIter := 100

	equations := []struct {
		name string
		f    func(float64) float64
		df   func(float64) float64
		a, b float64
	}{
		{"f₁(x) = arctan(x) - 1/(3x³)", f1, df1, -1.0, 1.0},
		{"f₂(x) = 2x³ - 9x² - 60x + 1", f2, df2, -5, 10},
		{"f₃(x) = log₂(-x)(x+2) + 1", f3, df3, -3.0, -0.1},
		{"f₄(x) = sin(x+π/3) - 0.5x", f4, df4, 0.0, 2},
	}

	var results []result

	// прохожусь по каждому уравнению и первым делом нахожу интервалы
	for id, eq := range equations {
		fmt.Printf("\n%s\n", strings.Repeat("=", 70))

		fmt.Printf("%s на [%.2f, %.2f]\n", eq.name, eq.a, eq.b)

		fmt.Println(strings.Repeat("=", 70))

		// отделение корней
		var intervals [][2]float64
		if eq.a < 0 && eq.b > 0 {
			intervals = append(intervals, separateRoots(eq.f, eq.a, -1e-16, 0.1)...)
			intervals = append(intervals, separateRoots(eq.f, 1e-16, eq.b, 0.1)...)
		} else {
			intervals = separateRoots(eq.f, eq.a, eq.b, 0.1)
		}

		fmt.Printf("\nНайдено интервалов с корнями: %d\n", len(intervals))

		// структура для хранения результатов по уравнению, чтобы построить график потом
		result := result{
			id:   id,
			name: eq.name,
			f:    eq.f,
			a:    eq.a,
			b:    eq.b,
		}

		// прохожусь по каждому интервалу и применяю все три метода
		for i, interval := range intervals {
			fmt.Printf("\nИнтервал #%d: [%.4f, %.4f]\n", i+1, interval[0], interval[1])

			// Бисекция
			root, iter := bisection(eq.f, interval[0], interval[1], eps)
			fmt.Printf("  Бисекция:     x = %.8f, итераций: %d, f(x) = %.2e\n", root, iter, eq.f(root))
			result.roots = append(result.roots, rootInfo{"Бисекция", root, iter, eq.f(root)})

			// Ньютон
			x0 := (interval[0] + interval[1]) / 2
			root, iter, success := newton(eq.f, eq.df, x0, eps, maxIter)
			if success {
				fmt.Printf("  Ньютон:       x = %.8f, итераций: %d, f(x) = %.2e\n", root, iter, eq.f(root))
				result.roots = append(result.roots, rootInfo{"Ньютон", root, iter, eq.f(root)})
			} else {
				fmt.Printf("  Ньютон:       не сошёлся\n")
			}

			// Секущие
			root, iter, success = secant(eq.f, interval[0], interval[1], eps, maxIter)
			if success {
				fmt.Printf("  Секущие:      x = %.8f, итераций: %d, f(x) = %.2e\n", root, iter, eq.f(root))
				result.roots = append(result.roots, rootInfo{"Секущие", root, iter, eq.f(root)})
			} else {
				fmt.Printf("  Секущие:      не сошёлся\n")
			}
		}

		results = append(results, result)
	}

	generateHTML(results)
}
