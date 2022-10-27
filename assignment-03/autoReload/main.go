package main

import (
	"autoReload/controllers"
	"autoReload/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/condition", inDB.CreateData)
	router.GET("/conditions", inDB.GetAllData)
	router.PUT("/condition/:id", inDB.UpdateData)
	router.DELETE("/condition/:id", inDB.DeleteData)
	router.Run(":3000")
}
