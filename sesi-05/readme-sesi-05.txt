===========================
  Channels, Defer & Exit,
      Error Handling
===========================

-> Channels >> c := make(chan string)
    - merupakan sebuah mekanisme untuk Goroutine saling 
      berkomunikasi dengan Goroutine lainnya
    - c <- value >> Mengirim data kepada channel
    - result := <- c >> Menerima data dari channel
    - channel bersifat unbuffered atau tidak di buffer 
      di memori
    - c := make(chan int, 3) >> 
      argumen kedua merupakan kapasitas buffer
    - function time.Sleep >> memberikan jeda
    - Select >> merupakan sebuah fitur pada bahasa Go yang 
                memungkinkan kita untuk dapat menggunakan 
                lebih dari satu channel untuk proses komunikasi 
                antara Goroutine satu dengan yang lainnya

-> Defer & Exit
    - defer >>
      merupakan sebuah keyword pada bahasa Go yang digunakan
      untuk memanggil sebuah function yang dimana proses 
      eksekusi akan di tahan hingga block sebuah function selesai
    - Funcion exit yang yang berasal dari package os berguna 
      untuk menghentikan suatu program secara paksa

-> Error, Panic & Recover
    - error >> me-return sebuah error jika ada error yang terjadi
    - function strconv.Atoi >> mengkonversi tipe data string 
                               menjadi tipe data int
    - Error() >> menampilkan pesan jika terjadi sebuah error
    - errors.New >> pesan error yang dibuat
    - panic >> menampilkan stack trace error sekaligus menghentikan
               flow goroutine (karena main() juga merupakan goroutine,
               maka behaviour yang sama juga berlaku).
    - Function recover digunakan untuk menangkap error yang terjadi
      pada sebuah Goroutine