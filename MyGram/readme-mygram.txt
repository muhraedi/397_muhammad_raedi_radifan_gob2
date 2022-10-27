Nama lengkap        : Muhammad Raedi Radifan
kode peserta        : 149368582100-397
Final Project       : MyGram

panduan aplikasi :
- membuat database postgree dengan nama "mygram"
- masuk ke terminal
- masuk ke directory "MyGram"
- jalankan perintah go run main.go
- allow firewall yang muncul
- buka postman kemudian masukkan url sesuai router

Tahapan Pembuatan Final Project
1. membuat folder MyGram
2. membuat file main.go
3. go mod init MyGram
4. Install/go get package : 	
    - github.com/gin-gonic/gin
    - gorm.io/driver/postgres
    - gorm.io/gorm
    - github.com/asaskevich/govalidator
    - github.com/dgrijalva/jwt-go
    - golang.org/x/crypto/bcrypt
5. buat beberapa folder yaitu controller, router, database, middleware, helper, models
6. melakukan config untuk connect ke database
7. membuat models untuk auto migration di config
8. membuat controller
    - user
    - photo
    - comment
    - social media
9. membuat helpers
    - jwt
    - bcrypt
10. membuat middlewares
    - authentication
    - authorization
11. membuat router controller

token 1: 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJwZXJ0YW1hQGdtYWlsLmNvbSIsImlkIjoxfQ.hTUzdt75noQDM9PAC8wQgv7WMfmNhTM4MhWyOmcpRzY
token 2:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJrZWR1YUBnbWFpbC5jb20iLCJpZCI6Mn0.pjOlcizYPZ12Xm84JDWDIRwbsMCMrsZH0r_ZsFZGmio
token 3:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJrZXRpZ2FAZ21haWwuY29tIiwiaWQiOjN9.7kiXm9uu94lCKkO51ZbzRAOO6VLGjo4fPKaRki0USig
token 4:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJrZWVtcGF0QGdtYWlsLmNvbSIsImlkIjo0fQ.RnNUqDpaCB98xsjU6_4ovcdeJJwn-zNwRBwNW9Yzyj4