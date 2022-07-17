package service

import (
	"context"

	"github.com/pikomonde/i-view-nityo/model"
	r "github.com/pikomonde/i-view-nityo/repository"
)

func NewServiceInvitation(
	ctx context.Context,
	config model.Config,
	rUser r.User,
	rInvitation r.Invitation,
) Invitation {
	return &ServiceInvitation{
		context:              ctx,
		config:               config,
		repositoryUser:       rUser,
		repositoryInvitation: rInvitation,
	}
}

type ServiceInvitation struct {
	context              context.Context
	config               model.Config
	repositoryUser       r.User
	repositoryInvitation r.Invitation
}

func (s *ServiceInvitation) CreateInvitation(userID int) error {
	return nil
}

func (s *ServiceInvitation) GetInvitations() ([]Invitation, error) {
	return nil, nil
}
