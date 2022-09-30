package main

import "fmt"

func main() {
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hi everyone")
	fmt.Println("Welcome back to Go learning center")
}
