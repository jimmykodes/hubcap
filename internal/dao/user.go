package dao

import (
	"context"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(ctx context.Context, u *dto.User) error
	Get(ctx context.Context, id int64) (*dto.User, error)
	Select(ctx context.Context, sf SearchFilters) ([]*dto.User, error)
	Update(ctx context.Context, u *dto.User) error
	Delete(ctx context.Context, id int64) error
	Close() error
}

const (
	createUser stmt = iota
	getUser
	updateUser
	deleteUser
)

type user struct {
	db    *sqlx.DB
	stmts statements
}

func newUser(db *sqlx.DB) (*user, error) {
	queries := map[stmt]string{
		createUser: "INSERT INTO vehicles.users (email, api_key, super_user) value (?, ?, ?);",
		getUser:    "SELECT id, email, api_key, super_user FROM vehicles.users WHERE id = ?;",
		updateUser: "UPDATE vehicles.users SET email = ?, super_user = ? WHERE id = ?;",
		deleteUser: "DELETE FROM vehicles.users WHERE id = ?",
	}
	s, err := prepareStatements(db, queries)
	if err != nil {
		return nil, err
	}
	return &user{db: db, stmts: s}, nil
}

func (u *user) Create(ctx context.Context, user *dto.User) error {
	_, err := u.stmts[createUser].ExecContext(ctx, user)
	return err
}

func (u *user) Get(ctx context.Context, id int64) (*dto.User, error) {
	user := &dto.User{}
	if err := u.stmts[getUser].GetContext(ctx, user, id); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) Select(ctx context.Context, sf SearchFilters) ([]*dto.User, error) {
	panic("implement me")
}

func (u *user) Update(ctx context.Context, user *dto.User) error {
	_, err := u.stmts[updateUser].ExecContext(ctx, user)
	return err
}

func (u *user) Delete(ctx context.Context, id int64) error {
	_, err := u.stmts[deleteUser].ExecContext(ctx, id)
	return err
}

func (u *user) Close() error {
	return u.stmts.Close()
}
