package store

import (
	"matester/pkg/api"
	"matester/pkg/db"
)

type Users interface {
	List(limit int) ([]api.User, error)
}

type UsersImpl struct {
	db db.Database
}

func NewUsersProvider(database db.Database) Users {
	return &UsersImpl{
		database,
	}
}

func (usersStore *UsersImpl) List(limit int) ([]api.User, error) {
	var res []api.User
	for i := 0; i < limit; i++ {
		row := usersStore.db.Next()
		res = append(res, api.User{
			Name: row.Value,
		})
	}

	return res, nil
}