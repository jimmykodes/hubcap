package dao

import (
	"github.com/jimmykodes/vehicle_maintenance/internal/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/multierr"
)

type DAO struct {
	Service     Service
	ServiceType ServiceType
	User        User
	Vehicle     Vehicle
}

func New(dbSettings settings.DB) (*DAO, error) {
	db, err := sqlx.Open(dbSettings.DriveName, dbSettings.DNS())
	if err != nil {
		return nil, err
	}
	vehicle, err := newVehicle(db)
	if err != nil {
		return nil, err
	}
	service, err := newService(db)
	if err != nil {
		return nil, err
	}
	serviceType, err := newServiceType(db)
	if err != nil {
		return nil, err
	}
	user, err := newUser(db)
	if err != nil {
		return nil, err
	}
	return &DAO{
		Vehicle:     vehicle,
		Service:     service,
		ServiceType: serviceType,
		User:        user,
	}, nil
}

func (d DAO) Close() error {
	return multierr.Combine(
		d.Service.Close(),
		d.ServiceType.Close(),
		d.User.Close(),
		d.Vehicle.Close(),
	)
}
