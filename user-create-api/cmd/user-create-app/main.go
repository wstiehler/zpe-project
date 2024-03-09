package main

import (
	"fmt"
	"log"

	uuid "github.com/google/uuid"
	"github.com/wstiehler/zpecreateuser-api/internal/api"
	"github.com/wstiehler/zpecreateuser-api/internal/domain/createuser"
	"github.com/wstiehler/zpecreateuser-api/internal/environment"
	config "github.com/wstiehler/zpecreateuser-api/internal/infrastructure/database"
	"github.com/wstiehler/zpecreateuser-api/internal/infrastructure/logger"
	"github.com/wstiehler/zpecreateuser-api/internal/infrastructure/logger/logwrapper"
	"github.com/wstiehler/zpecreateuser-api/internal/worker"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	env := environment.GetInstance()
	zaplogger, dispose := logger.New()
	defer dispose()

	logger := logwrapper.New(&logwrapper.Zap{Logger: *zaplogger}).SetVersion(env.APP_VERSION)
	logger.Info("Starting Backend Application ZPECreateUser")

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

	err = config.DB.AutoMigrate(createuser.UserEntity{})

	if err != nil {
		logger.Fatal(fmt.Sprintf("Captured panic: %v", err))
	}

	repository := createuser.NewRepository(config.DB, createuser.MysqlAdapter{})

	service := createuser.NewService(*repository)

	setupWorker(logger, *service, *mySqlConfig)

	setupApi(logger, *service, *mySqlConfig)

	config.CloseConnection(mySqlConfig)

}

func setupApi(logger logwrapper.LoggerWrapper, service createuser.Service, db gorm.DB) {
	input := api.Input{
		Logger: logger,
	}

	api.Start(input, service, &db)
}

func setupWorker(logger logwrapper.LoggerWrapper, service createuser.Service, db gorm.DB) {
	input := worker.Input{
		Logger:   logger,
		ConfigDB: &db,
	}
	worker.Start(input, service)
}
