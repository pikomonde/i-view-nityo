package service

import "github.com/pikomonde/i-view-nityo/model"

type Login interface {
	RegisterAdminIfNotExist() error
	LoginByUsernamePassword(username, password string) (string, error)
	LoginByInvitationToken(invitationToken string) error
}

type Invitation interface {
	CreateInvitation() (model.Invitation, error)
	GetInvitations() ([]model.Invitation, error)
	DisableInvitation(invitationToken string) error
}
