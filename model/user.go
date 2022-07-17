package model

type UserRole string

const (
	UserRole_Admin      UserRole = "admin"
	UserRole_Invitation UserRole = "invitation"
)

type User struct {
	ID           int64
	InvitationID int64
	Username     string
	Password     string
	Role         UserRole
	CreatedAt    int64
}
