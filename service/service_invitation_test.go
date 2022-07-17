package service_test

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/pikomonde/i-view-nityo/model"
	"github.com/pikomonde/i-view-nityo/repository"
	"github.com/pikomonde/i-view-nityo/service"
	"github.com/stretchr/testify/assert"
)

func TestNewServiceInvitation(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	service.NewServiceInvitation(ctx, model.Config{},
		repository.NewMockUser(mockCtrl),
		repository.NewMockInvitation(mockCtrl),
	)

	mockCtrl.Finish()
}

func TestCreateInvitation_ErrorCheckExist(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceInvitation(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoInvitation.EXPECT().IsInvitationExist(gomock.Any()).Return(false, errors.New("")),
	)

	_, err := s.CreateInvitation()
	mockCtrl.Finish()
	assert.Equal(t, errors.New(""), err)
}

func TestCreateInvitation_Success(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceInvitation(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoInvitation.EXPECT().IsInvitationExist(gomock.Any()).Return(false, nil),
		mockRepoInvitation.EXPECT().CreateInvitation(gomock.Any()),
	)

	_, _ = s.CreateInvitation()
	mockCtrl.Finish()
}

func TestCreateInvitation_SuccessWithInvitationExistOnce(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceInvitation(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoInvitation.EXPECT().IsInvitationExist(gomock.Any()).Return(true, nil),
		mockRepoInvitation.EXPECT().IsInvitationExist(gomock.Any()).Return(false, nil),
		mockRepoInvitation.EXPECT().CreateInvitation(gomock.Any()),
	)

	_, _ = s.CreateInvitation()
	mockCtrl.Finish()
}

func TestGetInvitations_Success(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceInvitation(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoInvitation.EXPECT().GetInvitations(),
	)

	_, _ = s.GetInvitations()
	mockCtrl.Finish()
}

func TestDisableInvitation_Success(t *testing.T) {
	ctx := context.Background()
	invitationToken := "xxxxxxxxxx"

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceInvitation(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoInvitation.EXPECT().UpdateInvitationStatus(invitationToken, model.InvitationStatus_Disabled),
	)

	_ = s.DisableInvitation(invitationToken)
	mockCtrl.Finish()
}
