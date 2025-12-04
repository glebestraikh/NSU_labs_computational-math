package direct

// a - нижняя диагональ
// b - главная диагональ
// c - верхняя диагональ
// d - вектор правой части
func TridiagonalSolve(a, b, c, d []float64) []float64 {
	n := len(d)
	alpha := make([]float64, n)
	beta := make([]float64, n)
	x := make([]float64, n)

	// Прямой ход
	// 	Для первого уравнения трёхдиагональной системы: b_0 x_0 + c_0 x_1 = d_0
	alpha[0] = -c[0] / b[0] // первый коэффициент прогонки
	beta[0] = d[0] / b[0]   //  первый свободный член

	// a_i x_{i-1} + b_i x_i + c_i x_{i+1} = d_i
	// a_i (\alpha_{i-1} x_i + \beta_{i-1}) + b_i x_i + c_i x_{i+1} = d_i
	// x_i = -\frac{c_i}{b_i + a_i \alpha_{i-1}} x_{i+1} + \frac{d_i - a_i \beta_{i-1}}{b_i + a_i \alpha_{i-1}}
	for i := 1; i < n; i++ {
		denom := b[i] + a[i]*alpha[i-1] // x_i = \alpha_i x_{i+1} + \beta_i
		if i < n-1 {
			alpha[i] = -c[i] / denom
		}
		beta[i] = (d[i] - a[i]*beta[i-1]) / denom
	}

	// Обратный ход
	x[n-1] = beta[n-1] // x_{n-1} = \beta_{n-1}
	for i := n - 2; i >= 0; i-- {
		x[i] = alpha[i]*x[i+1] + beta[i] // x_i = \alpha_i \cdot x_{i+1} + \beta_i
	}

	return x
}
