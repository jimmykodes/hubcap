package dao

import (
	"context"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type Service interface {
	Create(ctx context.Context, s *dto.Service) error
	Get(ctx context.Context, id int64) (*dto.Service, error)
	Select(ctx context.Context, sf SearchFilters) ([]*dto.Service, error)
	Update(ctx context.Context, s *dto.Service) error
	Delete(ctx context.Context, id int64) error
	Close() error
}

type service struct {
	db    *sqlx.DB
	stmts statements
}

func newService(db *sqlx.DB) (*service, error) {
	return &service{db: db}, nil
}

func (s *service) Create(ctx context.Context, service *dto.Service) error {
	panic("implement me")
}

func (s *service) Get(ctx context.Context, id int64) (*dto.Service, error) {
	panic("implement me")
}

func (s *service) Select(ctx context.Context, sf SearchFilters) ([]*dto.Service, error) {
	panic("implement me")
}

func (s *service) Update(ctx context.Context, service *dto.Service) error {
	panic("implement me")
}

func (s *service) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}

func (s *service) Close() error {
	return s.stmts.Close()
}
