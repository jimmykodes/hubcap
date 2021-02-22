package dao

import (
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type ServiceType interface {
	Create(st *dto.ServiceType) error
	Get(id int64) (*dto.ServiceType, error)
	Select(sf SearchFilters) ([]*dto.ServiceType, error)
	Update(st *dto.ServiceType) error
	Delete(id int64) error
	Close() error
}

type serviceType struct {
	db    *sqlx.DB
	stmts statements
}

func newServiceType(db *sqlx.DB) (*serviceType, error) {
	return &serviceType{db: db}, nil
}

func (st *serviceType) Create(serviceType *dto.ServiceType) error {
	panic("implement me")
}

func (st *serviceType) Get(id int64) (*dto.ServiceType, error) {
	panic("implement me")
}

func (st *serviceType) Select(sf SearchFilters) ([]*dto.ServiceType, error) {
	panic("implement me")
}

func (st *serviceType) Update(serviceType *dto.ServiceType) error {
	panic("implement me")
}

func (st *serviceType) Delete(id int64) error {
	panic("implement me")
}

func (st *serviceType) Close() error {
	return st.stmts.Close()
}
