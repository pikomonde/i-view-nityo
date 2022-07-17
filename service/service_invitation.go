package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/pikomonde/i-view-nityo/helper"
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

func (s *ServiceInvitation) CreateInvitation() (model.Invitation, error) {
	rnd := rand.NewSource(time.Now().UnixNano())
	token := ""
	for {
		digit := 6 + rnd.Int63()%7 // 6-12
		token = helper.RandomString(int(digit))
		isExist, err := s.repositoryInvitation.IsInvitationExist(token)
		if err != nil {
			return model.Invitation{}, err
		}

		if !isExist {
			break
		}
	}

	invitation := model.Invitation{
		Token:     token,
		Status:    model.InvitationStatus_Inactive,
		CreatedAt: time.Now().UnixNano(),
	}

	createdInvitation, err := s.repositoryInvitation.CreateInvitation(invitation)

	return createdInvitation, err
}

func (s *ServiceInvitation) GetInvitations() ([]model.Invitation, error) {
	return s.repositoryInvitation.GetInvitations()
}

func (s *ServiceInvitation) DisableInvitation(invitationToken string) error {
	return s.repositoryInvitation.UpdateInvitationStatus(invitationToken, model.InvitationStatus_Disabled)
}
