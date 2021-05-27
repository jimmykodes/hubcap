package dao

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

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

type serviceDAO struct {
	conn *pgxpool.Pool

	createQuery string
	getQuery    string
	searchQuery string
	updateQuery string
	deleteQuery string

	filterFields fields
	searchFields fields
}

func newService(conn *pgxpool.Pool) (*serviceDAO, error) {
	var getQuery = `
SELECT s.id, s.date, s.odometer, s.data, s.user_id, s.vehicle_id, s.service_type_id, st.name as service_type_name, v.name as vehicle_name FROM services s
	JOIN service_types st on st.id = s.service_type_id
	JOIN vehicles v on v.id = s.vehicle_id
WHERE s.id = $1 AND s.user_id = $2;
`
	var searchQuery = `
SELECT s.id, s.date, s.odometer, s.data, s.user_id, s.vehicle_id, s.service_type_id, st.name as service_type_name, v.name as vehicle_name FROM services s
	JOIN service_types st on st.id = s.service_type_id
	JOIN vehicles v on v.id = s.vehicle_id
WHERE s.user_id = $1
`
	// todo: figure out odometer/date gt lt logic and data logic
	// 	might be JSON_EXTRACT_SCALAR or something?
	ff := fields{"vehicle_id": true, "service_type_id": true}
	sf := fields{}
	return &serviceDAO{
		conn: conn,

		createQuery: "INSERT INTO services (date, odometer, data, user_id, vehicle_id, service_type_id) VALUES ($1, $2, $3, $4, $5, $6);",
		getQuery:    getQuery,
		searchQuery: searchQuery,
		updateQuery: "UPDATE services SET date = $1, odometer = $2, data = $3, vehicle_id = $4, service_type_id = $5 WHERE id = $6 AND user_id = $7;",
		deleteQuery: "DELETE FROM services WHERE id = $1 AND user_id = $2;",

		filterFields: ff,
		searchFields: sf,
	}, nil
}

func (s *serviceDAO) Create(ctx context.Context, service *dto.Service) error {
	_, err := s.conn.Exec(ctx, s.createQuery, service.Date, service.Odometer, service.Data, service.UserID, service.VehicleID, service.ServiceTypeID)
	return err
}

func (s *serviceDAO) Get(ctx context.Context, id, userID int64) (*dto.Service, error) {
	service := &dto.Service{}
	row := s.conn.QueryRow(ctx, s.getQuery, id, userID)
	if err := row.Scan(&service.ID, &service.Date, &service.Odometer, &service.Data, &service.UserID, &service.VehicleID, &service.ServiceTypeID, &service.ServiceTypeName, &service.VehicleName); err != nil {
		return nil, err
	}
	return service, nil

}

func (s *serviceDAO) Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.Service, error) {
	wc := sf.whereClause(s.searchFields, s.filterFields, 2)
	query := s.searchQuery
	args := []interface{}{userID}
	if q := wc.query(); q != "" {
		query = fmt.Sprintf("%s AND %s", s.searchQuery, wc.query())
		args = append(args, wc.args...)
	}
	var services []*dto.Service
	rows, err := s.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		service := &dto.Service{}
		if err := rows.Scan(&service.ID, &service.Date, &service.Odometer, &service.Data, &service.UserID, &service.VehicleID, &service.ServiceTypeID, &service.ServiceTypeName, &service.VehicleName); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil

}

func (s *serviceDAO) Update(ctx context.Context, service *dto.Service, id, userID int64) error {
	_, err := s.conn.Exec(ctx, s.updateQuery, service.Date, service.Odometer, service.Data, service.VehicleID, service.ServiceTypeID, id, userID)
	return err
}

func (s *serviceDAO) Delete(ctx context.Context, id, userID int64) error {
	_, err := s.conn.Exec(ctx, s.deleteQuery, id, userID)
	return err
}

func (s *serviceDAO) Close() error {
	return nil
}
