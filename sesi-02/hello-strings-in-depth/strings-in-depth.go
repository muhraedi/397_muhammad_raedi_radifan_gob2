package main

import (
	"fmt"
)

func main() {
	// indexing terhadap suatu string
	/* city := "Jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	} */

	// slice of bytes cetak nilai string
	/* var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}

	fmt.Println(string(city)) */

	// menggunakan function len()
	// str1 := "makan"
	// str2 := "mânca"

	// fmt.Printf("str1 byte length => %d\n", len(str1))
	// fmt.Printf("str2 byte length => %d\n", len(str2))

	// menggunakan function RuneCpuntInString
	// fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	// fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))

	// menggunakan range loop
	str := "mânca"

	for i, s := range str {
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}
}
