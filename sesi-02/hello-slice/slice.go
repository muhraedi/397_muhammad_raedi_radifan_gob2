package main

import (
	"fmt"
)

func main() {
	// Silce
	// var fruits = []string{"apple", "banana", "mango"}

	// _ = fruits

	// Slice (make function)
	// var fruits = make([]string, 3)
	// _ = fruits

	// fmt.Printf("%#v", fruits)

	// Slice (append function)
	// var fruits = make([]string, 3)

	// fruits = append(fruits, "apple", "banana", "mango")

	// fmt.Printf("%#v", fruits)

	// Slice (append function with ellipsis)
	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// fruits1 = append(fruits1, fruits2...)

	// fmt.Printf("%#v", fruits1)

	// Slice (copy function)
	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// nn := copy(fruits1, fruits2)

	// fmt.Println("Fruits1 =>", fruits1)
	// fmt.Println("Fruits2 =>", fruits2)
	// fmt.Println("Copied elements =>", nn)

	// Slice (Slicing)
	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	// var fruits2 = fruits1[1:4]
	// fmt.Printf("%#v\n", fruits2)

	// var fruits3 = fruits1[0:]
	// fmt.Printf("%#v\n", fruits3)

	// var fruits4 = fruits1[:3]
	// fmt.Printf("%#v\n", fruits4)

	// var fruits5 = fruits1[:] // sama dengan fruits[:len(fruits1)]
	// fmt.Printf("%#v\n", fruits5)

	// Slice (Combining slicing and append)
	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	// fruits1 = append(fruits1[:3], "rambutan")

	// fmt.Printf("%#v\n", fruits1)

	// Slice (Backing array)
	// var fruits1 = []string{"apple", "mango", "durian", "banana", "starfruit"}

	// var fruits2 = fruits1[2:4]

	// fruits2[0] = "rambutan"

	// fmt.Println("fruits1 => ", fruits1)
	// fmt.Println("fruits2 => ", fruits2)

	// Slice (Cap function)
	/* var fruits1 = []string{"apple", "mango", "durian", "banana"}

	fmt.Println("Fruits1 cap:", cap(fruits1))
	fmt.Println("Fruits1 len:", len(fruits1))

	fmt.Println(strings.Repeat("#", 20))

	var fruits2 = fruits1[0:3]
	fmt.Println("Fruits2 cap:", cap(fruits2))
	fmt.Println("Fruits2 len:", len(fruits2))

	fmt.Println(strings.Repeat("#", 20))

	var fruits3 = fruits1[1:]

	fmt.Println("Fruits3 cap:", cap(fruits3))
	fmt.Println("Fruits3 len:", len(fruits3))

	fmt.Println(strings.Repeat("#", 20)) */

	// Slice (Creating a new backing array)
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars", newCars)
}
