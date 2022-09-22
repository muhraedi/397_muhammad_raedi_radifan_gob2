package main

import "fmt"

func main() {
	// var currentyear = 2021

	// if age := currentyear - 1998; age < 17 {
	// 	fmt.Printf("kamu belum boleh membuat kartu sim => %d", age)
	// } else {
	// 	fmt.Println("kamu sudah boleh membuat kartu sim")
	// }

	var score = 9

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score >= 3):
		fmt.Println("not bad")
		fallthrough
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}
}
