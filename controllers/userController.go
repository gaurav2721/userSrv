package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	helpers "github.com/gaurav2721/userSrv/helpers"
	"github.com/gaurav2721/userSrv/models"

	"github.com/gaurav2721/userSrv/database"
	"github.com/gin-gonic/gin"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := helpers.ValidateUser(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()}) //TODO check this if error is suppossed to be this only
			return
		}

		insertErr := database.DbI.InsertUser(user)
		if insertErr != nil {
			msg := fmt.Sprintf("error user item was not created : %s", insertErr.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, user)

	}
}

// TODO: Have to implement pagination
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		users, err := database.DbI.GetAllUsers()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, *users)
		}
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		var _, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user *models.User
		user, err := database.DbI.GetUserFromId(userId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, *user)
		}
	}
}
