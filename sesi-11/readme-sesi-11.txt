==================
    Unit Test
==================
Bahasa Go telah menyediakan suatu package bernama testing, 
yang dapat kita gunakan untuk membuat pada aplikasi bahasa Go kita.
-> Terdapat aturan dari penamaan file test yang harus kita ikuti dalam 
   membuat file testing pada bahasa Go. Berikut merupakan aturannya:
    - File testing harus diakhiri dengan penamaan _test.go
    - Misalkan kita mempunyai sebuah file bernama calculation.go, dan 
      ketika kita ingin membuat file testing untuk file calculation.go,
      maka perlu membuat file dengan nama calculation_test.go. Perlu diingat
      disini bahwa penaman file testing boleh bebas, asalkan di akhiri dengan
      penaman _test.go
-> Terdapat juga aturan untuk membuat function testing pada bahasa Go. 
   Berikut merupakan aturannya:
    - Function untuk testing harus diwali dengan penamaan Test.
    - Misalkan kita mempunyai suatu function bernama sum dan kita ingin membuat 
      unit test dari function tersebut, maka kita dapat memberikan nama untuk 
      function unit test kita seperti TestSum.
    - Perlu diingat disini bahwa tidak ada aturan untuk nama belakang dari suatu 
      function testing, yang terpenting adalah harus diawali dengan penaman Test.
    - Kemudian function testing harus menerima satu parameter dengan tipe data 
      *testing.T dan tidak mengembalikan suatu return value apapun. Tipe data 
      *testing.T merupakan sebuah struct dari package testing untuk membuat suatu unit test.