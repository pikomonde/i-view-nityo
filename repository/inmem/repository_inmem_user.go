package inmem

import (
	"context"

	"github.com/pikomonde/i-view-nityo/model"
)

func NewRepositoryInMemUser(ctx context.Context, config model.Config) *RepositoryInMemUser {
	return &RepositoryInMemUser{
		context: ctx,
		data:    make(map[int64]model.User),
	}
}

type RepositoryInMemUser struct {
	context context.Context
	data    map[int64]model.User
}

func (r *RepositoryInMemUser) CreateUser(user model.User) (model.User, error) {
	newID := int64(len(r.data))
	r.data[newID] = model.User{
		ID:           newID,
		InvitationID: user.InvitationID,
		Username:     user.Username,
		Password:     user.Password,
		Role:         user.Role,
		CreatedAt:    user.CreatedAt,
	}
	return r.data[newID], nil
}

func (r *RepositoryInMemUser) GetUserByUsername(username string) (model.User, error) {
	for _, user := range r.data {
		if user.Username == username {
			return user, nil
		}
	}
	return model.User{}, model.Err_Repository_User_NotFound
}

func (r *RepositoryInMemUser) GetUserByID(id int64) (model.User, error) {
	user, isExist := r.data[id]
	if !isExist {
		return model.User{}, model.Err_Repository_User_NotFound
	}
	return user, nil
}
