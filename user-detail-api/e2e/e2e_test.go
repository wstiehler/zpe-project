package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wstiehler/zpedetailuser-api/e2e/cmd"
	"github.com/wstiehler/zpedetailuser-api/internal/domain/detailuser"
)

var url string

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

	roleCreated := detailuser.UserEntity{
		Name:     "William Teste Create",
		Email:    "teste-create@teste.com",
		Password: "teste-password",
		RoleId:   8,
	}
	t.Run("Getuserbyemail_When_return_must_be_success", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		users, err := client.GetUserByCriteria("email", roleCreated.Email)

		assert.Nil(t, err)
		assert.NotEmpty(t, users)
		assert.Equal(t, "teste-create@teste.com", users.Users[0].Email)
	})

	time.Sleep(1000)

	t.Run("Getuserbyname_When_return_must_be_success", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		users, err := client.GetUserByCriteria("name", roleCreated.Name)

		assert.Nil(t, err)
		assert.NotEmpty(t, users)
		assert.Equal(t, "William Teste Create", users.Users[0].Name)
	})

	time.Sleep(1000)

	t.Run("GetallUser_When_return_must_be_success", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		users, err := client.GetUserByCriteria("name", "aaa")

		assert.Nil(t, err)
		assert.NotEmpty(t, users)
		assert.Len(t, 1, users.Users[0])
	})
}
