============================
 Fundamental GO Programming
============================

-> Conditions (if else, switch)
    - Temporary variable
      suatu variable yang dimana variable tersebut hanya
      bisa diakses dan digunakan pada scope block dari 
      suatu kondisional
    - switch
      sifatnya fokus pada satu variabel kemudian di-cek nilainya
    - Switch with relational operators
      menggunakan switch dengan operator perbandingan pada
      sebuah kondisional dengan keyword if, else if dan else
    - Switch (fallthrough keyword)
      melanjutkan pengecekan kepada case selanjutnya walaupun
      suatu case telah terpenuhi kondisinya
    - Nested Conditions
      Kondisional bersarang merupakan sebuah proses kondisional 
      yang didalamnya terdapat proses kondisional kembali

-> Loopings (for)
    - break >> menghentikan sebuah proses looping
    - continue >> melanjutkan suatu proses looping
    - label (:) >> keyword break dan continue akan berlaku pada 
                   blok perulangan dimana ia digunakan saja

-> Array []
    - Modify Element Through Index
    - Loop through elements
    - Multidimensional Array

-> Slice
    - Slice tidak memiliki sifat fixed-length yang berarti
      panjang dari slice tidak tetap sehingga kita bisa 
      dengan leluasa menentukan panjang dari slice nya
    - make function
      membuat sebuah slice dengan menggunakan fungsi make
    - append function
      Fungsi append akan mengembalikan nilai dari slice yang
      ditambahkannya, maka dari itu kita harus menyimpan
      fungsi append ke dalam suatu variable
    - append function with ellipsis (...)
      memasukkan seluruh element-element pada suatu array
      ke dalam array lainnya
    - copy function
      meng-copy seluruh element pada sebuah slice ke dalam slice lainnya
    - slicing [start:stop]
      Start sama dengan awal index yang ingin kita akses 
      dan stop berarti index akhirnya
    - Backing array
      menyimpan element pada slice, bukan slice nya sendiri
    - Cap function
      mengetahui kapasitas dari sebuah array maupun slice

-> Map
    - disimpan sebagai key:value pairs (pasang key dan value)
    - map[string]string >> tipe data key dari map harus string
                           dan value dari map juga harus string
    - Looping with map >> range loop
    - Deleting value >> menggunakan fungsi delete
    - Detecting a value >> mendeteksi keberadaan suatu value dengan cara
                           mengakses key dari map nya lalu memberikan dua
                           variable sebagai penampungnya
    - Combining slice with map

-> Aliase
    - digunakan sebagai nama alternative dari tipe data 
      yang sudah ada
    - byte >> uint8
    - rune >> uint32

-> Strings In Depth
    - slice of bytes
    - Looping Over String (byte-by-byte)
    - Looping Over String (rune-by-rune)
