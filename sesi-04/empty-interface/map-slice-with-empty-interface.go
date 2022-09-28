package main

import "fmt"

func main() {
	rs := []interface{}{1, "Airell", true, 2, "Ananda", true}

	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm
	fmt.Printf("slice : %+v\n", rs)
	fmt.Printf("map : %+v\n", rm)
}
