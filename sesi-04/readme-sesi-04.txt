=========================== 
    Interface, Reflect, 
  Go-routines & Waitgroup
===========================

-> Interface
    - adalah tipe data pada bahasa Go yang merupakan kumpulan dari 
      definisi satu atau lebih method
    - teknik type assertion berfungsi untuk mengembalikan suatu tipe
      data interface kepada tipe data aslinya
    - Format penulisan type assertion >> namaVariable.(tipe data asli)

-> Empty interface
    - merupakan suatu tipe data yang dapat menerima segala tipe data
      pada bahasa Go
    - Map & Slice with Empty Interface >>
      memiliki value dengan tipe data yang berbeda-beda
      
-> Reflect
    - digunakan untuk melakukan inspeksi variabel, mengambil informasi 
      dari variabel tersebut atau bahkan memanipulasinya.
    - reflect.ValueOf() >>
      mengembalikan objek dalam tipe reflect.Value, yang berisikan informasi
      yang berhubungan dengan nilai pada variabel yang dicari
    - reflect.TypeOf() >>
      mengembalikan objek dalam tipe reflect.Type. Objek tersebut berisikan 
      informasi yang berhubungan dengan tipe data variabel yang dicari
    - Accessing Value Using Interface Method >> .Interface()
    - Identifying Method Information >>
      - Pengambilan refleksi method dilakukan menggunakan MethodByName dengan 
        argument adalah nama method yang diinginkan, atau bisa juga lewat 
        indeks method-nya (menggunakan Method(i))
      - Call() dipanggil untuk eksekusi method

-> Concurrency - Goroutines
    - Concurrency >> adalah mengeksekusi sebuah proses secara independen 
                     atau berhadapan dengan lebih dari satu tugas dalam
                     waktu yang sama
    - parallelism >> mengerjakan tugas yang banyak secara bersamaan 
                     dari awal hingga akhir
    - Goroutines >>
      sebuah thread ringan pada bahasa Go untuk melakukan concurrency
    - Goroutine bersifat asynchronous sehingga proses nya tidak saling 
      tunggu dengan Goroutine lainnya

-> Waitgroup
    - merupakan sebuah struct dari package sync, yang digunakan untuk 
      melakukan sinkronisasi terhadap go routine
    - sync.Waitgroup >> sebuah struct dari package sync
    - Add() >> menambah counter dari waitgroup
    - Done() >> memberitahu waitgroup tentang go routine yang telah
                menyelesaikan prosesnya
    - Wait() >> menunggu seluruh go routine menyelesaikan prosesnya
