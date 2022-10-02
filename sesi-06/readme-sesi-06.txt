============================
  GO Micro, Web Services & 
 Web Server Go Programming
============================

-> Web Server
    - net/http >> aplikasi berbasis web seperti contohnya routing,
                  templating, web server dan lain-lain
    - HandleFunc >> digunakan untuk keperluan routing yang dimana
                    function tersebut menerima 2 parameter
        - Parameter pertama digunakan untuk mendefinisikan 
          path routingnya
        - parameter kedua menerima sebuah function dengan 2 parameter
          yaitu http.ResponseWriter dan pointer http.Request
    - http.ResponseWriter >>
      sebuah interface dengan berbagai method yang digunakan untuk
      mengirim response balik kepada client yang mengirimkan request
    - http.Request >>
      sebuah struct yang digunakan untuk mendapat berbagai macam data
      request seperti form value, headers, url parameter dan lain-lain
    - ListenAndServe >> untuk menjalankan server aplikasi
    - API >> GET, POST
    - method Header dari interface http.ResponseWriter yang kemudian 
      dichaining dengan method Set dilakukan untuk menentukan bentuk 
      dari data response yang ingin kita kirimkan kepada client
    - method NewEncoder yang berasal dari package json, yang kemudian
      kita chaining dengan method Encode untuk mengkonversi datanya 
      menjadi bentuk JSON
    - http.StatusBadRequest >>
      sebuah konstanta dari package http.StatusBadRequest yang 
      merepresentasikan status 400
    - method FormValue yang berasal dari struct http.Request digunakan
      untuk mendapatkan nilai input form dari client
    - Ketika kita ingin menampilkan data pada file html, maka kita 
      perlu menyelipkan data tersebut kedalam dua tanda kurung {{}}, 
      lalu menempatkan tanda titik didalamnya. Tanda titik . yang 
      berada di dalam 2 tanda kurung tersebut merepresentasikan data
      yang ingin kita tampilkan.
    - function template.ParseFiles yang berasal dari package html/template
      digunakan untuk mem-parsing file html 
    - Format penulisan range loop adalah {{range $namaVariable := .}}, yang
      dimana cara penulisan nama variablenya diawali dengan tanda dolar $
    - {{end}} >> mengakhiri looping
-> Gin Framework
    - Gin >> sebuah framework untuk bahasa Go yang digunakan untuk keperluan 
             http routing
    - go get -u github.com/gin-gonic/gin >> install framework
    - Setting Up Environment >> 
      membuat 2 folfer dengan nama routers dan controllers
      - routers akan menjadi tempat kita dalam menaruh konfigurasi
        dari routingnya
      - controllers akan menjadi tempat kita untuk menaruh endpoint-endpoint 
        yang kita perlukan
    - *gin.Context mempunyai berbagai macam method yang dapat kita gunakan 
      untuk mendapat request body dari client, mengirim response dan lain-lain
    - ShouldBindJSON >>
      sebuah method dari tipe data *gin.Context yang digunakan untuk 
      mem-binding data JSON yang kirimkan oleh client sebagai request
      body kepada server
    - Method ShouldBindJSON menerima sebuah parameter pointer dari variable
      yang akan menampung hasil dari data binding tersebut
    - Method ShouldBindJSON akan mereturn sebuah error jika memang terjadi
      sebuah error, maka dari itu kita perlu memvalidasi terlebih dahulu 
      jika terjadi sebuah error
    - method AbortWithError digunakan untuk melempar error jika memang 
      ada error yang terjadi
    - Method AbortWithError menerima 2 parameter. Parameter pertama adalah 
      status errornya, kemudian parameter kedua adalah data errornya
    - Method JSON digunakan untuk mengirim response kepada client dengan 
      format data JSON. Method ini menerima 2 parameter. Parameter pertama
      adalah status response nya, dan parameter kedua adalah data response 
      yang di kirimkan kepada client
    - StartServer() >> menjalankan server dari aplikasi
    - *gin.Engine yang berasal dari Gin, dan kita gunakan untuk menjalankan 
      server, sebagai multiplexer dari routing dan lain-lain
    - method POST >> diberikan 2 parameter
      - Parameter pertama adalah route path nya
      - parameter kedua memerlukan handler atau endpoint nya. Endpoint 
        harus merupakan sebuah function yang menerima satu parameter 
        dengan tipe data *gin.Context
    - method ctx.Param merupakan sebuah method yang digunakan untuk 
      mendapat request parameter yang dikirimkan oleh client
    - Method ctx.Param menerima satu parameter dimana kita perlu meletekkan
      nama parameter yang kita buat pada router nya nanti
      