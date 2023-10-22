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

var server *gin.Engine = SetUp()

func SetUp() *gin.Engine {
	server := gin.Default()
	database.DbI = database.InitMemoryDb()
	server.POST("/users/signup", controllers.SignUp())
	server.GET("/users/:user_id", controllers.GetUser())
	server.GET("/users", controllers.GetUsers())
	return server
}

func TestSignUp(t *testing.T) {

	input := models.User{
		Id:         "1",
		Name:       "Gaurav",
		SignupTime: 1234567,
	}
	bodyReq, _ := json.Marshal(input)

	req, _ := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBuffer(bodyReq))
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	//testing logs + assertion
	t.Log(resp.Result().StatusCode)
	t.Log(resp.Body)
	assert.Equal(t, 200, resp.Result().StatusCode)
}

// this tests a get when user is present in db
func TestGetPresentUser(t *testing.T) {
	//adding the user into the db
	input := models.User{
		Id:         "1",
		Name:       "Gaurav",
		SignupTime: 1234567,
	}
	bodyReq, _ := json.Marshal(input)

	req, _ := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBuffer(bodyReq))
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	//making request to fetch the user

	req, _ = http.NewRequest(http.MethodGet, "/users/1", nil)
	resp = httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	//testing logs + assertion
	t.Log(resp.Result().StatusCode)
	t.Log(resp.Body)
	assert.Equal(t, 200, resp.Result().StatusCode)
}

func TestGetAllUser(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	//testing logs + assertion
	t.Log(resp.Result().StatusCode)
	t.Log(resp.Body)
	assert.Equal(t, 200, resp.Result().StatusCode)
}
