package main

import (
	"fmt"
	"os"
	"strings"
)

func generatePlot(x, yAnalytical, yEuler, yTrap, yRK4 []float64) error {
	file, err := os.Create("plot.html")
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert data to string format for JavaScript
	xStr := floatsToString(x)
	yAnalyticalStr := floatsToString(yAnalytical)
	yEulerStr := floatsToString(yEuler)
	yTrapStr := floatsToString(yTrap)
	yRK4Str := floatsToString(yRK4)

	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Finite Difference Methods</title>
    <script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
</head>
<body>
    <h1>Comparison of Finite Difference Methods</h1>
    <div id="plot"></div>
    <script>
        var trace1 = {
            x: %s,
            y: %s,
            mode: 'lines',
            name: 'Analytical Solution'
        };

        var trace2 = {
            x: %s,
            y: %s,
            mode: 'markers',
            name: 'Euler (k=1)'
        };

        var trace3 = {
            x: %s,
            y: %s,
            mode: 'markers',
            name: 'k=2'
        };

        var trace4 = {
            x: %s,
            y: %s,
            mode: 'markers',
            name: 'k=4'
        };

        var data = [trace1, trace2, trace3, trace4];

        var layout = {
            title: 'y\' = -y, y(0) = -1',
            xaxis: {
                title: 'x'
            },
            yaxis: {
                title: 'y(x)'
            }
        };

        Plotly.newPlot('plot', data, layout);
    </script>
</body>
</html>
`, xStr, yAnalyticalStr, xStr, yEulerStr, xStr, yTrapStr, xStr, yRK4Str)

	_, err = file.WriteString(htmlContent)
	return err
}

// Helper function to convert a slice of floats to a JavaScript array string
func floatsToString(data []float64) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range data {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%f", v))
	}
	sb.WriteString("]")
	return sb.String()
}
