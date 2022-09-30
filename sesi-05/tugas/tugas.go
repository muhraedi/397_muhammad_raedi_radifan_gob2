package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	var userame string
	//var password string

	fmt.Println("Masukkan username: ")
	fmt.Scan(&userame)
	fmt.Println("Masukkan password: ")
	//fmt.Scan(&password)
	password, _ := gopass.GetPasswdMasked()

	if valid, err := validPassword(userame, password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func validPassword(username string, password []byte) (string, error) {
	pass := string(password[:])
	if username == "ifan" && pass == "indonesia" {
		fmt.Println("Berhasil login!")
	} else {
		return "", errors.New("username atau password salah")
	}
	return "", nil
}
