package inmem

import (
	"context"

	"github.com/pikomonde/i-view-nityo/model"
)

func NewRepositoryInMemUser(ctx context.Context, config model.Config) *RepositoryInMemUser {
	return &RepositoryInMemUser{
		context: ctx,
	}
}

type RepositoryInMemUser struct {
	context context.Context
}

func (r *RepositoryInMemUser) GetUserByID(id int64) (user model.User, err error) {
	return
}
