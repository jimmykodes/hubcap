package dao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
)

type Vehicle interface {
	Create(ctx context.Context, v *dto.Vehicle) error
	Get(ctx context.Context, id, userID int64) (*dto.Vehicle, error)
	Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.Vehicle, error)
	Update(ctx context.Context, v *dto.Vehicle, id, userID int64) error
	Delete(ctx context.Context, id, userID int64) error
	Close() error
}

const (
	createVehicle stmt = iota
	getVehicle
	updateVehicle
	deleteVehicle
)

type vehicle struct {
	db           *sqlx.DB
	stmts        statements
	filterFields fields
	searchFields fields
	searchQuery  string
}

func newVehicle(db *sqlx.DB) (*vehicle, error) {
	queries := map[stmt]string{
		createVehicle: "INSERT INTO vehicles.vehicles (name, make, model, year, user_id) VALUE (?, ?, ?, ?, ?);",
		getVehicle:    "SELECT id, name, make, model, year, user_id FROM vehicles.vehicles WHERE id = ? AND user_id = ?;",
		updateVehicle: "UPDATE vehicles.vehicles SET name = ?, make = ?, model = ?, year = ? WHERE id = ? and user_id = ?;",
		deleteVehicle: "DELETE FROM vehicles.vehicles WHERE id = ? and user_id = ?;",
	}
	s, err := prepareStatements(db, queries)
	if err != nil {
		return nil, err
	}
	ff := fields{"make": true, "model": true, "year": true}
	sf := fields{"name": true}
	return &vehicle{
		db:           db,
		stmts:        s,
		filterFields: ff,
		searchFields: sf,
		searchQuery:  `SELECT id, name, make, model, year, user_id FROM vehicles.vehicles WHERE user_id = ?`,
	}, nil
}

func (v *vehicle) Create(ctx context.Context, vehicle *dto.Vehicle) error {
	_, err := v.stmts[createVehicle].ExecContext(ctx, vehicle.Name, vehicle.Make, vehicle.Model, vehicle.Year, vehicle.UserID)
	return err
}

func (v *vehicle) Get(ctx context.Context, id, userID int64) (*dto.Vehicle, error) {
	vehicle := &dto.Vehicle{}
	if err := v.stmts[getVehicle].GetContext(ctx, vehicle, id, userID); err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (v *vehicle) Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.Vehicle, error) {
	wc := sf.whereClause(v.searchFields, v.filterFields)
	query := v.searchQuery
	args := []interface{}{userID}
	if q := wc.query(); q != "" {
		query = fmt.Sprintf("%s AND %s", v.searchQuery, wc.query())
		args = append(args, wc.args...)
	}
	var rows []*dto.Vehicle
	if err := v.db.SelectContext(ctx, &rows, query, args...); err != nil {
		return nil, err
	}
	return rows, nil
}

func (v *vehicle) Update(ctx context.Context, vehicle *dto.Vehicle, id, userID int64) error {
	_, err := v.stmts[updateVehicle].ExecContext(ctx, vehicle.Name, vehicle.Make, vehicle.Model, vehicle.Year, id, userID)
	return err
}

func (v *vehicle) Delete(ctx context.Context, id, userID int64) error {
	_, err := v.stmts[deleteVehicle].ExecContext(ctx, id, userID)
	return err
}

func (v *vehicle) Close() error {
	return v.stmts.Close()
}
