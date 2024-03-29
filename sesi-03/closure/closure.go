package main

import (
	"fmt"
	"strings"
)

func main() {
	//Closure(Declare closure in variable)
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}

	var numbers = []int{4, 93, 77, 10, 52, 22, 34}
	fmt.Println(evenNumbers(numbers...))

	// Closure (IIFE)
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}
		return temp == str
	}("kakak")
	fmt.Println(isPalindrome)

	// Closure (Closure as a return value)
	var studentLists = []string{"Airell", "Nanda", "Schannel", "Marco"}
	var find = findStudent(studentLists)
	fmt.Println(find("airell"))

	// Closure (Callback)
	var numbers1 = []int{2, 5, 8, 10, 3, 99, 23}

	var find1 = findOddNumbers(numbers1, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find1)
}

func findStudent(students []string) func(string) string {

	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s doesn't exist!!!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position+1)
	}
}

func findOddNumbers(numbers1 []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers1 {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
