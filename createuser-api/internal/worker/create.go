package worker

import (
	"github.com/wstiehler/zpecreateuser-api/internal/domain/createuser"
	"github.com/wstiehler/zpecreateuser-api/internal/environment"
	"go.uber.org/zap"
)

type consumerWhoCreates struct{}

// PollingIntervalSeconds implements Consumer.
func (c *consumerWhoCreates) PollingIntervalSeconds() int64 {
	return environment.GetInstance().INTERVAL_GET_KEYS_TO_CREATE
}

// Handler implements Consumer.
func (c *consumerWhoCreates) Handler(input Input, user UserEntity) error {

	newUser := createuser.UserEntity{
		Name:  user.Name,
		Email: user.Email,
		// Role: createuser.RoleEntity{
		// 	Role: user.Role,
		// }
	}

	_, err := input.Service.CreateUser(input.ConfigDB, &newUser)

	if err != nil {
		input.Logger.Error("Error on create Product", zap.String("Error", err.Error()))
		return err
	}
	return nil
}

// URL implements Consumer.
func (c *consumerWhoCreates) QueueSubject() string {
	return environment.GetInstance().CREATE_USER_QUEUE_SUBJECT

}
