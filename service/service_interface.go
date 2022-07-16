package service

type Login interface {
	LoginByUsernamePassword(username, password string) error
	LoginByInvitationID(invitationID string) error
}

type Invitation interface {
	CreateInvitation(userID int) error
}
