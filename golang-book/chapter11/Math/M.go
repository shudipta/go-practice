// Finds the average of a series of numbers
package m

import "fmt"

func Sum(xs []float64) float64{
	total := float64(0)
	for _, x := range xs {
		total += x
	}

	return total
}

func Avg(xs []float64) float64 {
	return Sum(xs) / float64(len(xs))
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func main() {
	panic("salk")
	fmt.Println(Sum([]float64{1, 2}))
}