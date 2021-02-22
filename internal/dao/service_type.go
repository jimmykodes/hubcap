package dao

import (
	"context"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type ServiceType interface {
	Create(ctx context.Context, st *dto.ServiceType) error
	Get(ctx context.Context, id int64) (*dto.ServiceType, error)
	Select(ctx context.Context, sf SearchFilters) ([]*dto.ServiceType, error)
	Update(ctx context.Context, st *dto.ServiceType) error
	Delete(ctx context.Context, id int64) error
	Close() error
}

type serviceType struct {
	db    *sqlx.DB
	stmts statements
}

func newServiceType(db *sqlx.DB) (*serviceType, error) {
	return &serviceType{db: db}, nil
}

func (st *serviceType) Create(ctx context.Context, serviceType *dto.ServiceType) error {
	panic("implement me")
}

func (st *serviceType) Get(ctx context.Context, id int64) (*dto.ServiceType, error) {
	panic("implement me")
}

func (st *serviceType) Select(ctx context.Context, sf SearchFilters) ([]*dto.ServiceType, error) {
	panic("implement me")
}

func (st *serviceType) Update(ctx context.Context, serviceType *dto.ServiceType) error {
	panic("implement me")
}

func (st *serviceType) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}

func (st *serviceType) Close() error {
	return st.stmts.Close()
}
