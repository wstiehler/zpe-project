package routes

import (
	"net/http"

	"github.com/wstiehler/zpecreateuser-api/internal/domain/createuser"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func MakeCreateUserHandlers(r *gin.Engine, service createuser.Service, db *gorm.DB) {

	group := r.Group("/v1")
	{
		group.POST("/user", CreateUser(service, db))
		group.GET("/role/:name", GetRoleByName(service, db))
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

func GetRoleByName(service createuser.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		name := c.Param("name")

		responseView, err := service.GetRoleByName(name, db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get role"})
			return
		}

		c.JSON(http.StatusOK, responseView)
	}

}
