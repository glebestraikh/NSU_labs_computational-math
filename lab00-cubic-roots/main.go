package main

import (
	"fmt"
	"math"
	"strings"
)

type CubicEquation struct {
	a, b, c, d float64
}

// f(x) = ax³ + bx² + cx + d
func (eq CubicEquation) f(x float64) float64 {
	return eq.a*x*x*x + eq.b*x*x + eq.c*x + eq.d
}

// f'(x) = 3ax² + 2bx + c
func (eq CubicEquation) derivative(x float64) float64 {
	return 3*eq.a*x*x + 2*eq.b*x + eq.c
}

// Отделение корней методом 1: табулирование
func (eq CubicEquation) separateRoots(step, a, b float64) [][2]float64 {
	var intervals [][2]float64
	x1 := a
	y1 := eq.f(x1)
	skipNext := false

	for x2 := a + step; x2 <= b; x2 += step {
		y2 := eq.f(x2)

		// Если нашли точный корень
		if !math.IsNaN(y2) && math.Abs(y2) < 1e-10 {
			intervals = append(intervals, [2]float64{x2 - step/2, x2 + step/2})
			skipNext = true
			x1 = x2
			y1 = y2
			continue
		}

		// Смена знака (но не сразу после точного корня)
		if !skipNext && !math.IsNaN(y1) && !math.IsNaN(y2) && y1*y2 < 0 {
			intervals = append(intervals, [2]float64{x1, x2})
		}

		skipNext = false
		x1 = x2
		y1 = y2
	}
	return intervals
}

// Отделение корней методом 2: теорема об оценке корней многочлена
func (eq CubicEquation) getRootBounds() (float64, float64) {
	coeffs := []float64{math.Abs(eq.a), math.Abs(eq.b), math.Abs(eq.c), math.Abs(eq.d)}

	// Находим A = max(|a_k|) для k от 1 до n
	A := 0.0
	for i := 1; i < len(coeffs); i++ {
		if coeffs[i] > A {
			A = coeffs[i]
		}
	}

	// Находим B = max(|a_k|) для k от 0 до n-1
	B := 0.0
	for i := 0; i < len(coeffs)-1; i++ {
		if coeffs[i] > B {
			B = coeffs[i]
		}
	}

	// R = 1 + A/|a_0|
	R := 1.0 + A/math.Abs(eq.a)

	// r = 1/(1 + B/|a_n|)
	r := 1.0 / (1.0 + B/math.Abs(eq.d))

	return r, R
}

// Метод бисекции для уточнения корня на интервале [a, b]
func (eq CubicEquation) bisection(a, b, eps float64) (float64, int) {
	iter := 0
	for math.Abs(b-a) > eps {
		c := (a + b) / 2
		fc := eq.f(c)
		if math.Abs(fc) < eps {
			return c, iter
		}
		if eq.f(a)*fc < 0 {
			b = c
		} else {
			a = c
		}
		iter++
	}
	return (a + b) / 2, iter
}

// Аналитическое отделение корней для кубического уравнения
func (eq CubicEquation) analyticSeparation() {
	fmt.Println("\n=== Аналитическое отделение корней ===")

	// Для кубического уравнения ax³ + bx² + cx + d = 0
	// Производная: f'(x) = 3ax² + 2bx + c
	// Находим экстремумы (критические точки)
	discriminant := 4*eq.b*eq.b - 12*eq.a*eq.c

	fmt.Printf("Уравнение: %.2fx³ + %.2fx² + %.2fx + %.2f = 0\n", eq.a, eq.b, eq.c, eq.d)
	fmt.Printf("Производная: f'(x) = %.2fx² + %.2fx + %.2f\n", 3*eq.a, 2*eq.b, eq.c)
	fmt.Printf("Дискриминант производной: D = %.4f\n", discriminant)

	if discriminant < 0 {
		fmt.Println("Уравнение имеет ровно 1 действительный корень")
	} else if discriminant == 0 {
		x0 := -eq.b / (3 * eq.a)
		fmt.Printf("Производная имеет один корень (точка перегиба): x₀ = %.4f\n", x0)
		fmt.Println("Уравнение имеет 1 или 2 действительных корня")
	} else {
		// Два экстремума
		sqrtD := math.Sqrt(discriminant)
		x1 := (-2*eq.b - sqrtD) / (6 * eq.a)
		x2 := (-2*eq.b + sqrtD) / (6 * eq.a)

		if x1 > x2 {
			x1 = x2
			x2 = x1
		}

		f1 := eq.f(x1)
		f2 := eq.f(x2)

		fmt.Printf("Экстремумы: x₁ = %.4f, x₂ = %.4f\n", x1, x2)
		fmt.Printf("Значения функции: f(x₁) = %.4f, f(x₂) = %.4f\n", f1, f2)

		if f1*f2 > 0 {
			fmt.Println("Экстремумы имеют одинаковый знак -> 1 действительный корень")
		} else {
			fmt.Println("Экстремумы имеют разные знаки -> 2 или 3 действительных корня")
		}
	}
}

func main() {
	// Тестовые примеры
	testCases := []struct {
		name string
		eq   CubicEquation
	}{
		{
			name: "Пример 1: x³ - 6x² + 11x - 6 = 0 (корни: 1, 2, 3)",
			eq:   CubicEquation{a: 1, b: -6, c: 11, d: -6},
		},
		{
			name: "Пример 2: x³ - 3x + 2 = 0 (корни: -2, 1)",
			eq:   CubicEquation{a: 1, b: 0, c: -3, d: 2},
		},
		{
			name: "Пример 3: x³ + x² - x - 1 = 0 (корни: -1, 1)",
			eq:   CubicEquation{a: 1, b: 1, c: -1, d: -1},
		},
		{
			name: "Пример 4: 2x³ - 4x² + 2x = 0 (корни: 0, 1)",
			eq:   CubicEquation{a: 2, b: -4, c: 2, d: 0},
		},
		{
			name: "---- Пример 5: x³ + 3x² + 3x + 1 = 0 (корень: -1 кратности 3)",
			eq:   CubicEquation{a: 1, b: 3, c: 3, d: 1},
		},
		{
			name: "Пример 6: x³ + 2x² = 0 (корни: 0, -2)",
			eq:   CubicEquation{a: 1, b: 2, c: 0, d: 0},
		},
	}

	epsilon := 1e-12
	delta := 1e-1

	for _, tc := range testCases {
		fmt.Println("\n" + strings.Repeat("=", 70))
		fmt.Println(tc.name)
		fmt.Println(strings.Repeat("=", 70))

		// Аналитическое отделение
		tc.eq.analyticSeparation()

		// Метод 1: Табулирование
		fmt.Println("\n=== Метод 1: Отделение корней табулированием ===")
		r, R := tc.eq.getRootBounds()
		fmt.Printf("Оценка кольца корней методом 2: r = %.4f, R = %.4f\n", r, R)
		fmt.Printf("Используем интервал: [%.1f, %.1f] с шагом %.2f\n", -R, R, delta)

		intervals1 := tc.eq.separateRoots(delta, -R, R)
		fmt.Printf("Найдено интервалов со сменой знака: %d\n", len(intervals1))

		for i, interval := range intervals1 {
			fmt.Printf("\nИнтервал %d: [%.4f, %.4f]\n", i+1, interval[0], interval[1])
			root, iters := tc.eq.bisection(interval[0], interval[1], epsilon)
			if !math.IsNaN(root) {
				fmt.Printf("  Корень: x = %.8f (итераций: %d)\n", root, iters)
				fmt.Printf("  Проверка: f(%.8f) = %.2e\n", root, tc.eq.f(root))
			}
		}

		// Метод 2: Теорема об оценке корней
		fmt.Println("\n=== Метод 2: Теорема об оценке корней многочлена ===")
		fmt.Printf("Внутренний радиус: r = %.4f\n", r)
		fmt.Printf("Внешний радиус: R = %.4f\n", R)
		fmt.Printf("Все корни лежат в кольце: %.4f < |x| < %.4f\n", r, R)
	}
}
