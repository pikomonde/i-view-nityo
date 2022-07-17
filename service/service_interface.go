package service

type Login interface {
	RegisterAdminIfNotExist() error
	LoginByUsernamePassword(username, password string) (string, error)
	LoginByInvitationToken(invitationToken string) error
}

type Invitation interface {
	CreateInvitation(userID int) error
	GetInvitations() ([]Invitation, error)
}
