package main

import "fmt"

func main() {

	var student1, student2, student3 string = "satu", "dua", "tiga"
	var first, second, third int
	first, second, third = 1, 2, 3

	// short declaration
	var name, age, address = "Airell", 23, "Jalan Sudirman"
	pertama, kedua, ketiga := "1", 2, "3"

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)

	fmt.Println(name, age, address)
	fmt.Println(pertama, kedua, ketiga)
}
