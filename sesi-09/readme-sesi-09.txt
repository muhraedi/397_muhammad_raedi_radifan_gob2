=====================
    HTTP Request
=====================
-> Get Request
    - function http.Get >> 
      menerima satu parameter dengan tipe data string yang merupakan sebuah url
      untuk request tujuannya. Kemudian function ini akan mereturn sebuah nilai
      dengan tipe data pointer dari struct http.Response jika request nya berhasil
      dan akan mereturn sebuah tipe data error yang requestnya gagal
    - function ioutil.ReadAll >>
      mengubah nilai yang kita akses dari field Body menjadi sebuah nilai dengan
      tipe data slice of byte
    - menutup response body nya setelah kita selesai membacanya dengan method Close 
      untuk mencegah kebocoran memori atau memory 
-> Post Request
    - function bytes.NewBuffer >>
      mengubah tipe data dari data yang ingin kita kirim menjadi interface io.Reader
    