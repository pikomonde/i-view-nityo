package model

import "errors"

var (
	Err_Repository_User_NotFound              = errors.New("User Not Found")
	Err_Repository_Invitation_InvalidToken    = errors.New("Invalid Invitation Token")
	Err_Service_Login_WrongPassword           = errors.New("Invalid Username / Password")
	Err_Service_Login_ExpiredInvitationToken  = errors.New("Expired Invitation Token")
	Err_Service_Login_DisabledInvitationToken = errors.New("Invitation Token Has Been Disabled")
)
