package database

import (
	"errors"

	"github.com/gaurav2721/userSrv/models"
)

type MemoryDb struct {
	usersMap map[string]models.User //key: id , value: user obj
}

func InitMemoryDb() *MemoryDb {
	return &MemoryDb{usersMap: make(map[string]models.User, 0)}
}

func (mdb *MemoryDb) InsertUser(user models.User) error {
	mdb.usersMap[user.Id] = user
	return nil
}

func (mdb *MemoryDb) GetAllUsers() (*[]models.User, error) {
	users := make([]models.User, 0)
	for _, value := range mdb.usersMap {
		users = append(users, value)
	}
	return &users, nil
}

func (mdb *MemoryDb) GetUserFromId(userId string) (*models.User, error) {
	user, ok := mdb.usersMap[userId]
	if ok {
		return &user, nil
	} else {
		return nil, errors.New("user id does not exist")
	}
}
