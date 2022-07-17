package model

import "errors"

var (
	Err_Repository_User_NotFound              = errors.New("Err_Repository_User_NotFound")
	Err_Repository_Invitation_InvalidToken    = errors.New("Err_Repository_Invitation_InvalidToken")
	Err_Service_Login_WrongPassword           = errors.New("Err_Service_Login_WrongPassword")
	Err_Service_Login_ExpiredInvitationToken  = errors.New("Err_Service_Login_ExpiredInvitationToken")
	Err_Service_Login_DisabledInvitationToken = errors.New("Err_Service_Login_DisabledInvitationToken")
)
