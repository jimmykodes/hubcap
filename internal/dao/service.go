package dao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
)

type Service interface {
	Create(ctx context.Context, s *dto.Service) error
	Get(ctx context.Context, id, userID int64) (*dto.Service, error)
	Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.Service, error)
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
	db           *sqlx.DB
	stmts        statements
	filterFields fields
	searchFields fields
	searchQuery  string
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
	// todo: figure out odometer/date gt lt logic and data logic
	// 	might be JSON_EXTRACT_SCALAR or something?
	ff := fields{"vehicle_id": true, "service_type_id": true}
	sf := fields{}
	return &service{
		db:           db,
		stmts:        s,
		filterFields: ff,
		searchFields: sf,
		searchQuery:  "SELECT id, date, odometer, data, user_id, vehicle_id, service_type_id FROM vehicles.services WHERE user_id = ?",
	}, nil
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

func (s *service) Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.Service, error) {
	wc := sf.whereClause(s.searchFields, s.filterFields)
	query := s.searchQuery
	args := []interface{}{userID}
	if q := wc.query(); q != "" {
		query = fmt.Sprintf("%s AND %s", s.searchQuery, wc.query())
		args = append(args, wc.args...)
	}
	var rows []*dto.Service
	if err := s.db.SelectContext(ctx, &rows, query, args...); err != nil {
		return nil, err
	}
	return rows, nil

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
