package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/wstiehler/zpecreateuser-api/internal/domain/createuser"
	"github.com/wstiehler/zpecreateuser-api/internal/infrastructure/logger"
	"go.uber.org/zap"
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

func (project ProjectApi) GetRoleByName(name string) (*createuser.RoleDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	response := new(createuser.RoleDTO)

	req, err := http.NewRequest("GET", project.url+"/v1/role/"+name, nil)

	if err != nil {
		logger.Error("Failed to create request", zap.Error(err))
		return nil, err
	}

	httpClient := &http.Client{}
	resp, _ := httpClient.Do(req)

	err = newcheckHasError(resp)

	if err != nil {
		logger.Error("Error on get role by name", zap.String("error", err.Error()))
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to read response body", zap.Error(err))
		return nil, err
	}

	if err := json.Unmarshal(responseBody, response); err != nil {
		logger.Error("Failed to decode response body", zap.Error(err))
		return nil, err
	}

	return response, nil
}

func newcheckHasError(resp *http.Response) error {
	logger, dispose := logger.New()
	defer dispose()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		logger.Error("Unexpected response status", zap.Int("status", resp.StatusCode))
		return fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	if resp.StatusCode >= 500 {
		logger.Error("Unexpected response status", zap.Int("status", resp.StatusCode))
		return fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	return nil
}
