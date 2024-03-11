package cmd

import (
	"fmt"

	"github.com/dghubble/sling"
	"github.com/wstiehler/zpecreateuser-api/internal/domain/createuser"
	"github.com/wstiehler/zpecreateuser-api/internal/infrastructure/logger"
)

type ErrorInfo struct {
	Message string `json:"msg"`
}

type ErrorResponse struct {
	Errors []ErrorInfo `json:"errors"`
}

type HealthReturn struct {
	Status int `json:"status"`
}

type ProjectApi struct {
	url string
}

func NewProjectApi(url string) *ProjectApi {

	project := &ProjectApi{
		url: url,
	}

	return project
}

func (project ProjectApi) ApiHealth() (*HealthReturn, error) {
	logger, dispose := logger.New()
	defer dispose()

	response := new(HealthReturn)

	resp, err := sling.New().Base(project.url).Path("health").ReceiveSuccess(response)

	if err != nil {
		logger.Error("Error")
		logger.Error(err.Error())
		return nil, err
	}
	fmt.Printf("[Health] result: %v\n", resp)
	return response, nil
}

func (project ProjectApi) CreateUser(r createuser.UserEntity) (*createuser.UserDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	response := new(createuser.UserDTO)

	resp, err := sling.New().Base(project.url).Post("/v1/user").BodyJSON(r).ReceiveSuccess(response)
	if err != nil {
		logger.Error("Create error")
		fmt.Println(response, resp, err)
		return nil, err
	}
	return response, nil
}
