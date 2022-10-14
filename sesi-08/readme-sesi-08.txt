======================
  JSON, URL Parsing 
  on Go Programming
======================
-> Decoding & Parsing JSON Data
    - Decoding JSON To Struct
      - membuat data JSON sederhana menggunakan tanda backtick``
      - function json.Unmarshal >> mendecode data JSON kepada struct Employee
      - Argumen pertama dari function json.Unmarshal menerima sebuah nilai 
        dengan tipe data slice of byte. argument pertama itulah kita meletakkan 
        data JSON nya tetapi harus kita ubah terlebih dahulu menjadi tipe data 
        slice of byte. argumen kedua, kita meletakkan pointer dari variable 
        result agar setelah data JSON berhasil didecode, datanya akan disimpan 
        kepada variable result
    - Decoding JSON To Map
      - Kita juga bisa men-decode data JSON kepada sebuah tipe data map
      - tidak perlu membuat tag seperti pada sebuah struct
    - Decoding JSON To Empty Interface
      - ketika kita ingin mengakses field-fieldnya, maka harus dilakukan type 
        assertion dari empty interface menjadi tipe data map[string]interface{}
    - Decoding JSON Array To Slice Of Struct
      - Kita juga bisa men-decode data JSON array kepada sebuah tipe data 
        slice of struct
-> URL Parsing
    - Function url.Parse >> parsing string ke bentuk URL
    - Mengembalikan 2 data, variabel objek bertipe url.URL dan error
      (jika ada)
    - query yang ada pada url akan otomatis diparsing, menjadi bentuk 
      map[string][]string, dengan key adalah nama elemen query, dan value array
      string yang berisikan value elemen query.
-> Swaggo
    - library swagger >> membuat dokumentasi dari RESTAPI
    - Set Up
      - go get -u github.com/swaggo/swag/cmd/swag
      - go get -u github.com/swaggo/swag/http-swagger
      - go get -u github.com/alecthomas/template
      - go get -u github.com/gorilla/mux
    - Global >>
      go install github.com/swaggo/swag/cmd/swag
    - Add path >>
      export PATH=$(go env GOPATH)/bin:$PATH
    - Access swagger >>
      http://localhost:8080/swagger/index.html
    - Generate swagger/docs >> 
      swag init -g main.go