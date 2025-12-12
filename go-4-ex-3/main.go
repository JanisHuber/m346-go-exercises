package main

import "math"

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) (float64, float64) {
	// calculate the discriminant
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		panic("no real roots")
	}
	sqrtDiscriminant := math.Sqrt(discriminant)
	root1 := (-b + sqrtDiscriminant) / (2 * a)
	root2 := (-b - sqrtDiscriminant) / (2 * a)
	return root1, root2
}

func main() {
	// TODO: call the function computeQuadraticFormula
	var root1, root2 = computeQuadraticFormula(1, -3, 2)
	println(root1, root2)
}
