package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value)	:", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondPerson (value) 	:", *secondPerson)
	fmt.Println("secondPerson (memori address) :", secondPerson)

	fmt.Print("\n", strings.Repeat("#", 30), "\n\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value)	:", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondPerson (value) 	:", *secondPerson)
	fmt.Println("secondPerson (memori address) :", secondPerson)
}
