package main

import "fmt"

func main() {

	first, second := 1, "2"

	fmt.Printf("Tipe data variable first adalah %T \n", first)
	fmt.Printf("Tipe data variable second adalah %T \n", second)

	var name = "Airel"
	var age = 23
	var address = "Jalan sudirman"

	fmt.Printf("Halo nama ku %s, umur aku adalah %d, dan aku tinggal di %s", name, age, address)
}
