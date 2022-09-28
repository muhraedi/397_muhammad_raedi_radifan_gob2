package main

func main() {
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}
}
