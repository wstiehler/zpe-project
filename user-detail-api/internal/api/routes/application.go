package routes

import (
	"errors"
	"net/http"

	"github.com/wstiehler/zpedetailuser-api/internal/domain/detailuser"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func MakeCreateUserHandlers(r *gin.Engine, service detailuser.Service, db *gorm.DB) {
	group := r.Group("/v1")
	{
		group.GET("/user/:criteria/:value", GetUserByInformation(service, db))
	}
}

func GetUserByInformation(service detailuser.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		criteria := c.Param("criteria")
		value := c.Param("value")

		users, err := service.GetUserByCriteria(db, criteria, value)
		if err != nil {
			if errors.Is(err, detailuser.ErrUserNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}
