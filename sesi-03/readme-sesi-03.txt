==================================
    Function, Method, Pointer, 
   Struct & Exported-Unexported
==================================

-> Function
    - menggunakan keyword func lalu diikuti dengan nama function
      dan parameter yang digunakan
    - Ketika function lebih dari satu parameter yang memiliki
      tipe data sama, maka cukup menuliskan satu
    - return >> mengembalikan sebuah nilai
    - dapat mengembalikan multi value >> (float64, float64)
    - Predefined return value
      mendefinisikan nilai return dari sebuah function
    - Variadic function >>  dapat menerima argumen yang 
                            tak terbatas jumlahnya

-> Closure
    - merupakan sebuah anonymous function atau function tanpa nama
      yang dapat disimpan sebagai sebuah variable maupun dapat 
      dijadikan sebagai nilai return dari sebuah function
    - IIFE (immediately-invoked function expression) >> 
      merupakan sebuah closure yang dapat langsung tereksekusi 
      ketika pertama kali dideklarasikan
    - Closure juga bisa dijadikan sebagai nilai kembalian 
      dari suatu function
    - Callback adalah sebuah closure yang dijadikan sebagai 
      parameter pada sebuah function

-> Pointer
    - merupakan sebuah variable spesial yang digunakan untuk 
      menyimpan alamat memori variable lainnya
    - Memory address >>
      - mendapatkan alamat memori dari variable biasa dengan
        menggunakan tanda ampersand &
      - mendapatkan nilai asli dari variable pointer dengan 
        cara menggunakan tanda asterisk *
    - Changing value through pointer
      Pointer digunakan untuk menyimpan alamat memori, maka 
      ketika kita mengganti nilai dari sebuah pointer, maka
      variable lainnya yang mempunyai alamat memori yang 
      sama, akan ikut terganti nilainya
    - Pointer juga bisa dijadikan sebagai sebuah parameter 
      pada sebuah function

-> Struct
    - adalah sebuah tipe data berupa kumpulan/koleksi dari 
      berbagai macam property/field dan juga method
    - Giving value to struct
      memberikan nilai kepada field yang terdapat pada sebuah
      struct, kita perlu terlebih dahulu menyimpan tipe data 
      dari struct kepada sebuah variable
    - Initializing struct
      menginisialisasi sebuah struct sekaligus memberikan
      nilai-nilai nya
    - Pointer to a struct
      menggunakan pointer pada sebuah struct
    - Embedded struct
      mengandung tipe data struct lainnya dengan menjadikannya
      sebuah field
    - Anonymous struct
      sebuah struct yang tidak dideklerasikan di awal sebagai
      sebuah tipe data struct baru, melainkan langsung dideklerasikan 
      bersamaan dengan pembuatan variable
    - Slice of struct
      Slice juga dapat dikombinasikan dengan tipe data struct, cara 
      penulisannya mirip seperti slice of map
    - Slice of anonymous struct
      Anonymous struct juga dapat digabungkan dengan tipe data 
      slice dan pengisian nilainya pun dapat dilakukan secara langsung

-> Exported & Unexported
    - Di Go tidak ada istilah public modifier dan private modifier
    - exported ekuivalen dengan public modifier
    - unexported untuk private modifier