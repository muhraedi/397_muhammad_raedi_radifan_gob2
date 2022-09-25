package main

import "fmt"

type Person struct {
	Name      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var people = []Person{
		{
			Name:      "Airell",
			Alamat:    "Jl. Sudirman",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin menjadi golang developer",
		},
		{
			Name:      "Ananda",
			Alamat:    "Jl. Harimau",
			Pekerjaan: "Front-End Developer",
			Alasan:    "Iseng-iseng",
		},
		{
			Name:      "Mailo",
			Alamat:    "Jl. Landak",
			Pekerjaan: "Back-End Developer",
			Alasan:    "Mencoba hal baru",
		},
	}
	var nomor int
	fmt.Print("Masukkan nomor absen(1 ~ 3): ")
	fmt.Scan(&nomor)
	for i, v := range people {
		if nomor == i+1 {
			fmt.Printf("%+v\n", v)
		}
	}
}
