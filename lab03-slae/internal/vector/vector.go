package vector

import "fmt"

func CopyVector(v []float64) []float64 {
	result := make([]float64, len(v))
	copy(result, v)
	return result
}

func PrintVector(name string, v []float64) {
	fmt.Printf("%s = [", name)
	for i, val := range v {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%.6f", val)
	}
	fmt.Println("]")
}
