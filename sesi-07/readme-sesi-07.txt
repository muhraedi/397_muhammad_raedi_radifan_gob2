============================
  SQL & API Implementation
============================
-> Database >>
	Komponen utama penyusun suatu app yang terdiri dari sekumpulan data 
	atau informasi yang tersimpan secara sistematis
	- Untuk memanipulasi data pada relational database digunakan sebuah 
	  bahasa yang disebut dengan Structured Query Language (SQL), oleh 
	  karena itu Relational Database juga disebut dengan SQL database
	- install >> PostgreSQL >> relational database manajemen system (RDBMS)
				 pgAdmin >> platform opensource yang digunakan untuk manajemen,
				 			pengembangan, serta manipulasi database posgreSQL
-> Go SQL
	- menggunakan package database/sql dengan mengintegrasikannya dengan 
	  sebuah sql driver. install driver PostgreSQL >>
	  go get -u github.com/lib/pq
	- membuat database >> CREATE db-go-sql;
	- membangun koneksi ke database >>
		- memberikan tanda underscore _ pada saat mengimport diver dari Postgresql
		- membuat sebuah variable global berupa konstanta yang kita gunakan untuk 
		  menyimpan seluruh informasi dari database Postgresql pada sistem kita
		- db *sql.DB >>
		  Variable db memiliki tipe data pointer dari struct sql.DB yang nantinya
		  akan kita reassign dengan object dari sql.DB pada saat kita membangun 
		  koneksi pada function main
		- menggabungkan seluruh info dari Postgresql yang telah kita buat pada 
		  line diatasnya menjadi sebuah string panjang yang kita simpan pada 
		  variable psqlInfo
		- menggunakan function Open yang berasal dari package database/sql
		- memanggil method Ping dapat membangun koneksi ke database sekaligus 
		  memeriksa jika string panjang berupa info yang kita berikan pada 
		  function Open merupakan info yang 100% valid
	- Create data
		- membuat terlebih dahulu sebuah struct
		- membuat sebuah sql statement yang merupakan sebuah statement untuk 
		  melakukan create data pada database
		- method QueryRow >> mengeksekusi sebuah query sql tergantung dari 
							 statement sql yang kita berikan
		- method QueryRow kita chaning dengan method Scan agar kita dapat 
		  mendapatkan nilai-nilai balikan dari statement
	- Get data
		- menggunakan method Query yang dimana method Query biasa digunakan 
		  untuk mendapatkan banyak data dari suatu table pada database 
		  dikarenakan method ini dapat mengembalikan satu atau lebih rows 
		  dari suatu table pada database
		- melakukan perulangan/looping sebanyak data yang kita dapatkan dengan 
		  acuan rows.Next. Method rows.Next akan menghasilkan nilai true selama
		  data nya masih ada
		- memanggil function Get pada function main
	- Update data
		- menggunakan method Exec untuk mengupdate data
		- menggunakan method RowsAffected untuk mengetahui berapa jumlah row 
		  atau data yang baru diupdate
		- Method RowsAffected didapatkan dari interface sql.Result yang 
		  merupakan nilai kembalian dari method Exec
		- memanggil function Update pada function main dengan meng-comment
		  pemanggilan function sebelumnya
	- Delete data
		- memanggil function Delete untuk menghapus data yang dibuat
		- memanggil function Delete pada function main dengan meng-comment
		  pemanggilan function sebelumnya
-> Gorm >> 
	ORM untuk bahasa Go, yang menyediakan berbagai macam fitur seperti 
	auto migration, eager loading, association, query method, dan lain-lain
	- install gorm >> go get -u gorm.io/gorm
	- models >> dibuat menggunakan struct
		- Untuk membuat structnya, maka seluruh property maupun nama dari 
		  structnya perlu diawali dengan huruf besar
		- Gorm menyediakan tags yang dapat kita gunakan
		- Gorm menyediakan berbagai macam assosiasi, seperti one to one,
		  one to many, dan many to many
	- Connecting To Database And Table Migration
		- install driver >> go get gorm.io/driver/postgres
		- menampung config untuk menghubungkan kepada database
		- method Open digunakan untuk membangun koneksi kepada database
		- Method Debug digunakan sebagai debugging atau logger. Kemudian 
		  di chaining dengan method AutoMigrate
		- Method AutoMigrate digunakan untuk memigrasi secara otomatis 
		  dari struct-struct yang telah dibuat
		- memanggil function StartDB pada func main
	- Create >> *gorm.DB
		- Method untuk CRUD pada Gorm dengan property Error agar kita bisa 
		  langsung dapat memeriksa errornya jika memang ada
	- Get One User Data >> menggunakan method First
		- Method First dapat menerima 3 parameter, parameter pertamanya 
		  adalah pointer terhadap data yang ingin dicari. Kemudian parameter
		  keduanya adalah condition dari query nya, dan yang terakhiir 
		  adalah data dari conditionnya
		- Method First akan mengembalikan error berupa ErrRecordNotFound 
		  jika tidak ada data yang ditemukan
	- Update
		- mengupdate data menggunakan Gorm, maka kita perlu menggunakan 
		  method Model terlebih dahulu agar hasil dari update nya dapat 
		  langsung discan sehingga kita dapat langsung mengetahui hasilnya
		- membuat condition dapat menggunakan method Where sehingga 
		  method Model dapat dichaining dengan method Where
	- Hooks
		- Method BeforeCreate akan tereksekusi sebelum melakukan create data
	- Eager Loading
		- menggunakan method Preload dan kita perlu memberikan nama table untuk parameter 
		  method Preload
		- parameter Preload harus menggunakan uppercase
	- Delete
		- akan menghapus data