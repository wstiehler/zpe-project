package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wstiehler/zpeupdateuser-api/internal/domain/user"
	"gorm.io/gorm"
)

func MakeAuthHandlers(r *gin.Engine, service user.Service, db *gorm.DB) {
	{
		r.POST("auth/login", Login(service, db))
	}
}

func Login(service user.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user user.UserLoginEntity

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userCompanyResponse, err := service.Login(db, user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", userCompanyResponse.Token, 3600*24, "", "", false, true)

		c.JSON(http.StatusOK, userCompanyResponse)
	}
}
