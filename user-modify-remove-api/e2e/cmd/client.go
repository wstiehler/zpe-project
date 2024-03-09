package cmd

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"

// 	"github.com/dghubble/sling"
// 	"github.com/wstiehler/zpeupdateuser-api/internal/domain/createuser"
// 	"github.com/wstiehler/zpeupdateuser-api/internal/infrastructure/logger"
// 	"go.uber.org/zap"
// )

// type ErrorInfo struct {
// 	Message string `json:"msg"`
// }

// type ErrorResponse struct {
// 	Errors []ErrorInfo `json:"errors"`
// }

// type HealthReturn struct {
// 	Status int `json:"status"`
// }

// type ProjectApi struct {
// 	url string
// }

// func NewProjectApi(url string) *ProjectApi {

// 	project := &ProjectApi{
// 		url: url,
// 	}

// 	return project
// }

// func (project ProjectApi) ApiHealth() (*HealthReturn, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(HealthReturn)

// 	resp, err := sling.New().Base(project.url).Path("health").ReceiveSuccess(response)

// 	if err != nil {
// 		logger.Error("Error")
// 		logger.Error(err.Error())
// 		return nil, err
// 	}
// 	fmt.Printf("[Health] result: %v\n", resp)
// 	return response, nil
// }

// func (project ProjectApi) Signup(user createuser.UserEntity) (*createuser.UserSignupDTO, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(createuser.UserSignupDTO)

// 	resp, err := sling.New().Base(project.url).Post("/signup").BodyJSON(user).ReceiveSuccess(response)

// 	if err != nil {
// 		logger.Error("Create error")
// 		fmt.Println(response, resp, err)
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (project ProjectApi) Login(user createuser.UserEntity) (*createuser.UserCompanyResponseDTO, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(createuser.UserCompanyResponseDTO)

// 	token, err := sling.New().Base(project.url).Post("/login").BodyJSON(user).ReceiveSuccess(response)

// 	if err != nil {
// 		logger.Error("Login error")
// 		fmt.Println(response, token, err)
// 		return nil, err
// 	}

// 	return response, nil
// }

// func (project ProjectApi) CreateCompany(company createuser.CompanyEntity) (*createuser.CompanyDTO, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(createuser.CompanyDTO)

// 	resp, err := sling.New().Base(project.url).Post("/public/company").BodyJSON(company).ReceiveSuccess(response)
// 	if err != nil {
// 		logger.Error("Create error")
// 		fmt.Println(response, resp, err)
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (project ProjectApi) DeleteCompany(companyId string, token string) (*http.Response, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	req, err := http.NewRequest("DELETE", project.url+"/v1/company/"+companyId, nil)
// 	if err != nil {
// 		logger.Error("Failed to create request", zap.Error(err))
// 		return nil, err
// 	}

// 	req.Header.Add("Authorization", "Bearer "+token)

// 	httpClient := &http.Client{}
// 	resp, err := httpClient.Do(req)

// 	if err != nil {
// 		logger.Error("Request failed", zap.Error(err))
// 		return nil, err
// 	}
// 	err = newcheckHasError(resp, err, *resp)

// 	if err != nil {
// 		logger.Error("Error on delete company", zap.String("error", err.Error()))
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	return resp, err
// }

// func (project ProjectApi) CreateProduct(product createuser.ProductEntity, companyId string, token string) (*createuser.ProductDTO, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(createuser.ProductDTO)

// 	requestBody, err := json.Marshal(product)
// 	if err != nil {
// 		logger.Error("Failed to marshal product", zap.Error(err))
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", project.url+"/v1/product/"+companyId, bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		logger.Error("Failed to create request", zap.Error(err))
// 		return nil, err
// 	}

// 	req.Header.Add("Authorization", "Bearer "+token)

// 	httpClient := &http.Client{}
// 	resp, err := httpClient.Do(req)
// 	if err != nil {
// 		logger.Error("Request failed", zap.Error(err))
// 		return nil, err
// 	}
// 	err = newcheckHasError(resp, err, *resp)

// 	if err != nil {
// 		logger.Error("Error on create product", zap.String("error", err.Error()))
// 	}

// 	defer resp.Body.Close()

// 	responseBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Error("Failed to read response body", zap.Error(err))
// 		return nil, err
// 	}

// 	if err := json.Unmarshal(responseBody, response); err != nil {
// 		logger.Error("Failed to decode response body", zap.Error(err))
// 		return nil, err
// 	}

// 	return response, nil
// }

// func (project ProjectApi) GetProductsByCompanyId(companyId string, token string) (*CompanyListReturn, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(CompanyListReturn)

// 	req, err := http.NewRequest("GET", project.url+"/v1/products/"+companyId, nil)

// 	if err != nil {
// 		logger.Error("Failed to create request", zap.Error(err))
// 		return nil, err
// 	}

// 	req.Header.Add("Authorization", "Bearer "+token)

// 	httpClient := &http.Client{}
// 	resp, err := httpClient.Do(req)

// 	err = newcheckHasError(resp, err, *resp)

// 	if err != nil {
// 		logger.Error("Error on get product by companyID", zap.String("error", err.Error()))
// 	}

// 	defer resp.Body.Close()

// 	responseBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Error("Failed to read response body", zap.Error(err))
// 		return nil, err
// 	}

// 	if err := json.Unmarshal(responseBody, response); err != nil {
// 		logger.Error("Failed to decode response body", zap.Error(err))
// 		return nil, err
// 	}

// 	return response, nil
// }

// func (project ProjectApi) GetProductsActivesByCompanyId(companyId string) (*CompanyListReturn, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(CompanyListReturn)

// 	resp, err := sling.New().Base(project.url).Get("/public/products/actives/" + companyId).ReceiveSuccess(response)

// 	err = newcheckHasError(resp, err, *resp)

// 	if err != nil {
// 		logger.Error("Error on get product by companyID", zap.String("error", err.Error()))
// 		return nil, err
// 	}

// 	return response, nil
// }

// func (project ProjectApi) UpdateProduct(productId string, product createuser.ProductEntity, token string) (*createuser.ProductDTO, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(createuser.ProductDTO)

// 	requestBody, err := json.Marshal(product)
// 	if err != nil {
// 		logger.Error("Failed to marshal product", zap.Error(err))
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("PATCH", project.url+"/v1/product/"+productId, bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		logger.Error("Failed to create request", zap.Error(err))
// 		return nil, err
// 	}

// 	req.Header.Add("Authorization", "Bearer "+token)

// 	httpClient := &http.Client{}
// 	resp, err := httpClient.Do(req)
// 	if err != nil {
// 		logger.Error("Request failed", zap.Error(err))
// 		return nil, err
// 	}
// 	err = newcheckHasError(resp, err, *resp)

// 	if err != nil {
// 		logger.Error("Error on update product", zap.String("error", err.Error()))
// 	}

// 	defer resp.Body.Close()

// 	responseBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Error("Failed to read response body", zap.Error(err))
// 		return nil, err
// 	}

// 	if err := json.Unmarshal(responseBody, response); err != nil {
// 		logger.Error("Failed to decode response body", zap.Error(err))
// 		return nil, err
// 	}

// 	return response, nil
// }

// func (project ProjectApi) ActiveProduct(productId string, product createuser.ProductEntity, token string) (*createuser.ProductIsAvailableDTO, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(createuser.ProductIsAvailableDTO)

// 	requestBody, err := json.Marshal(product)
// 	if err != nil {
// 		logger.Error("Failed to marshal product", zap.Error(err))
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("PATCH", project.url+"/v1/product/active/"+productId, bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		logger.Error("Failed to create request", zap.Error(err))
// 		return nil, err
// 	}

// 	req.Header.Add("Authorization", "Bearer "+token)

// 	httpClient := &http.Client{}
// 	resp, err := httpClient.Do(req)
// 	if err != nil {
// 		logger.Error("Request failed", zap.Error(err))
// 		return nil, err
// 	}
// 	err = newcheckHasError(resp, err, *resp)

// 	if err != nil {
// 		logger.Error("Error on update product", zap.String("error", err.Error()))
// 	}

// 	defer resp.Body.Close()

// 	responseBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Error("Failed to read response body", zap.Error(err))
// 		return nil, err
// 	}

// 	if err := json.Unmarshal(responseBody, response); err != nil {
// 		logger.Error("Failed to decode response body", zap.Error(err))
// 		return nil, err
// 	}

// 	return response, nil
// }

// func (project ProjectApi) InactiveProduct(productId string, product createuser.ProductEntity, token string) (*createuser.ProductIsAvailableDTO, error) {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	response := new(createuser.ProductIsAvailableDTO)

// 	requestBody, err := json.Marshal(product)
// 	if err != nil {
// 		logger.Error("Failed to marshal product", zap.Error(err))
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("PATCH", project.url+"/v1/product/inactive/"+productId, bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		logger.Error("Failed to create request", zap.Error(err))
// 		return nil, err
// 	}

// 	req.Header.Add("Authorization", "Bearer "+token)

// 	httpClient := &http.Client{}
// 	resp, err := httpClient.Do(req)
// 	if err != nil {
// 		logger.Error("Request failed", zap.Error(err))
// 		return nil, err
// 	}
// 	err = newcheckHasError(resp, err, *resp)

// 	if err != nil {
// 		logger.Error("Error on update product", zap.String("error", err.Error()))
// 	}

// 	defer resp.Body.Close()

// 	responseBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Error("Failed to read response body", zap.Error(err))
// 		return nil, err
// 	}

// 	if err := json.Unmarshal(responseBody, response); err != nil {
// 		logger.Error("Failed to decode response body", zap.Error(err))
// 		return nil, err
// 	}

// 	return response, nil
// }

// func newcheckHasError(resp *http.Response, err error, response http.Response) error {
// 	logger, dispose := logger.New()
// 	defer dispose()

// 	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
// 		logger.Error("Unexpected response status", zap.Int("status", resp.StatusCode))
// 		return fmt.Errorf("unexpected response status: %d", resp.StatusCode)
// 	}

// 	if resp.StatusCode >= 500 {
// 		logger.Error("Unexpected response status", zap.Int("status", resp.StatusCode))
// 		return fmt.Errorf("unexpected response status: %d", resp.StatusCode)
// 	}

// 	return nil
// }
