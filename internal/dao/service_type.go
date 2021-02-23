package dao

import (
	"context"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jmoiron/sqlx"
)

type ServiceType interface {
	Create(ctx context.Context, st *dto.ServiceType) error
	Get(ctx context.Context, id, userID int64) (*dto.ServiceType, error)
	Select(ctx context.Context, sf SearchFilters) ([]*dto.ServiceType, error)
	Update(ctx context.Context, st *dto.ServiceType, id, userID int64) error
	Delete(ctx context.Context, id, userID int64) error
	Close() error
}

const (
	createServiceType stmt = iota
	getServiceType
	updateServiceType
	deleteServiceType
)

type serviceType struct {
	db    *sqlx.DB
	stmts statements
}

func newServiceType(db *sqlx.DB) (*serviceType, error) {
	q := queries{
		createServiceType: "INSERT INTO vehicles.service_types (name, freq_miles, freq_days, user_id) VALUE (?, ?, ?, ?);",
		getServiceType:    "SELECT id, name, freq_miles, freq_days, user_id FROM vehicles.service_types WHERE id = ? AND user_id = ?;",
		updateServiceType: "UPDATE vehicles.service_types SET name = ?, freq_miles = ?, freq_days = ? WHERE id = ? AND user_id = ?;",
		deleteServiceType: "DELETE FROM vehicles.service_types WHERE id = ? AND user_id = ?;",
	}
	s, err := prepareStatements(db, q)
	if err != nil {
		return nil, err
	}
	return &serviceType{db: db, stmts: s}, nil
}

func (st *serviceType) Create(ctx context.Context, serviceType *dto.ServiceType) error {
	_, err := st.stmts[createServiceType].ExecContext(ctx, serviceType.Name, serviceType.FreqMiles, serviceType.FreqDays, serviceType.UserID)
	return err
}

func (st *serviceType) Get(ctx context.Context, id, userID int64) (*dto.ServiceType, error) {
	obj := &dto.ServiceType{}
	if err := st.stmts[getServiceType].GetContext(ctx, obj, id, userID); err != nil {
		return nil, err
	}
	return obj, nil
}

func (st *serviceType) Select(ctx context.Context, sf SearchFilters) ([]*dto.ServiceType, error) {
	panic("implement me")
}

func (st *serviceType) Update(ctx context.Context, serviceType *dto.ServiceType, id, userID int64) error {
	_, err := st.stmts[updateServiceType].ExecContext(ctx, serviceType.Name, serviceType.FreqMiles, serviceType.FreqDays, id, userID)
	return err
}

func (st *serviceType) Delete(ctx context.Context, id, userID int64) error {
	_, err := st.stmts[deleteServiceType].ExecContext(ctx, id, userID)
	return err
}

func (st *serviceType) Close() error {
	return st.stmts.Close()
}
