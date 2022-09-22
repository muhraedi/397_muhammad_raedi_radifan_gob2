package main

import "fmt"

func main() {
	// belajar hello
	// fmt.Println("Hello")
	fmt.Println("Arif", "Setiawan", "Belajar Ngoding")

	/*
		belajar print
		tahap 2
	*/
	fmt.Println("Hello")
	fmt.Println("Hello", "Arif")
	fmt.Println("====================")
	fmt.Print("Hello Arif\n")
	fmt.Print("Hello", "Arif\n")
	fmt.Print("Hello", " ", "Arif\n")
	fmt.Println("====================")

	// belajar variabel
	var namaLengkap string
	var umur int = 99

	namaLengkap = "Arif"
	umur = 28

	fmt.Println("Hello", namaLengkap)
	fmt.Println("Umur", umur)
	fmt.Println("====================")

	// belajar variabel without type data
	var namaAnda = "Anonymous"
	var umurAnda = 28

	fmt.Printf("%T, %T\n", namaAnda, umurAnda)
	fmt.Println("====================")

	// belajar short declaration
	namaSaya := "Arif"
	umurSaya := 28

	fmt.Printf("%T, %T\n", namaSaya, umurSaya)
	fmt.Println("====================")

	// belajar multivariabel declaration
	var one, two, three string = "1", "2", "3"
	var angkaOne, angkaTwo, angkaThree int = 1, 2, 3
	fmt.Println("test data string multiple =>", one, two, three)
	fmt.Println("test data int multiple =>", angkaOne, angkaTwo, angkaThree)
	fmt.Println("====================")

	// belajar multivariabel declaration short
	var oneData, twoData, threeData = 1, true, "Arif"
	dataOne, dataTwo, dataThree := true, 55.5, "Setiawan"
	fmt.Println("test data string multiple =>", oneData, twoData, threeData)
	fmt.Println("test data int multiple =>", dataOne, dataTwo, dataThree)
	fmt.Println("====================")

	// belajar underscore variabel
	var testVariabel string
	var oneName, twoName, alamatSaya, dataUmur = "Arif", "Setiawa", "Lampung", 28
	_, _, _, _, _ = testVariabel, oneName, twoName, alamatSaya, dataUmur
	// fmt.Printf Usage
	fmt.Printf("test underscore variabel => %T\n", testVariabel)
	fmt.Printf("halo nama saya %s, umur saya adalah %d, saya beralamat di %s\n", oneName, dataUmur, alamatSaya)
	fmt.Println("====================")

	// belajar konstanta
	const full_name string = "Airell Jordan"
	fmt.Printf("Hello %s\n", full_name)
	fmt.Println("====================")

	// belajar operator aritmatika
	//var value = 2 + 2 * 3
	var value = (2 + 2) * 3
	fmt.Println(value)

}
