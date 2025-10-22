package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func generateHTML(results []result) {
	html := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Решение нелинейных уравнений</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.9.1/chart.min.js"></script>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; background: #f5f5f5; }
        .container { max-width: 1200px; margin: 0 auto; }
        h1 { color: #333; text-align: center; }
        .equation { background: white; padding: 20px; margin: 20px 0; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .chart-container { height: 400px; margin: 20px 0; }
        table { width: 100%; border-collapse: collapse; margin: 10px 0; }
        th, td { padding: 10px; text-align: left; border-bottom: 1px solid #ddd; }
        th { background: #4CAF50; color: white; }
        .root { background: #e8f5e9; padding: 5px; border-radius: 3px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Решение нелинейных уравнений</h1>
`
	for _, r := range results {
		html += fmt.Sprintf(`
        <div class="equation">
            <h2>%s</h2>
            <p><strong>Интервал:</strong> [%.2f, %.2f]</p>
            <div class="chart-container">
                <canvas id="chart%d"></canvas>
            </div>
            <h3>Найденные корни:</h3>
            <table>
                <tr>
                    <th>Метод</th>
                    <th>Корень</th>
                    <th>Итерации</th>
                    <th>f(x)</th>
                </tr>
`, r.name, r.a, r.b, r.id)

		for _, root := range r.roots {
			html += fmt.Sprintf(`
                <tr>
                    <td>%s</td>
                    <td class="root">%.8f</td>
                    <td>%d</td>
                    <td>%.2e</td>
                </tr>
`, root.method, root.value, root.iterations, root.fValue)
		}

		html += `
            </table>
        </div>
`
	}

	html += `
    </div>
    <script>
`
	for _, r := range results {
		html += generateChartJS(r)
	}

	html += `
    </script>
</body>
</html>
`
	os.WriteFile("results.html", []byte(html), 0644)
	fmt.Println("\nРезультаты сохранены в файл results.html")
}

func generateChartJS(r result) string {
	// Генерация точек для графика
	points := 500
	step := (r.b - r.a) / float64(points)
	var dataPoints []string

	for i := 0; i <= points; i++ {
		x := r.a + float64(i)*step
		y := r.f(x)
		if math.IsNaN(y) || math.IsInf(y, 0) || math.Abs(y) > 100 {
			// Разрыв линии
			dataPoints = append(dataPoints, fmt.Sprintf("{x: %.6f, y: null}", x))
		} else {
			dataPoints = append(dataPoints, fmt.Sprintf("{x: %.6f, y: %.6f}", x, y))
		}
	}

	// Точки корней
	var rootPoints []string
	for _, root := range r.roots {
		if root.method == "Бисекция" {
			rootPoints = append(rootPoints, fmt.Sprintf("{x: %.6f, y: %.6f}", root.value, root.fValue))
		}
	}

	// Точки для линии y=0
	zeroPoints := fmt.Sprintf("[{x: %.6f, y: 0}, {x: %.6f, y: 0}]", r.a, r.b)

	return fmt.Sprintf(`
        new Chart(document.getElementById('chart%d'), {
            type: 'scatter',
            data: {
                datasets: [{
                    label: 'f(x)',
                    data: [%s],
                    borderColor: 'rgb(75, 192, 192)',
                    backgroundColor: 'rgba(75, 192, 192, 0.1)',
                    borderWidth: 2,
                    pointRadius: 0,
                    showLine: true,
                    tension: 0.1
                }, {
                    label: 'y = 0',
                    data: %s,
                    borderColor: 'rgba(255, 99, 132, 0.5)',
                    borderWidth: 2,
                    borderDash: [5, 5],
                    pointRadius: 0,
                    showLine: true
                }, {
                    label: 'Корни',
                    data: [%s],
                    backgroundColor: 'rgb(255, 99, 132)',
                    borderColor: 'rgb(255, 99, 132)',
                    pointRadius: 8,
                    pointHoverRadius: 10,
                    showLine: false
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    title: { display: true, text: '%s', font: { size: 16 } },
                    legend: { display: true, position: 'top' }
                },
                scales: {
                    x: { 
                        type: 'linear',
                        title: { display: true, text: 'x', font: { size: 14 } },
                        grid: { color: 'rgba(0, 0, 0, 0.1)' }
                    },
                    y: { 
                        title: { display: true, text: 'f(x)', font: { size: 14 } },
                        grid: { color: 'rgba(0, 0, 0, 0.1)' }
                    }
                }
            }
        });
`, r.id, strings.Join(dataPoints, ","), zeroPoints, strings.Join(rootPoints, ","), r.name)
}
