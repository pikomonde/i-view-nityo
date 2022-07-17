package inmem

import (
	"context"

	"github.com/pikomonde/i-view-nityo/model"
)

func NewRepositoryInMemInvitation(ctx context.Context, config model.Config) *RepositoryInMemInvitation {
	return &RepositoryInMemInvitation{
		context: ctx,
	}
}

type RepositoryInMemInvitation struct {
	context context.Context
	data    map[string]model.Invitation
}

func (r *RepositoryInMemInvitation) CreateInvitation() (invitation model.Invitation, err error) {
	return
}
