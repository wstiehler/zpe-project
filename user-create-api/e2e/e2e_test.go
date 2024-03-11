//go:build e2e
// +build e2e

package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wstiehler/zpecreateuser-api/e2e/cmd"
	"github.com/wstiehler/zpecreateuser-api/internal/domain/createuser"
)

var url string
var timeout int

func readEnv() (string, int) {
	url = os.Getenv("APPLICATION_URL")
	timeoutString := os.Getenv("TEST_TIMEOUT")

	timeout, err := strconv.Atoi(timeoutString)
	if err != nil {
		timeout = 3
	}
	return url, timeout
}

func TestApiHealth(t *testing.T) {
	assert := assert.New(t)
	readEnv()

	t.Run("Health status", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		health, err := client.ApiHealth()

		fmt.Printf("health: %+v\n", health)
		assert.Equal(err, nil)
	})
}

func TestMethodCompany(t *testing.T) {
	assert := assert.New(t)
	readEnv()

	roleCreated := createuser.UserEntity{
		Name:     "William Teste Create",
		Email:    "teste-create@teste.com",
		Password: "teste-password",
		RoleId:   1,
	}

	t.Run("Createuser_When_return_must_be_success", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		userr, err := client.CreateUser(roleCreated)

		assert.Equal(userr.Email, "teste-create@teste.com")
		assert.Equal(err, nil)
	})

	t.Run("CreateUser_When_return_must_be_failure", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		_, err := client.CreateUser(roleCreated)

		assert.NotNil(t, err)
	})

}
