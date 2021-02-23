package dao

import (
	"context"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type Service interface {
	Create(ctx context.Context, s *dto.Service) error
	Get(ctx context.Context, id, userID int64) (*dto.Service, error)
	Select(ctx context.Context, sf SearchFilters) ([]*dto.Service, error)
	Update(ctx context.Context, s *dto.Service, id, userID int64) error
	Delete(ctx context.Context, id, userID int64) error
	Close() error
}

const (
	createService stmt = iota
	getService
	updateService
	deleteService
)

type service struct {
	db    *sqlx.DB
	stmts statements
}

func newService(db *sqlx.DB) (*service, error) {
	q := queries{
		createService: "INSERT INTO vehicles.services (date, odometer, data, user_id, vehicle_id, service_type_id) VALUE (?, ?, ?, ?, ?, ?);",
		getService:    "SELECT id, date, odometer, data, user_id, vehicle_id, service_type_id FROM vehicles.services WHERE id = ? AND user_id = ?;",
		updateService: "UPDATE vehicles.services SET date = ?, odometer = ?, data = ?, vehicle_id = ?, service_type_id = ? WHERE id = ? AND user_id = ?;",
		deleteService: "DELETE FROM vehicles.services WHERE id = ? AND user_id = ?;",
	}
	s, err := prepareStatements(db, q)
	if err != nil {
		return nil, err
	}
	return &service{db: db, stmts: s}, nil
}

func (s *service) Create(ctx context.Context, service *dto.Service) error {
	_, err := s.stmts[createService].ExecContext(ctx, service.Date, service.Odometer, service.Data, service.UserID, service.VehicleID, service.ServiceTypeID)
	return err
}

func (s *service) Get(ctx context.Context, id, userID int64) (*dto.Service, error) {
	obj := &dto.Service{}
	if err := s.stmts[getService].GetContext(ctx, obj, id, userID); err != nil {
		return nil, err
	}
	return obj, nil

}

func (s *service) Select(ctx context.Context, sf SearchFilters) ([]*dto.Service, error) {
	panic("implement me")
}

func (s *service) Update(ctx context.Context, service *dto.Service, id, userID int64) error {
	_, err := s.stmts[updateService].ExecContext(ctx, service.Date, service.Odometer, service.Data, service.VehicleID, service.ServiceTypeID, id, userID)
	return err
}

func (s *service) Delete(ctx context.Context, id, userID int64) error {
	_, err := s.stmts[deleteService].ExecContext(ctx, id, userID)
	return err
}

func (s *service) Close() error {
	return s.stmts.Close()
}
