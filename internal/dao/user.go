package dao

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
)

type User interface {
	Create(ctx context.Context, u *dto.User) (*dto.User, error)
	CreateSession(ctx context.Context, user *dto.User, expires time.Time) (string, error)
	Get(ctx context.Context, id int64) (*dto.User, error)
	GetFromApiKey(ctx context.Context, apiKey string) (*dto.User, error)
	GetFromSession(ctx context.Context, session string, time int64) (*dto.User, error)
	GetFromUsername(ctx context.Context, username string) (*dto.User, error)
	Select(ctx context.Context, sf SearchFilters) ([]*dto.User, error)
	Update(ctx context.Context, u *dto.User, id int64) error
	UpdateAPIKey(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
	Close() error
}

const (
	createUser stmt = iota
	createSession
	getUser
	getUserFromApiKey
	getUserFromSession
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
		createSession:       "INSERT INTO vehicles.sessions (`key`, user_id, expires) value (?, ?, ?)",
		getUser:             "SELECT id, username, api_key, super_user FROM vehicles.users WHERE id = ?;",
		getUserFromApiKey:   "SELECT id, username, api_key, super_user FROM vehicles.users WHERE api_key = ?;",
		getUserFromSession:  "SELECT u.id, u.username, u.api_key, u.super_user FROM vehicles.users u JOIN vehicles.sessions s on u.id = s.user_id WHERE s.`key` = ? and s.expires > ?;",
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

func (u *user) Create(ctx context.Context, user *dto.User) (*dto.User, error) {
	apiKey, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	res, err := u.stmts[createUser].ExecContext(ctx, user.Username, apiKey, user.SuperUser)
	if err != nil {
		return nil, err
	}
	user.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ApiKey = apiKey.String()
	return user, nil
}

func (u *user) CreateSession(ctx context.Context, user *dto.User, expires time.Time) (string, error) {
	sessionKey, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	_, err = u.stmts[createSession].ExecContext(ctx, sessionKey, user.ID, expires.Unix())
	if err != nil {
		return "", err
	}
	return sessionKey.String(), nil
}

func (u *user) Get(ctx context.Context, id int64) (*dto.User, error) {
	user := &dto.User{}
	if err := u.stmts[getUser].GetContext(ctx, user, id); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) getUser(ctx context.Context, s stmt, arg ...interface{}) (*dto.User, error) {
	user := &dto.User{}
	if err := u.stmts[s].GetContext(ctx, user, arg...); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) GetFromApiKey(ctx context.Context, apiKey string) (*dto.User, error) {
	return u.getUser(ctx, getUserFromApiKey, apiKey)
}
func (u *user) GetFromUsername(ctx context.Context, username string) (*dto.User, error) {
	return u.getUser(ctx, getUserFromUsername, username)
}
func (u *user) GetFromSession(ctx context.Context, session string, time int64) (*dto.User, error) {
	return u.getUser(ctx, getUserFromSession, session, time)
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
