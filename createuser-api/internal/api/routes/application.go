package routes

import (
	"net/http"

	"github.com/wstiehler/zpecreateuser-api/internal/domain/createuser"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func MakeCreateUserHandlers(r *gin.Engine, service createuser.Service, db *gorm.DB) {

	publicGroup := r.Group("/v1")
	{
		publicGroup.POST("/user", CreateUser(service, db))
	}
}

func CreateUser(service createuser.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user createuser.UserEntity

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userCreated, err := service.CreateUser(db, &user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, userCreated)
	}
}
