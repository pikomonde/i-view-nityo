package model

type InvitationStatus string

const (
	InvitationStatus_Disabled UserRole = "disabled"
	InvitationStatus_Inactive UserRole = "inactive"
	InvitationStatus_Active   UserRole = "active"
)

type Invitation struct {
	ID        int64
	Token     string
	Status    InvitationStatus
	CreatedBy int64
	CreatedAt int64
}
