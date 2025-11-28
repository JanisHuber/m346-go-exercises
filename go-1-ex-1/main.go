package main

import "fmt"

func main() {
	var firstName string = "Janis"
	var lastName string = "Huber"
	var dayOfBirth int = 15
	var monthOfBirth int = 9
	var yearOfBirth int = 2007
	var numberOfSiblings int = 1
	var heightInMeters float32 = 1.70
	var zodiacSign rune = '\U+264D'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
