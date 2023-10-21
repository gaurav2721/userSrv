package database

import (
	"github.com/gaurav2721/userSrv/models"
)

var DbI DbInterface

type DbInterface interface {
	InsertUser(models.User) error
	GetUserFromId(user_id string) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
}
