package store

import (
	"matester/pkg/api"
	"matester/pkg/db"
)

type UsersController interface {
	List(limit int) []api.User
	Friends(userId int, limit int) []api.User
}

type UsersImpl struct {
	db db.Database
}

func NewUsersController(database db.Database) UsersController {
	return &UsersImpl{
		database,
	}
}

func (usersStore *UsersImpl) List(limit int) []api.User {
	//TODO: Pagination and limits
	return usersStore.db.QueryUsersList()
}

func (usersStore *UsersImpl) Friends(userId int, limit int) []api.User {
	//TODO: Pagination and limits
	return usersStore.db.QueryFriendsList(userId)
}
