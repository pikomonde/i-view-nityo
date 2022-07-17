package service_test

import (
	"context"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/pikomonde/i-view-nityo/helper"
	"github.com/pikomonde/i-view-nityo/model"
	"github.com/pikomonde/i-view-nityo/repository"
	"github.com/pikomonde/i-view-nityo/service"
	"github.com/stretchr/testify/assert"
)

func TestNewServiceLogin(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	service.NewServiceLogin(ctx, model.Config{},
		repository.NewMockUser(mockCtrl),
		repository.NewMockInvitation(mockCtrl),
	)

	mockCtrl.Finish()
}

func TestRegisterAdminIfNotExist_AdminExist(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceLogin(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoUser.EXPECT().GetUserByID(int64(1)).Return(model.User{}, nil),
	)

	s.RegisterAdminIfNotExist()
	mockCtrl.Finish()
}

func TestRegisterAdminIfNotExist_Success(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceLogin(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoUser.EXPECT().GetUserByID(int64(1)).Return(model.User{}, model.Err_Repository_User_NotFound),
		mockRepoUser.EXPECT().CreateUser(gomock.Any()).Return(model.User{}, nil),
	)

	s.RegisterAdminIfNotExist()
	mockCtrl.Finish()
}

func TestLoginByUsernamePassword_UserNotExist(t *testing.T) {
	ctx := context.Background()
	username := "username"
	password := "password"

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceLogin(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoUser.EXPECT().GetUserByUsername(username).Return(model.User{}, model.Err_Repository_User_NotFound),
	)

	_, err := s.LoginByUsernamePassword(username, password)
	mockCtrl.Finish()
	assert.Equal(t, model.Err_Repository_User_NotFound, err)
}

func TestLoginByUsernamePassword_WrongPassword(t *testing.T) {
	ctx := context.Background()
	username := "username"
	password := "password"

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceLogin(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoUser.EXPECT().GetUserByUsername(username).Return(model.User{}, nil),
	)

	_, err := s.LoginByUsernamePassword(username, password)
	mockCtrl.Finish()
	assert.Equal(t, model.Err_Service_Login_WrongPassword, err)
}

func TestLoginByUsernamePassword_Success(t *testing.T) {
	ctx := context.Background()
	username := "username"
	password := "password"
	user := model.User{
		Password: password,
	}
	user.Password = helper.HashUserPassword(user)

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceLogin(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoUser.EXPECT().GetUserByUsername(username).Return(user, nil),
	)

	_, err := s.LoginByUsernamePassword(username, password)
	mockCtrl.Finish()
	assert.Equal(t, nil, err)
}

func TestLoginByInvitationToken_InvitationNotExist(t *testing.T) {
	ctx := context.Background()
	invitationToken := "xxxxxxxxxx"

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceLogin(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoInvitation.EXPECT().GetInvitationByToken(invitationToken).Return(model.Invitation{}, model.Err_Repository_Invitation_InvalidToken),
	)

	_, err := s.LoginByInvitationToken(invitationToken)
	mockCtrl.Finish()
	assert.Equal(t, model.Err_Repository_Invitation_InvalidToken, err)
}

func TestLoginByInvitationToken_InvitationExpired(t *testing.T) {
	ctx := context.Background()
	invitationToken := "xxxxxxxxxx"

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceLogin(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoInvitation.EXPECT().GetInvitationByToken(invitationToken).Return(model.Invitation{}, nil),
	)

	_, err := s.LoginByInvitationToken(invitationToken)
	mockCtrl.Finish()
	assert.Equal(t, model.Err_Service_Login_ExpiredInvitationToken, err)
}

func TestLoginByInvitationToken_InvitationStatusDisabled(t *testing.T) {
	ctx := context.Background()
	invitationToken := "xxxxxxxxxx"
	now := time.Now()

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceLogin(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoInvitation.EXPECT().GetInvitationByToken(invitationToken).Return(model.Invitation{
			Status:    model.InvitationStatus_Disabled,
			CreatedAt: now.UnixNano(),
		}, nil),
	)

	_, err := s.LoginByInvitationToken(invitationToken)
	mockCtrl.Finish()
	assert.Equal(t, model.Err_Service_Login_DisabledInvitationToken, err)
}

func TestLoginByInvitationToken_Success(t *testing.T) {
	ctx := context.Background()
	invitationToken := "xxxxxxxxxx"
	now := time.Now()

	mockCtrl := gomock.NewController(t)
	mockRepoUser := repository.NewMockUser(mockCtrl)
	mockRepoInvitation := repository.NewMockInvitation(mockCtrl)
	s := service.NewServiceLogin(ctx, model.Config{},
		mockRepoUser,
		mockRepoInvitation,
	)

	gomock.InOrder(
		mockRepoInvitation.EXPECT().GetInvitationByToken(invitationToken).Return(model.Invitation{
			CreatedAt: now.UnixNano(),
		}, nil),
		mockRepoInvitation.EXPECT().UpdateInvitationStatus(invitationToken, model.InvitationStatus_Active),
	)

	_, err := s.LoginByInvitationToken(invitationToken)
	mockCtrl.Finish()
	assert.Equal(t, nil, err)
}
