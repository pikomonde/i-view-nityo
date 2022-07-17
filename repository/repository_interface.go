package repository

import "github.com/pikomonde/i-view-nityo/model"

type User interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByID(id int64) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
}

type Invitation interface {
	CreateInvitation(invitation model.Invitation) (model.Invitation, error)
	IsInvitationExist(invitationToken string) (bool, error)
	GetInvitations() ([]model.Invitation, error)
	DisableInvitation(invitationToken string) error
}
