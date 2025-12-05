package main

import (
	"fmt"
	"log"
)

// TODO: implement the function computeGrade
func computeGrade(gotPoints float32, maxPoints float32) float32 {
	if gotPoints > maxPoints {
		panic("max points exceeded")
	}
	return 5/maxPoints*gotPoints + 1
}

func main() {
	// TODO: call the function computeGrade
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	var grade float32 = computeGrade(21, 20)
	// Why is it such a complicated expression instead of just a try and catch block?
	// And which languages named it exeptions "panic" hahaha
	fmt.Printf("%.2f\n", grade)
}
