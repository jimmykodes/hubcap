package dao

import (
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type Service interface {
	Create(s *dto.Service) error
	Get(id int64) (*dto.Service, error)
	Select(sf SearchFilters) ([]*dto.Service, error)
	Update(s *dto.Service) error
	Delete(id int64) error
	Close() error
}

type service struct {
	db    *sqlx.DB
	stmts statements
}

func newService(db *sqlx.DB) (*service, error) {
	return &service{db: db}, nil
}

func (s *service) Create(service *dto.Service) error {
	panic("implement me")
}

func (s *service) Get(id int64) (*dto.Service, error) {
	panic("implement me")
}

func (s *service) Select(sf SearchFilters) ([]*dto.Service, error) {
	panic("implement me")
}

func (s *service) Update(service *dto.Service) error {
	panic("implement me")
}

func (s *service) Delete(id int64) error {
	panic("implement me")
}

func (s *service) Close() error {
	return s.stmts.Close()
}
