package worker

import (
	"github.com/wstiehler/zpecreateuser-api/internal/domain/createuser"
	"github.com/wstiehler/zpecreateuser-api/internal/infrastructure/logger/logwrapper"
	"gorm.io/gorm"
)

type Input struct {
	Logger   logwrapper.LoggerWrapper
	Service  createuser.Service
	ConfigDB *gorm.DB
}

func Start(input Input, service createuser.Service) {

	go createPolling(input, new(consumerWhoCreates))
}
