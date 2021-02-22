package dao

import (
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(u *dto.User) error
	Get(id int64) (*dto.User, error)
	Select(sf SearchFilters) ([]*dto.User, error)
	Update(u *dto.User) error
	Delete(id int64) error
	Close() error
}

type user struct {
	db    *sqlx.DB
	stmts statements
}

func newUser(db *sqlx.DB) (*user, error) {
	return &user{db: db}, nil
}

func (u *user) Create(user *dto.User) error {
	panic("implement me")
}

func (u *user) Get(id int64) (*dto.User, error) {
	panic("implement me")
}

func (u *user) Select(sf SearchFilters) ([]*dto.User, error) {
	panic("implement me")
}

func (u *user) Update(user *dto.User) error {
	panic("implement me")
}

func (u *user) Delete(id int64) error {
	panic("implement me")
}

func (u *user) Close() error {
	return u.stmts.Close()
}
