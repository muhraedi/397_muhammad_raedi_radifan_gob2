===================
    What is GO?
===================

Go Progamming Introduction
-> Instalasi Go
    - installer >> https://golang.org/dl/
    - cek keberhasilan instalasi >> go version
-> Inisialisasi Project
    menggunakan Go Modules >> go mod init <nama-project>
    sebaiknya disamakan dengan nama folder
-> Command
    - go run
    - go build
-> Membuat Project dengan Go
    - touch .go >> package >> import >> main function
    - komen >> // multiline /* inline */

Variables
-> dituliskan tipe datanya 
    - var namaVariable typedata
-> tidak ditulis tipe datanya atau biasa
    - var namaVariable
-> Fungsi fmt.Printf 
    mempunyai kegunaan yang sama dengan fungsi fmt.Println.
    fungsi fmt.Printf ditentukan dari flag yang kita berikan
-> Short Declaration
    menggunakan :=
-> Multiple variable declarations
    variableSatu, variableDua, variableTiga := "valueSatu", valueDua, "valueTiga"
-> Underscore variable (_)
    Dengan variable underscore maka kita dapat menghindari 
    error yang akan terjadi jika kita mempunyai suatu variable 
    yang menganggur atau tidak dipakai
-> fmt.Printf Usage
    - %T >> mengetahui tipe data dari variable
    - \n >> baris baru
    - %s >> memformat suatu nilai dengan tipe data string
    - %d >> memformat tipe data int

Data Types
-> Number (integer)
    - int: Bilangan cacah (Bilangan positif)
    - uint: Bilangan bulat (bilangan positif maupun negatif)
-> Number (float)
    - %f >> memformat angka desimal atau tipe data float 
            menjadi 6 digit angka desimal
    - %.nf >> huruf n = jumlah digit yang ingin kita hasilnya
-> Bool
    - terdiri dari 2 nilai yaitu false, dan true. 
    - %t >> memformat tipe data bool menjadi tipe data string
-> String
    - nilainya di apit oleh tanda quote atau petik dua ("")
    - deklarasinya bisa dengan tanda grave accent/backticks (``)
-> nil & Zero Value
    - nil bukan merupakan tipe data, melainkan sebuah nilai
    - zero value >> nilai default tipe data

Constants & Operators
-> Constant (const)
    variabel yang nilainya tidak dapat diubah
-> Operators
    - Operator aritmatika (+, -, *, /, %, ++, --)
    - Operator perbandingan/relasional
        (==, !=, >, <, >=, <=)
    - Operator logika
        (&&, ||, !)