package main

import (
	"fmt"
	"log"

	uuid "github.com/google/uuid"
	"github.com/wstiehler/zpeupdateuser-api/internal/api"
	"github.com/wstiehler/zpeupdateuser-api/internal/domain/user"
	"github.com/wstiehler/zpeupdateuser-api/internal/environment"
	config "github.com/wstiehler/zpeupdateuser-api/internal/infrastructure/database"
	"github.com/wstiehler/zpeupdateuser-api/internal/infrastructure/logger"
	"github.com/wstiehler/zpeupdateuser-api/internal/infrastructure/logger/logwrapper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	env := environment.GetInstance()
	zaplogger, dispose := logger.New()
	defer dispose()

	logger := logwrapper.New(&logwrapper.Zap{Logger: *zaplogger}).SetVersion(env.APP_VERSION)
	logger.Info("Starting Backend Application ZPEUpdateUser")

	RoutineID := uuid.New().String()

	err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	logger.Info("env",
		zap.String("MYSQL_DBNAME", env.MYSQL_DBNAME),
		zap.String("LOG_LEVEL", env.LOG_LEVEL),
		zap.String("ENVIRONMENT", env.ENVIRONMENT),
		zap.String("ROUTINE_ID", RoutineID),
	)

	mySqlConfig := config.DB

	defer func() {
		if r := recover(); r != nil {
			config.CloseConnection(mySqlConfig)
			logger.Fatal(fmt.Sprintf("Captured panic: %v", r))
		}
	}()

	repository := user.NewRepository(config.DB, user.MysqlAdapter{})

	service := user.NewService(repository)

	setupApi(logger, *service, *mySqlConfig)

	config.CloseConnection(mySqlConfig)

}

func setupApi(logger logwrapper.LoggerWrapper, service user.Service, db gorm.DB) {
	input := api.Input{
		Logger: logger,
	}

	api.Start(input, service, &db)
}
