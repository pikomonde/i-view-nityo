package service

import (
	"context"

	"github.com/pikomonde/i-view-nityo/model"
	r "github.com/pikomonde/i-view-nityo/repository"
)

func NewServiceLogin(
	ctx context.Context,
	config model.Config,
	rUser r.User,
	rInvitation r.Invitation,
) Login {
	return &ServiceLogin{
		context:              ctx,
		config:               config,
		repositoryUser:       rUser,
		repositoryInvitation: rInvitation,
	}
}

type ServiceLogin struct {
	context              context.Context
	config               model.Config
	repositoryUser       r.User
	repositoryInvitation r.Invitation
}

func (s *ServiceLogin) LoginByUsernamePassword(username, password string) error {
	return nil
}

func (s *ServiceLogin) LoginByInvitationID(invitationID string) error {
	return nil
}
