package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wstiehler/zpedetailuser-api/internal/api/middlewares"
	"github.com/wstiehler/zpedetailuser-api/internal/api/routes"
	"github.com/wstiehler/zpedetailuser-api/internal/domain/detailuser"
	"github.com/wstiehler/zpedetailuser-api/internal/environment"
	"github.com/wstiehler/zpedetailuser-api/internal/infrastructure/logger/logwrapper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Input struct {
	Logger logwrapper.LoggerWrapper
}

func Start(input Input, companyService detailuser.Service, db *gorm.DB) {
	r := gin.New()
	env := environment.GetInstance()

	logger := input.Logger

	logger.Info("Starting ZPECreateUser-API")

	applicationPort := resolvePort()

	r.Use(middlewares.Context())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.Logger(logger))

	if !env.IsDevelopment() {
		r.Use(middlewares.Recovery(&zap.Logger{}, true))
	}

	r.SetTrustedProxies([]string{env.APPLICATION_ADDRESS})
	routes.MakeCreateUserHandlers(r, companyService, db)
	routes.MakeHealthHandle(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	})

	if err := r.Run(applicationPort); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}

}

func resolvePort() string {
	const CHAR string = ":"
	env := environment.GetInstance()
	port := env.APPLICATION_PORT
	fisrtChar := port[:1]
	if fisrtChar != CHAR {
		port = CHAR + port
	}
	return port
}
