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

// Ein Schüler (Student) hat einen Vor- und Nachnamen.
//Eine Klasse (Class) besteht aus einer Reihe von Schülern.
//Ein Modul hat eine eindeutige Nummer (z.B. 346) und wird von einer Reihe von Klassen besucht.
//Erstelle die notwendigen Datenstrukturen mit entsprechenden Beispieldaten (d.h. mindestens zwei Klassen mit je drei
//Schülern und mindestens drei Module, die von einer oder mehreren Klassen besucht werden). Gib die Daten anschliessend
//per fmt.Println() auf die Konsole aus.
