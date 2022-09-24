package main

import "fmt"

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee = Employee{}

	employee.name = "Airell"
	employee.age = 23
	employee.division = "Curriculum Developer"

	var employee1 = Employee{name: "Ananda", age: 23, division: "Finance"}

	fmt.Printf("Employee: %+v\n", employee)
	fmt.Printf("Employee1: %+v\n", employee1)
}
