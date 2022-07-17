package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pikomonde/i-view-nityo/model"
)

func NewRepositoryInMemInvitation(ctx context.Context, config model.Config, cli *sqlx.DB) *RepositoryInMemInvitation {
	return &RepositoryInMemInvitation{
		context: ctx,
		cli:     cli,
	}
}

type RepositoryInMemInvitation struct {
	context context.Context
	cli     *sqlx.DB
}

func (r *RepositoryInMemInvitation) CreateInvitation(invitation model.Invitation) (model.Invitation, error) {
	query := `insert into invitation (token, status, created_at) values (?, ?, ?)`

	res, err := r.cli.Exec(query, invitation.Token, invitation.Status, invitation.CreatedAt)
	if err != nil {
		return model.Invitation{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return model.Invitation{}, err
	}

	return model.Invitation{
		ID:        lastID,
		Token:     invitation.Token,
		Status:    invitation.Status,
		CreatedAt: invitation.CreatedAt,
	}, nil
}

func (r *RepositoryInMemInvitation) IsInvitationExist(invitationToken string) (bool, error) {
	query := `select id from invitation where token = ?`

	var id int64
	err := r.cli.QueryRow(query, invitationToken).
		Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *RepositoryInMemInvitation) GetInvitations() ([]model.Invitation, error) {
	query := `select id, token, status, created_at from invitation order by id`

	invitations := make([]model.Invitation, 0)
	rows, err := r.cli.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		invitation := model.Invitation{}
		err = rows.Scan(&invitation.ID, &invitation.Token, &invitation.Status, &invitation.CreatedAt)
		if err != nil {
			return nil, err
		}
		invitations = append(invitations, invitation)
	}
	return invitations, nil
}

func (r *RepositoryInMemInvitation) GetInvitationByToken(invitationToken string) (model.Invitation, error) {
	query := `select id, token, status, created_at from invitation where token = ?`

	invitation := model.Invitation{}
	err := r.cli.QueryRow(query, invitationToken).
		Scan(&invitation.ID, &invitation.Token, &invitation.Status, &invitation.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Invitation{}, model.Err_Repository_Invitation_InvalidToken
		}
		return model.Invitation{}, err
	}
	return invitation, nil
}

func (r *RepositoryInMemInvitation) UpdateInvitationStatus(invitationToken string, updatedStatus model.InvitationStatus) error {
	query := `update invitation set status = ? where token = ?`

	res, err := r.cli.Exec(query, updatedStatus, invitationToken)
	if err != nil {
		return err
	}

	nUpdated, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if nUpdated == 0 {
		return model.Err_Repository_Invitation_InvalidToken
	}

	return nil
}
