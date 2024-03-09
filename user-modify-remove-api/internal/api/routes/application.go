package routes

import (
	"net/http"

	"github.com/wstiehler/zpeupdateuser-api/internal/api/middlewares"
	"github.com/wstiehler/zpeupdateuser-api/internal/domain/user"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func MakeUpdateUserHandlers(r *gin.Engine, service user.Service, db *gorm.DB) {

	group := r.Group("/v1")
	{
		group.PATCH("/user/:id", middlewares.RequireAuthAndPermission("modifier"), UpdateUser(service, db))
		group.DELETE("/user/:id", middlewares.RequireAuthAndPermission("deleter"), DeleteUser(service, db))
	}
}

func UpdateUser(service user.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user *user.UserEntity

		id := c.Param("id")

		user, err := service.GetUserByID(id, db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on get user"})
			return
		}

		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Json invalid data"})
		}

		userDTO, err := service.UpdateUser(id, *user, db)

		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Error on update user"})
			return
		}

		c.JSON(http.StatusOK, userDTO)
	}
}

func DeleteUser(service user.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		err := service.DeleteUser(id, db)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"user deleted": id})
		}
	}

}
