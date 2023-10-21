package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gaurav2721/userSrv/controllers"
	"github.com/gaurav2721/userSrv/database"
	"github.com/gaurav2721/userSrv/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {

	server := gin.Default()

	database.DbI = database.InitMemoryDb()

	input := models.User{
		Id:         "1",
		Name:       "Gaurav",
		SignupTime: 1234567,
	}
	bodyReq, _ := json.Marshal(input)

	server.POST("/users/signup", controllers.SignUp())

	req, _ := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBuffer(bodyReq))
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	//testing logs + assertion
	t.Log(resp.Result().StatusCode)
	t.Log(resp.Body)
	assert.Equal(t, 200, resp.Result().StatusCode)
}

func TestGetUser(t *testing.T) {

	server := gin.Default()

	database.DbI = database.InitMemoryDb()

	server.GET("/users/:user_id", controllers.GetUser())

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	//testing logs + assertion
	t.Log(resp.Result().StatusCode)
	t.Log(resp.Body)
	assert.Equal(t, 500, resp.Result().StatusCode)
}

func TestGetAllUser(t *testing.T) {

	server := gin.Default()

	database.DbI = database.InitMemoryDb()

	server.GET("/users", controllers.GetUsers())

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	//testing logs + assertion
	t.Log(resp.Result().StatusCode)
	t.Log(resp.Body)
	assert.Equal(t, 200, resp.Result().StatusCode)
}
