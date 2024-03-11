package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/wstiehler/zpedetailuser-api/internal/domain/detailuser"
	"github.com/wstiehler/zpedetailuser-api/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type ProjectApi struct {
	url        string
	httpClient *sling.Sling
}

func NewProjectApi(url string) *ProjectApi {
	return &ProjectApi{
		url:        url,
		httpClient: sling.New().Base(url),
	}
}

type HealthReturn struct {
	Status int `json:"status"`
}

type UserListResponse struct {
	Users []detailuser.UserDTO `json:"users"`
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

func (project *ProjectApi) GetUserByCriteria(criteria string, value string) (detailuser.UserListDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	response := new(detailuser.UserListDTO)

	req, err := http.NewRequest("GET", project.url+"/v1/user/"+criteria+"/"+value, nil)
	if err != nil {
		logger.Error("Failed to create request", zap.Error(err))
		return detailuser.UserListDTO{}, err
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Error("Request failed", zap.Error(err))
		return detailuser.UserListDTO{}, err
	}

	if err != nil {
		logger.Error("Request failed", zap.Error(err))
		return detailuser.UserListDTO{}, err
	}
	err = newcheckHasError(resp, err, *resp)

	if err != nil {
		logger.Error("Error on get user", zap.String("error", err.Error()))
		return detailuser.UserListDTO{}, err
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to read response body", zap.Error(err))
		return detailuser.UserListDTO{}, err
	}

	if err := json.Unmarshal(responseBody, response); err != nil {
		logger.Error("Failed to decode response body", zap.Error(err))
		return detailuser.UserListDTO{}, err
	}

	return *response, err
}

func newcheckHasError(resp *http.Response, err error, response http.Response) error {
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
