package dao

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
)

type User interface {
	Create(ctx context.Context, u *dto.User) error
	CreateSession(ctx context.Context, user *dto.User, expires time.Time) (string, error)
	Get(ctx context.Context, id int64) (*dto.User, error)
	GetFromApiKey(ctx context.Context, apiKey string) (*dto.User, error)
	GetFromSession(ctx context.Context, session string, time int64) (*dto.User, error)
	GetFromUsername(ctx context.Context, username string) (*dto.User, error)
	Update(ctx context.Context, u *dto.User, id int64) error
	UpdateAPIKey(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
}

type userDAO struct {
	conn *pgxpool.Pool

	createUserQuery          string
	createSessionQuery       string
	getUserQuery             string
	getUserFromApiKeyQuery   string
	getUserFromSessionQuery  string
	getUserFromUsernameQuery string
	updateUserQuery          string
	updateApiKeyQuery        string
	deleteUserQuery          string
}

func newUserDAO(conn *pgxpool.Pool) (*userDAO, error) {
	return &userDAO{
		conn: conn,

		createUserQuery:          "INSERT INTO users (username, api_key, super_user) values ($1, $2, $3);",
		createSessionQuery:       "INSERT INTO sessions (key, user_id, expires) values ($1, $2, $3)",
		getUserQuery:             "SELECT id, username, api_key, super_user FROM users WHERE id = $1;",
		getUserFromApiKeyQuery:   "SELECT id, username, api_key, super_user FROM users WHERE api_key = $1;",
		getUserFromSessionQuery:  "SELECT u.id, u.username, u.api_key, u.super_user FROM users u JOIN sessions s on u.id = s.user_id WHERE s.key = $1 and s.expires > $2;",
		getUserFromUsernameQuery: "SELECT id, username, api_key, super_user FROM users WHERE username = $1;",
		updateUserQuery:          "UPDATE users SET username = $1, super_user = $2 WHERE id = $3;",
		updateApiKeyQuery:        "UPDATE users SET api_key = $1 WHERE id = $2;",
		deleteUserQuery:          "DELETE FROM users WHERE id = $1",
	}, nil
}

func (u *userDAO) Create(ctx context.Context, user *dto.User) error {
	apiKey, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	_, err = u.conn.Exec(ctx, u.createUserQuery, user.Username, apiKey, user.SuperUser)
	return err
}

func (u *userDAO) CreateSession(ctx context.Context, user *dto.User, expires time.Time) (string, error) {
	sessionKey, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	_, err = u.conn.Exec(ctx, u.createSessionQuery, sessionKey, user.ID, expires.Unix())
	if err != nil {
		return "", err
	}
	return sessionKey.String(), nil
}

func (u *userDAO) getUser(ctx context.Context, q string, arg ...interface{}) (*dto.User, error) {
	user := &dto.User{}
	row := u.conn.QueryRow(ctx, q, arg...)
	if err := row.Scan(&user.ID, &user.Username, &user.ApiKey, &user.SuperUser); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userDAO) Get(ctx context.Context, id int64) (*dto.User, error) {
	return u.getUser(ctx, u.getUserQuery, id)
}

func (u *userDAO) GetFromApiKey(ctx context.Context, apiKey string) (*dto.User, error) {
	return u.getUser(ctx, u.getUserFromApiKeyQuery, apiKey)
}

func (u *userDAO) GetFromUsername(ctx context.Context, username string) (*dto.User, error) {
	return u.getUser(ctx, u.getUserFromUsernameQuery, username)
}

func (u *userDAO) GetFromSession(ctx context.Context, session string, time int64) (*dto.User, error) {
	return u.getUser(ctx, u.getUserFromSessionQuery, session, time)
}

func (u *userDAO) Update(ctx context.Context, user *dto.User, id int64) error {
	_, err := u.conn.Exec(ctx, u.updateUserQuery, user.Username, user.SuperUser, id)
	return err
}
func (u *userDAO) UpdateAPIKey(ctx context.Context, id int64) error {
	apiKey, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	_, err = u.conn.Exec(ctx, u.updateApiKeyQuery, apiKey, id)
	return err
}

func (u *userDAO) Delete(ctx context.Context, id int64) error {
	_, err := u.conn.Exec(ctx, u.deleteUserQuery, id)
	return err
}
