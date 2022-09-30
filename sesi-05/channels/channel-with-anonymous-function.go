package main

import "fmt"

func main() {

	c := make(chan string)

	student := []string{"Airell", "Mailo", "Indah"}

	for _, v := range student {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c <- result
		}(v)
	}

	for i := 0; i <= 3; i++ {
		print(c)
	}

	close(c)
}

func print(c chan string) {
	fmt.Println(<-c)
}
