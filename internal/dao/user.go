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

type user struct {
	db    *sqlx.DB
	stmts statements
}

func newUser(db *sqlx.DB) (*user, error) {
	return &user{db: db}, nil
}

func (u *user) Create(ctx context.Context, user *dto.User) error {
	panic("implement me")
}

func (u *user) Get(ctx context.Context, id int64) (*dto.User, error) {
	panic("implement me")
}

func (u *user) Select(ctx context.Context, sf SearchFilters) ([]*dto.User, error) {
	panic("implement me")
}

func (u *user) Update(ctx context.Context, user *dto.User) error {
	panic("implement me")
}

func (u *user) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u *user) Close() error {
	return u.stmts.Close()
}
