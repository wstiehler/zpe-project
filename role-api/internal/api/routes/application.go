package routes

import (
	"net/http"

	"github.com/wstiehler/role-api/internal/domain/role"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func MakeCreateUserHandlers(r *gin.Engine, service role.Service, db *gorm.DB) {

	group := r.Group("/v1")
	{
		group.POST("/role", CreateRole(service, db))
		group.POST("/permission", CreatePermission(service, db))
	}
}

func CreateRole(service role.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var role role.RoleEntity

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

func CreatePermission(service role.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var permission role.PermissionEntity

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
