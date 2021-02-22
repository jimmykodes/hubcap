package dao

import (
	"context"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type Vehicle interface {
	Create(ctx context.Context, v *dto.Vehicle) error
	Get(ctx context.Context, id int64) (*dto.Vehicle, error)
	Select(ctx context.Context, sf SearchFilters) ([]*dto.Vehicle, error)
	Update(ctx context.Context, v *dto.Vehicle) error
	Delete(ctx context.Context, id int64) error
	Close() error
}

type vehicle struct {
	db    *sqlx.DB
	stmts statements
}

func newVehicle(db *sqlx.DB) (*vehicle, error) {
	return &vehicle{db: db}, nil
}

func (v *vehicle) Create(ctx context.Context, vehicle *dto.Vehicle) error {
	panic("implement me")
}

func (v *vehicle) Get(ctx context.Context, id int64) (*dto.Vehicle, error) {
	panic("implement me")
}

func (v *vehicle) Select(ctx context.Context, sf SearchFilters) ([]*dto.Vehicle, error) {
	panic("implement me")
}

func (v *vehicle) Update(ctx context.Context, vehicle *dto.Vehicle) error {
	panic("implement me")
}

func (v *vehicle) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}

func (v *vehicle) Close() error {
	return v.stmts.Close()
}
