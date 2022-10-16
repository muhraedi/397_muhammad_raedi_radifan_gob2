Nama lengkap        : Muhammad Raedi Radifan
kode peserta        : 149368582100-397
tugas assignment 2  : REST API

panduan aplikasi :
- membuat database mysql dengan nama "ass02db"
- masuk ke terminal
- masuk ke directory "assignment-02/restAPI"
- jalankan perintah go run main.go
- allow firewall yang muncul
- buka postman kemudian masukkan url sesuai router

Url postman:
POST http://localhost:3000/order
Contoh request body :
                        {
                            "orderedAt":"2019-11-09T21:21:46+00:00",
                            "customerName":"Tom Jerry",
                            "items":[
                                {
                                    "itemCode":"123",
                                    "description":"IPhone 10x",
                                    "quantity":1
                                }
                            ]
                        }
GET http://localhost:3000/orders
PUT http://localhost:3000/order/:orderId
Contoh request body :
                        {
                            "customerName":"Spyke Tyke",
                            "orderedAt":"2019-11-09T21:21:46Z",
                            "items":[
                                {
                                    "lineItemId":1,
                                    "itemCode":"123",
                                    "description":"IPhone 10x",
                                    "quantity":10
                                }
                            ]
                        }
DELETE http://localhost:3000/order/:orderId
Link postman:
    https://github.com/muhraedi/397_muhammad_raedi_radifan_gob2/tree/master/assignment-02/postman