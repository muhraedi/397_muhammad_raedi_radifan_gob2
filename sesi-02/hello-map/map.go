package main

import "fmt"

func main() {
	// Cara 1
	/* var person map[string]string // Deklarasi

	person = map[string]string{} // Inisialisasi

	person["name"] = "Airell"

	person["age"] = "23"

	person["address"] = "Jalan Sudirman" */

	// Cara 2
	/* var person = map[string]string{
		"name":    "Airell",
		"age":     "23",
		"address": "Jalan Sudirman",
	} */

	/* fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"]) */

	// Map (Looping with map)
	/* for key, value := range person {
		fmt.Println(key, ":", value)
	} */

	// Map (Deleting value)
	/* fmt.Println("Before deleting:", person)

	delete(person, "age")

	fmt.Println("After deleting:", person) */

	// Map (Detecting a value)
	/* value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't  exist")
	}

	delete(person, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	} */

	// Map (Combining slice with map)
	var people = []map[string]string{
		{"name": "Airell", "age": "23"},
		{"name": "Nanda", "age": "23"},
		{"name": "Mailo", "age": "15"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
