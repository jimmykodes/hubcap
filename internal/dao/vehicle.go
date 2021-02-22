package dao

import (
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type Vehicle interface {
	Create(v *dto.Vehicle) error
	Get(id int64) (*dto.Vehicle, error)
	Select(sf SearchFilters) ([]*dto.Vehicle, error)
	Update(v *dto.Vehicle) error
	Delete(id int64) error
	Close() error
}

type vehicle struct {
	db    *sqlx.DB
	stmts statements
}

func newVehicle(db *sqlx.DB) (*vehicle, error) {
	return &vehicle{db: db}, nil
}

func (v *vehicle) Create(vehicle *dto.Vehicle) error {
	panic("implement me")
}

func (v *vehicle) Get(id int64) (*dto.Vehicle, error) {
	panic("implement me")
}

func (v *vehicle) Select(sf SearchFilters) ([]*dto.Vehicle, error) {
	panic("implement me")
}

func (v *vehicle) Update(vehicle *dto.Vehicle) error {
	panic("implement me")
}

func (v *vehicle) Delete(id int64) error {
	panic("implement me")
}

func (v *vehicle) Close() error {
	return v.stmts.Close()
}
