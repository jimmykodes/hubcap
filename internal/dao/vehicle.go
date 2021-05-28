package dao

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
)

type Vehicle interface {
	Create(ctx context.Context, v *dto.Vehicle) error
	Get(ctx context.Context, id, userID int64) (*dto.Vehicle, error)
	Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.Vehicle, error)
	Update(ctx context.Context, v *dto.Vehicle, id, userID int64) error
	Delete(ctx context.Context, id, userID int64) error
}

type vehicleDAO struct {
	conn *pgxpool.Pool

	createQuery string
	getQuery    string
	searchQuery string
	updateQuery string

	deleteQuery  string
	filterFields fields
	searchFields fields
}

func newVehicle(conn *pgxpool.Pool) (*vehicleDAO, error) {
	ff := fields{"make": true, "model": true, "year": true}
	sf := fields{"name": true}
	return &vehicleDAO{
		conn: conn,

		createQuery: "INSERT INTO vehicles (name, make, model, year, user_id) VALUES ($1, $2, $3, $4, $5);",
		getQuery:    "SELECT id, name, make, model, year, user_id FROM vehicles WHERE id = $1 AND user_id = $2;",
		searchQuery: "SELECT id, name, make, model, year, user_id FROM vehicles WHERE user_id = $1",
		updateQuery: "UPDATE vehicles SET name = $1, make = $2, model = $3, year = $4 WHERE id = $5 and user_id = $6;",
		deleteQuery: "DELETE FROM vehicles WHERE id = $1 and user_id = $2;",

		filterFields: ff,
		searchFields: sf,
	}, nil
}

func (v *vehicleDAO) Create(ctx context.Context, vehicle *dto.Vehicle) error {
	_, err := v.conn.Exec(ctx, v.createQuery, vehicle.Name, vehicle.Make, vehicle.Model, vehicle.Year, vehicle.UserID)
	return err
}

func (v *vehicleDAO) Get(ctx context.Context, id, userID int64) (*dto.Vehicle, error) {
	vehicle := &dto.Vehicle{}
	row := v.conn.QueryRow(ctx, v.getQuery, id, userID)
	if err := row.Scan(&vehicle.ID, &vehicle.Name, &vehicle.Make, &vehicle.Model, &vehicle.Year, &vehicle.UserID); err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (v *vehicleDAO) Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.Vehicle, error) {
	wc := sf.whereClause(v.searchFields, v.filterFields, 2)
	query := v.searchQuery
	args := []interface{}{userID}
	if q := wc.query(); q != "" {
		query = fmt.Sprintf("%s AND %s", v.searchQuery, wc.query())
		args = append(args, wc.args...)
	}
	var vehicles []*dto.Vehicle

	rows, err := v.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		vehicle := &dto.Vehicle{}
		if err := rows.Scan(&vehicle.ID, &vehicle.Name, &vehicle.Make, &vehicle.Model, &vehicle.Year, &vehicle.UserID); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}

func (v *vehicleDAO) Update(ctx context.Context, vehicle *dto.Vehicle, id, userID int64) error {
	_, err := v.conn.Exec(ctx, v.updateQuery, vehicle.Name, vehicle.Make, vehicle.Model, vehicle.Year, id, userID)
	return err
}

func (v *vehicleDAO) Delete(ctx context.Context, id, userID int64) error {
	_, err := v.conn.Exec(ctx, v.deleteQuery, id, userID)
	return err
}
