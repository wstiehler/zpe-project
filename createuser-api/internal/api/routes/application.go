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
		group.POST("/role", CreateRole(service, db))
		group.POST("/permission", CreatePermission(service, db))
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

func CreateRole(service createuser.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var role createuser.RoleEntity

		if err := c.BindJSON(&role); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		roleCreated, err := service.CreateRole(db, &role)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, roleCreated)
	}
}

func CreatePermission(service createuser.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var permission createuser.PermissionEntity

		if err := c.BindJSON(&permission); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		permissionCreated, err := service.CreatePermission(db, &permission)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, permissionCreated)
	}
}
