package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pikomonde/i-view-nityo/helper"
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

func (s *ServiceLogin) RegisterAdminIfNotExist() error {
	if _, err := s.repositoryUser.GetUserByID(0); err != model.Err_Repository_User_NotFound {
		// already exist
		return nil
	}

	password := helper.RandomString(24)
	fmt.Printf(`
======================================
              ADMIN INFO
======================================
  username: admin
  password: %s
======================================

`, password)

	user := model.User{
		Username:  "admin",
		Password:  password,
		Role:      model.UserRole_Admin,
		CreatedAt: time.Now().UnixNano(),
	}
	hashedPassword := helper.HashUserPassword(user)
	user.Password = hashedPassword

	_, err := s.repositoryUser.CreateUser(user)

	return err
}

func (s *ServiceLogin) LoginByUsernamePassword(username, password string) (string, error) {
	user, err := s.repositoryUser.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	// comparing password
	hashedStoredPassword := user.Password

	userWithUnhashedInputPassword := user
	userWithUnhashedInputPassword.Password = password
	hashedInputPassword := helper.HashUserPassword(userWithUnhashedInputPassword)

	if hashedStoredPassword != hashedInputPassword {
		return "", model.Err_Service_Login_WrongPassword
	}

	// hashing user data for jwt
	userStr, err := json.Marshal(model.User{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	if err != nil {
		return "", err
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": string(userStr),
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})
	token, err := sign.SignedString([]byte(s.config.App.JWTSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *ServiceLogin) LoginByInvitationToken(invitationToken string) error {
	return nil
}
