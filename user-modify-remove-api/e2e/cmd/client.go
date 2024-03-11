package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/wstiehler/zpeupdateuser-api/internal/domain/user"
	"github.com/wstiehler/zpeupdateuser-api/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type ErrorInfo struct {
	Message string `json:"msg"`
}

type ErrorResponse struct {
	Errors []ErrorInfo `json:"errors"`
}

type Token string

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

func (project ProjectApi) Login(u user.UserEntity) (*user.UserLoginDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	response := new(user.UserLoginDTO)

	token, err := sling.New().Base(project.url).Post("/auth/login").BodyJSON(u).ReceiveSuccess(response)

	if err != nil {
		logger.Error("Login error")
		fmt.Println(response, token, err)
		return nil, err
	}

	return response, nil
}

func (project ProjectApi) DeleteUser(id string, token string) (*http.Response, error) {
	logger, dispose := logger.New()
	defer dispose()

	req, err := http.NewRequest("DELETE", project.url+"/v1/user/"+id, nil)
	if err != nil {
		logger.Error("Failed to create request", zap.Error(err))
		return nil, err
	}

	cookie := &http.Cookie{
		Name:  "Authorization",
		Value: token,
		Path:  "/",
	}

	req.AddCookie(cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		logger.Error("Request failed", zap.Error(err))
		return nil, err
	}
	err = newcheckHasError(resp)

	if err != nil {
		logger.Error("Error on delete company", zap.String("error", err.Error()))
		return nil, err
	}

	defer resp.Body.Close()

	return resp, err
}

func (project ProjectApi) UpdateUser(userId string, u user.UserEntity, token string) (*user.UserDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	response := new(user.UserDTO)

	requestBody, err := json.Marshal(u)
	if err != nil {
		logger.Error("Failed to marshal user", zap.Error(err))
		return nil, err
	}

	req, err := http.NewRequest("PATCH", project.url+"/v1/user/"+userId, bytes.NewBuffer(requestBody))
	if err != nil {
		logger.Error("Failed to create request", zap.Error(err))
		return nil, err
	}

	cookie := &http.Cookie{
		Name:  "Authorization",
		Value: token,
		Path:  "/",
	}

	req.AddCookie(cookie)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Error("Request failed", zap.Error(err))
		return nil, err
	}
	err = newcheckHasError(resp)

	if err != nil {
		logger.Error("Error on update product", zap.String("error", err.Error()))
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
