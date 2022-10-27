package middlewares

import (
	"MyGram/database"
	"MyGram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		UserID := uint(userData["id"].(float64))
		User := models.User{}

		err := db.Where("id = ?", UserID).First(&User).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doest'n exist",
			})
			return
		}
		if User.ID != UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
		}
		c.Next()
	}
}
