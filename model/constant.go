package model

import "errors"

var (
	Err_Repository_User_NotFound    = errors.New("Err_Repository_User_NotFound")
	Err_Service_Login_WrongPassword = errors.New("Err_Service_Login_WrongPassword")
)
