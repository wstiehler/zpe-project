package routes

import (
	"net/http"
	"strconv"

	"github.com/wstiehler/role-api/internal/domain/role"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func MakeRoleHandlers(r *gin.Engine, service role.Service, db *gorm.DB) {

	group := r.Group("/v1")
	{
		group.POST("/role", CreateRole(service, db))
		group.GET("/role/:id", GetRoleByID(service, db))
		group.POST("/permission", CreatePermission(service, db))
		group.GET("/permission/:id", GetPermissionByRoleID(service, db))
	}
}

func CreateRole(service role.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var role role.RoleEntity

		if err := c.BindJSON(&role); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		roleCreated, err := service.CreateRole(db, &role)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
			return
		}

		c.JSON(http.StatusCreated, roleCreated)
	}
}

func CreatePermission(service role.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var permission role.PermissionEntity

		if err := c.BindJSON(&permission); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		permissionCreated, err := service.CreatePermission(db, &permission)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create permission"})
			return
		}

		c.JSON(http.StatusCreated, permissionCreated)
	}
}

func GetRoleByID(service role.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		idInt, _ := strconv.ParseUint(id, 10, 64)

		var uintNum uint = uint(idInt)

		responseView, err := service.GetRoleByID(uintNum, db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get role"})
			return
		}

		c.JSON(http.StatusOK, responseView)
	}

}

func GetPermissionByRoleID(service role.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		idInt, _ := strconv.ParseUint(id, 10, 64)

		var uintNum uint = uint(idInt)

		responseView, err := service.GetPermissionsByRoleID(uintNum, db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get permissions"})
			return
		}

		c.JSON(http.StatusOK, responseView)
	}

}
