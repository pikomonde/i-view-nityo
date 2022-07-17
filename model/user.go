package model

type UserRole string

const (
	UserRole_Admin      UserRole = "admin"
	UserRole_Invitation UserRole = "invitation"
)

type User struct {
	ID              int64    `json:"id"`
	InvitationToken string   `json:"invitation_token"`
	Username        string   `json:"username"`
	Password        string   `json:"password"`
	Role            UserRole `json:"role"`
	CreatedAt       int64    `json:"created_at"`
}
