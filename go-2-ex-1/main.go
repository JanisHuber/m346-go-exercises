package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Profile struct {
	FullName         FullName
	BirthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "Janis",
			LastName:  "Huber",
		},
		BirthDate: BirthDate{
			Day:   15,
			Month: 9,
			Year:  2007,
		},
		NumberOfSiblings: 1,
		ZodiacSign:       'J',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings = 2
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
