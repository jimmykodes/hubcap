package dao

import (
	"context"

	"github.com/google/uuid"
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(ctx context.Context, u *dto.User) error
	Get(ctx context.Context, id int64) (*dto.User, error)
	GetFromApiKey(ctx context.Context, apiKey string) (*dto.User, error)
	GetFromUsername(ctx context.Context, username string) (*dto.User, error)
	Select(ctx context.Context, sf SearchFilters) ([]*dto.User, error)
	Update(ctx context.Context, u *dto.User, id int64) error
	UpdateAPIKey(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
	Close() error
}

const (
	createUser stmt = iota
	getUser
	getUserFromApiKey
	getUserFromUsername
	updateUser
	updateApiKey
	deleteUser
)

type user struct {
	db    *sqlx.DB
	stmts statements
}

func newUser(db *sqlx.DB) (*user, error) {
	queries := map[stmt]string{
		createUser:          "INSERT INTO vehicles.users (username, api_key, super_user) value (?, ?, ?);",
		getUser:             "SELECT id, username, api_key, super_user FROM vehicles.users WHERE id = ?;",
		getUserFromApiKey:   "SELECT id, username, api_key, super_user FROM vehicles.users WHERE api_key = ?;",
		getUserFromUsername: "SELECT id, username, api_key, super_user FROM vehicles.users WHERE username = ?;",
		updateUser:          "UPDATE vehicles.users SET username = ?, super_user = ? WHERE id = ?;",
		updateApiKey:        "UPDATE vehicles.users SET api_key = ? WHERE id = ?;",
		deleteUser:          "DELETE FROM vehicles.users WHERE id = ?",
	}
	s, err := prepareStatements(db, queries)
	if err != nil {
		return nil, err
	}
	return &user{db: db, stmts: s}, nil
}

func (u *user) Create(ctx context.Context, user *dto.User) error {
	apiKey, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	_, err = u.stmts[createUser].ExecContext(ctx, user.Username, apiKey, user.SuperUser)
	return err
}

func (u *user) Get(ctx context.Context, id int64) (*dto.User, error) {
	user := &dto.User{}
	if err := u.stmts[getUser].GetContext(ctx, user, id); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) GetFromApiKey(ctx context.Context, apiKey string) (*dto.User, error) {
	user := &dto.User{}
	if err := u.stmts[getUserFromApiKey].GetContext(ctx, user, apiKey); err != nil {
		return nil, err
	}
	return user, nil
}
func (u *user) GetFromUsername(ctx context.Context, username string) (*dto.User, error) {
	user := &dto.User{}
	if err := u.stmts[getUserFromUsername].GetContext(ctx, user, username); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) Select(ctx context.Context, sf SearchFilters) ([]*dto.User, error) {
	panic("implement me")
}

func (u *user) Update(ctx context.Context, user *dto.User, id int64) error {
	_, err := u.stmts[updateUser].ExecContext(ctx, user.Username, user.SuperUser, id)
	return err
}
func (u *user) UpdateAPIKey(ctx context.Context, id int64) error {
	apiKey, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	_, err = u.stmts[updateApiKey].ExecContext(ctx, apiKey, id)
	return err
}

func (u *user) Delete(ctx context.Context, id int64) error {
	_, err := u.stmts[deleteUser].ExecContext(ctx, id)
	return err
}

func (u *user) Close() error {
	return u.stmts.Close()
}
