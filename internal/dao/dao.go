package dao

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	"github.com/jimmykodes/vehicle_maintenance/internal/settings"
)

type DAO struct {
	conn *pgxpool.Pool

	Service     Service
	ServiceType ServiceType
	User        User
	Vehicle     Vehicle
}

func New(dbSettings settings.DB, logger *zap.Logger) (*DAO, error) {
	conn, err := pgxpool.Connect(context.Background(), dbSettings.DSN())
	if err != nil {
		return nil, err
	}
	vehicle, err := newVehicle(conn)
	if err != nil {
		return nil, fmt.Errorf("error creating vehicleDAO dao: %w", err)
	}
	service, err := newService(conn, logger)
	if err != nil {
		return nil, fmt.Errorf("error creating serviceDAO dao: %w", err)
	}
	serviceType, err := newServiceType(conn)
	if err != nil {
		return nil, fmt.Errorf("error creating serviceDAO type dao: %w", err)
	}
	user, err := newUserDAO(conn)
	if err != nil {
		return nil, fmt.Errorf("error creating userDAO dao: %w", err)
	}
	return &DAO{
		conn:        conn,
		Vehicle:     vehicle,
		Service:     service,
		ServiceType: serviceType,
		User:        user,
	}, nil
}

func (d DAO) Close() {
	d.conn.Close()
}
