package model

type InvitationStatus string

const (
	InvitationStatus_Disabled InvitationStatus = "disabled"
	InvitationStatus_Inactive InvitationStatus = "inactive"
	InvitationStatus_Active   InvitationStatus = "active"
)

type Invitation struct {
	ID        int64            `json:"id"`
	Token     string           `json:"token"`
	Status    InvitationStatus `json:"status"`
	CreatedAt int64            `json:"created_at"`
}
