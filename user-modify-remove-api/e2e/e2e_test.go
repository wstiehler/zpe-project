//go:build e2e
// +build e2e

package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wstiehler/zpeupdateuser-api/e2e/cmd"
	"github.com/wstiehler/zpeupdateuser-api/internal/domain/user"
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

	var token string

	userUpdate := user.UserEntity{
		Name: "William Teste AAAAAA",
	}

	var userID string

	userLogin := user.UserEntity{
		Email:    "teste-create@teste.com",
		Password: "teste-password",
	}

	t.Run("Login_When_return_must_be_success", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		userResponse, err := client.Login(userLogin)

		token = userResponse.Token
		userID = userResponse.Id

		assert.Equal(userResponse.Email, "teste-create@teste.com")
		assert.Equal(err, nil)
	})

	t.Run("Edit_When_return_must_be_success", func(t *testing.T) {
		client := cmd.NewProjectApi(url)

		user, err := client.UpdateUser(userID, userUpdate, token)

		assert.Equal(user.Name, "William Teste AAAAAA")
		assert.Equal(err, nil)
	})

	t.Run("DeleteUser_When_return_must_be_success", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		resp, err := client.DeleteUser(userID, token)

		assert.Equal(resp.StatusCode, 200)

		assert.Equal(err, nil)
	})

}
