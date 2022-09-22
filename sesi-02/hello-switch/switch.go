package main

import "fmt"

func main() {
	// var score = 8

	// switch score {
	// case 8:
	// 	fmt.Println("perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// default:
	// 	fmt.Println("not bad")
	// }

	// Switch with relational operators
	var score = 9

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score >= 3):
		fmt.Println("not bad")
		fallthrough // Switch (fallthrough keyword)
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}
}
