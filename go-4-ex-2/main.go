package main

import "math"

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt

func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func main() {
	// TODO: call the function computeHypotenuse
	var hypotenuse float64 = computeHypotenuse(3, 4)
	println(hypotenuse) // expected output: 5
}
