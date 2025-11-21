package main

import "fmt"

func main() {

	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}

	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Student
	}
	var SINA24aL Class = Class{
		Students: []Student{
			{FirstName: "Janis", LastName: "Huber"},
			{FirstName: "levin", LastName: "Schmid"},
			{FirstName: "Andri", LastName: "Gruenig"},
		},
	}
	var SINA24bL Class = Class{
		Students: []Student{
			{FirstName: "Weiss", LastName: "deren"},
			{FirstName: "doch", LastName: "nicht"},
		},
	}

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int][]Class{
		1: {SINA24aL, SINA24bL},
		2: {SINA24aL},
	}

	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
