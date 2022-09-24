package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	profile("Airell", "pasta", "ayam geprek", "ikan roa", "sate padang")

	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sum(numberLists...)
	fmt.Println("Result:", result)

	studentLists := print("Airell", "Nanda", "Mailo", "Schannel", "Marco")
	fmt.Printf("%v", studentLists)

	var diameter float64 = 20
	var area, circumference = calculate(diameter)
	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)

	var diameter1 float64 = 15
	var area1, circumference1 = calculate1(diameter1)
	fmt.Println("Area:", area1)
	fmt.Println("Circumference:", circumference1)

	// greet("Airell", "Jl. sudirman")
	/* var names = []string{"Airell", "Jordan"}
	var printMsg = greet("Heii", names)

	fmt.Println(printMsg) */
}

// fungsi dua parameter yang memiliki tipe data beda
/* func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
} */

// fungsi dua parameter yang memiliki tipe data sama
/* func greet(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
} */

// Function (Return)
/* func greet(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
} */

// Function (Returning multiple values)
func calculate(d float64) (float64, float64) {
	// Menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 3)

	// Menghitung keliling
	var circumference = math.Pi * d

	return area, circumference
}

// Function (Predefined return value)
func calculate1(diameter1 float64) (area1 float64, circumference1 float64) {
	// Menghitung luas
	area1 = math.Pi * math.Pow(diameter1/2, 2)

	// Menghitung keliling
	circumference1 = math.Pi * diameter1

	return area1, circumference1
}

// Function (Variadic function #1)
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

// Function (Variadic function #2)
func sum(numbers ...int) int {

	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

// Function (Variadic function #3)
func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
