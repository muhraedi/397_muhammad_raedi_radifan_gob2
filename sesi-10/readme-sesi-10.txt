=================================
    Middleware & JsonWebToken
=================================
-> Middleware >>
    merupakan sebuah fungsi yang akan terkesekusi sesudah maupun sebelum mencapai sebuah endpoint.
    Biasa middleware digunakan untuk logging atau untuk mengamankan sebuah endpoint seperti contohnya
    proses autentikasi dan autorisasi. Untuk membuat middleware pada bahasa Go, kita akan menggunakan 
    package net/http dengan menggunakan multiplexer nya agar kita dapat melakukan kustomisasi
-> JWT >>
    - adalah sebuah token berbentuk string panjang yang digunakan untuk pertukaran informasi antara 
      dua belah pihak atau client dan server
    - JsonWebToken biasa digunakan pada aplikasi web yang berbasis SPA atau Single Page Application
    
go get github.com/asaskevich/govalidator
go get github.com/dgrijalva/jwt-go
go get github.com/gin-gonic/gin
go get golang.org/x/crypto
go get gorm.io/driver/postgres
go get gorm.io/gorm

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFpcmVsbEBnbWFpbC5jb20iLCJpZCI6MX0._4Dv5YHe92K_duuFhn8ZmE-UwDkcLt-phFKRBDgBrYo