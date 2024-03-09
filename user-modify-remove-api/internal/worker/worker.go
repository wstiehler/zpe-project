package worker

import (
	"github.com/wstiehler/zpeupdateuser-api/internal/domain/user"
	"github.com/wstiehler/zpeupdateuser-api/internal/infrastructure/logger/logwrapper"
	"gorm.io/gorm"
)

type Input struct {
	Logger   logwrapper.LoggerWrapper
	Service  user.Service
	ConfigDB *gorm.DB
}

func Start(input Input, companyService user.Service) {

	go createPolling(input, new(consumerWhoCreates))
}
