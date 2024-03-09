//e2e
//go:build e2e
// +build e2e

package e2e

//package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/wstiehler/zpecreateuser-api/e2e/cmd"
// 	"github.com/wstiehler/zpecreateuser-api/internal/domain/candystore"
// )

// var url string
// var timeout int

// func readEnv() (string, int) {
// 	url = os.Getenv("APPLICATION_URL")
// 	timeoutString := os.Getenv("TEST_TIMEOUT")

// 	timeout, err := strconv.Atoi(timeoutString)
// 	if err != nil {
// 		timeout = 3
// 	}
// 	return url, timeout
// }

// func TestApiHealth(t *testing.T) {
// 	assert := assert.New(t)
// 	readEnv()

// 	t.Run("Health status", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)
// 		health, err := client.ApiHealth()

// 		fmt.Printf("health: %+v\n", health)
// 		assert.Equal(err, nil)
// 	})
// }

// func TestMethodCompany(t *testing.T) {
// 	assert := assert.New(t)
// 	readEnv()

// 	companyCreated := candystore.CompanyEntity{
// 		CompanyName: "Confeitaria da Anna",
// 	}

// 	var companyID string

// 	productCreated := candystore.ProductEntity{
// 		SKU:            55555,
// 		Title:          "Teste",
// 		Description:    "Teste",
// 		Price:          10,
// 		Installments:   5,
// 		AvailableTypes: "Teste",
// 		Information:    "500gr",
// 		IsAvailable:    true,
// 		IsSchedule:     false,
// 	}

// 	var productID string

// 	var token string

// 	t.Run("CreateCompany_When_return_must_be_success", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)
// 		company, err := client.CreateCompany(companyCreated)

// 		companyID = company.ID

// 		assert.Equal(company.CompanyName, "Confeitaria da Anna")
// 		assert.Equal(err, nil)
// 	})

// 	userCreated := candystore.UserEntity{
// 		Email:     "teste@teste.com",
// 		Password:  "teste@teste",
// 		Name:      "Laurecir Moreira",
// 		CompanyId: companyID,
// 	}

// 	userLogin := candystore.UserEntity{
// 		Email:    "teste@teste.com",
// 		Password: "teste@teste",
// 		Name:     "Laurecir Moreira",
// 	}

// 	t.Run("Signup_When_return_must_be_success", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)
// 		user, err := client.Signup(userCreated)

// 		assert.Equal(user.Email, "teste@teste.com")
// 		assert.Equal(err, nil)
// 	})

// 	t.Run("Login_When_return_must_be_success", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)
// 		userCompanyResponse, err := client.Login(userLogin)

// 		token = userCompanyResponse.User.Token

// 		assert.Equal(userCompanyResponse.User.Name, "Laurecir Moreira")
// 		assert.Equal(err, nil)
// 	})

// 	t.Run("CreateProduct_When_return_must_be_success", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)

// 		product, err := client.CreateProduct(productCreated, companyID, token)

// 		productID = product.Id

// 		assert.Equal(product.Title, "Teste")
// 		assert.Equal(err, nil)
// 	})

// 	productUpdate := candystore.ProductEntity{
// 		Id:             productID,
// 		SKU:            55555,
// 		Title:          "Novo Titulo",
// 		Description:    "Teste",
// 		Price:          10,
// 		Installments:   5,
// 		AvailableTypes: "Teste",
// 		Information:    "500gr",
// 		IsAvailable:    false,
// 		IsSchedule:     false,
// 		CompanyId:      companyID,
// 	}

// 	productActive := candystore.ProductEntity{
// 		Id:          productID,
// 		IsAvailable: true,
// 		CompanyId:   companyID,
// 	}

// 	productInactive := candystore.ProductEntity{
// 		Id:          productID,
// 		IsAvailable: false,
// 		CompanyId:   companyID,
// 	}

// 	t.Run("UpdateProduct_When_return_must_be_success", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)
// 		product, err := client.UpdateProduct(productID, productUpdate, token)

// 		assert.Equal(product.Title, "Novo Titulo")
// 		assert.Equal(err, nil)
// 	})

// 	t.Run("GetProductByCompanyId_When_return_must_be_one", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)

// 		products, err := client.GetProductsByCompanyId(companyID, token)

// 		assert.Equal(len(products.Items), 1)
// 		assert.Equal(err, nil)

// 	})

// 	t.Run("ActiveProduct_When_return_must_be_true", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)
// 		product, err := client.ActiveProduct(productID, productActive, token)

// 		assert.Equal(product.IsAvailable, true)
// 		assert.Equal(err, nil)
// 	})

// 	t.Run("GetProductActivesByCompanyId_When_return_must_be_one", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)

// 		products, err := client.GetProductsActivesByCompanyId(companyID)

// 		assert.Equal(len(products.Items), 1)
// 		assert.Equal(err, nil)
// 	})

// 	t.Run("InactiveProduct_When_return_must_be_false", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)
// 		product, err := client.InactiveProduct(productID, productInactive, token)

// 		assert.Equal(product.IsAvailable, false)
// 		assert.Equal(err, nil)
// 	})

// 	t.Run("GetProductActivesByCompanyId_When_return_must_be_zero", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)

// 		products, err := client.GetProductsActivesByCompanyId(companyID)

// 		assert.Equal(len(products.Items), 0)
// 		assert.Equal(err, nil)

// 	})

// 	t.Run("DeleteCompany_When_return_must_be_success", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)
// 		resp, err := client.DeleteCompany(companyID, token)

// 		assert.Equal(resp.StatusCode, 200)

// 		assert.Equal(err, nil)
// 	})

// }
