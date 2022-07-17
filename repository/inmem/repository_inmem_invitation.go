package inmem

import (
	"context"
	"sort"

	"github.com/pikomonde/i-view-nityo/model"
)

func NewRepositoryInMemInvitation(ctx context.Context, config model.Config) *RepositoryInMemInvitation {
	return &RepositoryInMemInvitation{
		context: ctx,
		data:    make(map[int64]model.Invitation),
	}
}

type RepositoryInMemInvitation struct {
	context context.Context
	data    map[int64]model.Invitation
	lastID  int64
}

func (r *RepositoryInMemInvitation) CreateInvitation(invitation model.Invitation) (model.Invitation, error) {
	r.lastID++
	newID := int64(r.lastID)
	r.data[newID] = model.Invitation{
		ID:        newID,
		Token:     invitation.Token,
		Status:    invitation.Status,
		CreatedAt: invitation.CreatedAt,
	}
	return r.data[newID], nil
}

func (r *RepositoryInMemInvitation) IsInvitationExist(invitationToken string) (bool, error) {
	for _, invitation := range r.data {
		if invitation.Token == invitationToken {
			return true, nil
		}
	}
	return false, nil
}

func (r *RepositoryInMemInvitation) GetInvitations() ([]model.Invitation, error) {
	invitations := make([]model.Invitation, 0)
	for _, invitation := range r.data {
		invitations = append(invitations, invitation)
	}
	sort.Slice(invitations, func(i, j int) bool {
		return invitations[i].ID < invitations[j].ID
	})
	return invitations, nil
}

func (r *RepositoryInMemInvitation) GetInvitationByToken(invitationToken string) (model.Invitation, error) {
	for _, invitation := range r.data {
		if invitation.Token == invitationToken {
			return invitation, nil
		}
	}
	return model.Invitation{}, model.Err_Repository_Invitation_InvalidToken
}

func (r *RepositoryInMemInvitation) UpdateInvitationStatus(invitationToken string, updatedStatus model.InvitationStatus) error {
	for _, invitation := range r.data {
		if invitation.Token == invitationToken {
			invitation.Status = updatedStatus
			r.data[invitation.ID] = invitation
			return nil
		}
	}
	return model.Err_Repository_Invitation_InvalidToken
}
