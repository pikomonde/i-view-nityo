package repository

import "github.com/pikomonde/i-view-nityo/model"

type User interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByID(id int64) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
}

type Invitation interface {
	CreateInvitation() (model.Invitation, error)
}
