package main

import (
	"restAPI/controllers"
	"restAPI/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/order", inDB.CreateOrder)
	router.GET("/orders", inDB.GetOrders)
	router.PUT("/order/:orderId", inDB.UpdateOrder)
	router.DELETE("/order/:orderId", inDB.DeleteOrder)
	router.Run(":3000")
}
