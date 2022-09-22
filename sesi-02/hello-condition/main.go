package main

import "fmt"

func main() {
	var currentyear = 2021

	if age := currentyear - 1998; age < 17 {
		fmt.Printf("kamu belum boleh membuat kartu sim => %d", age)
	} else {
		fmt.Println("kamu sudah boleh membuat kartu sim")
	}
}
