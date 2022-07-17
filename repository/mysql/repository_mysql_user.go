package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pikomonde/i-view-nityo/model"
)

func NewRepositoryInMemUser(ctx context.Context, config model.Config, cli *sqlx.DB) *RepositoryInMemUser {
	return &RepositoryInMemUser{
		context: ctx,
		cli:     cli,
	}
}

type RepositoryInMemUser struct {
	context context.Context
	cli     *sqlx.DB
}

func (r *RepositoryInMemUser) CreateUser(user model.User) (model.User, error) {
	query := `insert into user (invitation_token, username, password, role, created_at) values (?, ?, ?, ?, ?)`

	res, err := r.cli.Exec(query, "", user.Username, user.Password, user.Role, user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:        lastID,
		Username:  user.Username,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *RepositoryInMemUser) GetUserByUsername(username string) (model.User, error) {
	query := `select id, invitation_token, username, password, role, created_at from user where username = ?`

	user := model.User{}
	err := r.cli.QueryRow(query, username).
		Scan(&user.ID, &user.InvitationToken, &user.Username, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, model.Err_Repository_User_NotFound
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *RepositoryInMemUser) GetUserByID(id int64) (model.User, error) {
	query := `select id, invitation_token, username, password, role, created_at from user where id = ?`

	user := model.User{}
	err := r.cli.QueryRow(query, id).
		Scan(&user.ID, &user.InvitationToken, &user.Username, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, model.Err_Repository_User_NotFound
		}
		return model.User{}, err
	}
	return user, nil
}
