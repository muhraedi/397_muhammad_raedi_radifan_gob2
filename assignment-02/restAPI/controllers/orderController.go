package controllers

import (
	"net/http"

	"restAPI/structs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type InDB struct {
	DB *gorm.DB
}

// create new data to the database
func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)
	c.ShouldBindJSON(&order)
	idb.DB.Create(&order)
	result = gin.H{
		"result": order,
	}
	c.JSON(http.StatusOK, result)
}

// get all data in orders
func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []structs.Order
		result gin.H
	)

	idb.DB.Preload("Items").Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}

	c.JSON(http.StatusOK, result)
}

// update data with {id} as query
func (idb *InDB) UpdateOrder(c *gin.Context) {
	var (
		order    structs.Order
		newOrder structs.Order
		result   gin.H
	)
	id := c.Param("orderId")
	c.ShouldBindJSON(&newOrder)
	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	order.CustomerName = newOrder.CustomerName
	order.OrderedAt = newOrder.OrderedAt
	order.Items = newOrder.Items
	err = idb.DB.Model(&newOrder).Updates(&order).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)
	id := c.Param("orderId")
	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Select("Items").Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
