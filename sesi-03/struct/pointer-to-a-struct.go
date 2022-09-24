package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee = Employee{name: "Airell", age: 23, division: "Curriculum Developer"}

	var employee1 *Employee = &employee

	fmt.Printf("Employee name: %+v\n", employee.name)
	fmt.Printf("Employee1 name: %+v\n", employee1.name)

	fmt.Println(strings.Repeat("#", 20))

	employee1.name = "Ananda"

	fmt.Printf("Employee name: %+v\n", employee.name)
	fmt.Printf("Employee1 name: %+v\n", employee1.name)
}
