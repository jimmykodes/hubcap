package dao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
)

type ServiceType interface {
	Create(ctx context.Context, st *dto.ServiceType) error
	Get(ctx context.Context, id, userID int64) (*dto.ServiceType, error)
	Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.ServiceType, error)
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
	db           *sqlx.DB
	stmts        statements
	filterFields fields
	searchFields fields
	searchQuery  string
}

func newServiceType(db *sqlx.DB, database string) (*serviceType, error) {
	q := queries{
		createServiceType: fmt.Sprintf("INSERT INTO %s.service_types (name, freq_miles, freq_days, questions, user_id) VALUE (?, ?, ?, ?, ?);", database),
		getServiceType:    fmt.Sprintf("SELECT id, name, freq_miles, freq_days, questions, user_id FROM %s.service_types WHERE id = ? AND user_id = ?;", database),
		updateServiceType: fmt.Sprintf("UPDATE %s.service_types SET name = ?, freq_miles = ?, freq_days = ?, questions = ? WHERE id = ? AND user_id = ?;", database),
		deleteServiceType: fmt.Sprintf("DELETE FROM %s.service_types WHERE id = ? AND user_id = ?;", database),
	}
	s, err := prepareStatements(db, q)
	if err != nil {
		return nil, err
	}
	return &serviceType{
		db:           db,
		stmts:        s,
		filterFields: fields{"freq_days": true, "freq_miles": true},
		searchFields: fields{"name": true},
		searchQuery:  fmt.Sprintf("SELECT id, name, freq_miles, freq_days, questions, user_id FROM %s.service_types WHERE user_id = ?", database),
	}, nil
}

func (st *serviceType) Create(ctx context.Context, serviceType *dto.ServiceType) error {
	_, err := st.stmts[createServiceType].ExecContext(ctx, serviceType.Name, serviceType.FreqMiles, serviceType.FreqDays, serviceType.Questions, serviceType.UserID)
	return err
}

func (st *serviceType) Get(ctx context.Context, id, userID int64) (*dto.ServiceType, error) {
	obj := &dto.ServiceType{}
	if err := st.stmts[getServiceType].GetContext(ctx, obj, id, userID); err != nil {
		return nil, err
	}
	return obj, nil
}

func (st *serviceType) Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.ServiceType, error) {
	wc := sf.whereClause(st.searchFields, st.filterFields)
	query := st.searchQuery
	args := []interface{}{userID}
	if q := wc.query(); q != "" {
		query = fmt.Sprintf("%s AND %s", st.searchQuery, wc.query())
		args = append(args, wc.args...)
	}
	var rows []*dto.ServiceType
	if err := st.db.SelectContext(ctx, &rows, query, args...); err != nil {
		return nil, err
	}
	return rows, nil
}

func (st *serviceType) Update(ctx context.Context, serviceType *dto.ServiceType, id, userID int64) error {
	_, err := st.stmts[updateServiceType].ExecContext(ctx, serviceType.Name, serviceType.FreqMiles, serviceType.FreqDays, serviceType.Questions, id, userID)
	return err
}

func (st *serviceType) Delete(ctx context.Context, id, userID int64) error {
	_, err := st.stmts[deleteServiceType].ExecContext(ctx, id, userID)
	return err
}

func (st *serviceType) Close() error {
	return st.stmts.Close()
}
