package repository

import "github.com/pikomonde/i-view-nityo/model"

type User interface {
	GetUserByID(id int64) (model.User, error)
}

type Invitation interface {
	CreateInvitation() (model.Invitation, error)
}
