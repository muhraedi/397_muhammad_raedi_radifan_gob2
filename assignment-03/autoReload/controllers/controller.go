package controllers

import (
	"autoReload/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type InDB struct {
	DB *gorm.DB
}

// create new data to the database
func (idb *InDB) CreateData(c *gin.Context) {
	var (
		condition models.Condition
		result    gin.H
	)

	water := c.PostForm("water")
	wind := c.PostForm("wind")

	intWater, _ := strconv.Atoi(water)
	intWind, _ := strconv.Atoi(wind)

	condition.Water = intWater
	condition.Wind = intWind

	idb.DB.Create(&condition)
	result = gin.H{
		"status": condition,
	}

	c.JSON(http.StatusOK, result)
}

// get all data
func (idb *InDB) GetAllData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	var (
		conditions models.Condition
		result     gin.H
	)

	err := idb.DB.Find(&conditions).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
			"status": nil,
		})
		return
	} else {
		result = gin.H{
			"status": gin.H{
				"water": conditions.Water,
				"wind":  conditions.Wind,
			},
		}
	}

	c.JSON(http.StatusOK, result)
}

// update data with {id}
func (idb *InDB) UpdateData(c *gin.Context) {
	id := c.Param("id")

	water := c.PostForm("water")
	wind := c.PostForm("wind")

	intWater, _ := strconv.Atoi(water)
	intWind, _ := strconv.Atoi(wind)

	var (
		condition    models.Condition
		newCondition models.Condition
		result       gin.H
	)

	err := idb.DB.Find(&condition, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newCondition.Water = intWater
	newCondition.Wind = intWind

	err = idb.DB.Model(&condition).Updates(newCondition).Error
	if err != nil {
		result = gin.H{
			"status": "update failed",
		}
	} else {
		result = gin.H{
			"status": "successfully updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (idb *InDB) DeleteData(c *gin.Context) {
	var (
		condition models.Condition
		result    gin.H
	)

	id := c.Param("id")

	err := idb.DB.Find(&condition, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = idb.DB.Delete(&condition).Error
	if err != nil {
		result = gin.H{
			"status": "delete failed",
		}
	} else {
		result = gin.H{
			"status": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
