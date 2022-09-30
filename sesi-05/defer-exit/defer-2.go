package main

import "fmt"

func main() {
	callDeferFunc()
	fmt.Println("Hi everyone!!!")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
