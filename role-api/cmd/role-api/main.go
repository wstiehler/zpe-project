package main

import (
	"fmt"
	"log"

	uuid "github.com/google/uuid"
	"github.com/wstiehler/role-api/internal/api"
	"github.com/wstiehler/role-api/internal/domain/role"
	"github.com/wstiehler/role-api/internal/environment"
	config "github.com/wstiehler/role-api/internal/infrastructure/database"
	"github.com/wstiehler/role-api/internal/infrastructure/logger"
	"github.com/wstiehler/role-api/internal/infrastructure/logger/logwrapper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	env := environment.GetInstance()
	zaplogger, dispose := logger.New()
	defer dispose()

	logger := logwrapper.New(&logwrapper.Zap{Logger: *zaplogger}).SetVersion(env.APP_VERSION)
	logger.Info("Starting Backend Application ZPERole")

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

	err = config.DB.AutoMigrate(role.RoleEntity{}, role.PermissionEntity{})

	if err != nil {
		logger.Fatal(fmt.Sprintf("Captured panic: %v", err))
	}

	repository := role.NewRepository(config.DB, role.MysqlAdapter{})

	companyService := role.NewService(*repository)

	setupApi(logger, *companyService, *mySqlConfig)

	config.CloseConnection(mySqlConfig)

}

func setupApi(logger logwrapper.LoggerWrapper, roleService role.Service, db gorm.DB) {
	input := api.Input{
		Logger: logger,
	}

	api.Start(input, roleService, &db)
}
