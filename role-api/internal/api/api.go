package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wstiehler/role-api/internal/api/middlewares"
	"github.com/wstiehler/role-api/internal/api/routes"
	"github.com/wstiehler/role-api/internal/domain/role"
	"github.com/wstiehler/role-api/internal/environment"
	"github.com/wstiehler/role-api/internal/infrastructure/logger/logwrapper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Input struct {
	Logger logwrapper.LoggerWrapper
}

func Start(input Input, roleService role.Service, db *gorm.DB) {
	r := gin.New()
	env := environment.GetInstance()

	logger := input.Logger

	logger.Info("Starting ZPERole-API")

	applicationPort := resolvePort()

	r.Use(middlewares.Context())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.Logger(logger))

	if !env.IsDevelopment() {
		r.Use(middlewares.Recovery(&zap.Logger{}, true))
	}

	r.SetTrustedProxies([]string{env.APPLICATION_ADDRESS})
	routes.MakeRoleHandlers(r, roleService, db)
	routes.MakeHealthHandle(r)

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
