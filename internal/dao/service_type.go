package dao

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
)

type ServiceType interface {
	Create(ctx context.Context, st *dto.ServiceType) error
	Get(ctx context.Context, id, userID int64) (*dto.ServiceType, error)
	Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.ServiceType, error)
	Update(ctx context.Context, st *dto.ServiceType, id, userID int64) error
	Delete(ctx context.Context, id, userID int64) error
}

type serviceTypeDAO struct {
	conn *pgxpool.Pool

	createQuery string
	getQuery    string
	searchQuery string
	updateQuery string
	deleteQuery string

	filterFields fields
	searchFields fields
}

func newServiceType(conn *pgxpool.Pool) (*serviceTypeDAO, error) {
	return &serviceTypeDAO{
		conn: conn,

		createQuery: "INSERT INTO service_types (name, freq_miles, freq_days, questions, user_id) VALUES ($1, $2, $3, $4, $5);",
		getQuery:    "SELECT id, name, freq_miles, freq_days, questions, user_id FROM service_types WHERE id = $1 AND user_id = $2;",
		searchQuery: "SELECT id, name, freq_miles, freq_days, questions, user_id FROM service_types WHERE user_id = $1",
		updateQuery: "UPDATE service_types SET name = $1, freq_miles = $2, freq_days = $3, questions = $4 WHERE id = $5 AND user_id = $6;",
		deleteQuery: "DELETE FROM service_types WHERE id = $1 AND user_id = $2;",

		filterFields: fields{"freq_days": true, "freq_miles": true},
		searchFields: fields{"name": true},
	}, nil
}

func (st *serviceTypeDAO) Create(ctx context.Context, serviceType *dto.ServiceType) error {
	_, err := st.conn.Exec(ctx, st.createQuery, serviceType.Name, serviceType.FreqMiles, serviceType.FreqDays, serviceType.Questions, serviceType.UserID)
	return err
}

func (st *serviceTypeDAO) Get(ctx context.Context, id, userID int64) (*dto.ServiceType, error) {
	serviceType := &dto.ServiceType{}
	row := st.conn.QueryRow(ctx, st.getQuery, id, userID)
	if err := row.Scan(&serviceType.ID, &serviceType.Name, &serviceType.FreqMiles, &serviceType.FreqDays, &serviceType.Questions, &serviceType.UserID); err != nil {
		return nil, err
	}
	return serviceType, nil
}

func (st *serviceTypeDAO) Select(ctx context.Context, sf SearchFilters, userID int64) ([]*dto.ServiceType, error) {
	wc := sf.whereClause(st.searchFields, st.filterFields, 2)
	query := st.searchQuery
	args := []interface{}{userID}
	if q := wc.query(); q != "" {
		query = fmt.Sprintf("%s AND %s", st.searchQuery, wc.query())
		args = append(args, wc.args...)
	}
	var serviceTypes []*dto.ServiceType
	rows, err := st.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		serviceType := &dto.ServiceType{}
		if err := rows.Scan(&serviceType.ID, &serviceType.Name, &serviceType.FreqMiles, &serviceType.FreqDays, &serviceType.Questions, &serviceType.UserID); err != nil {
			return nil, err
		}
		serviceTypes = append(serviceTypes, serviceType)
	}
	return serviceTypes, nil
}

func (st *serviceTypeDAO) Update(ctx context.Context, serviceType *dto.ServiceType, id, userID int64) error {
	_, err := st.conn.Exec(ctx, st.updateQuery, serviceType.Name, serviceType.FreqMiles, serviceType.FreqDays, serviceType.Questions, id, userID)
	return err
}

func (st *serviceTypeDAO) Delete(ctx context.Context, id, userID int64) error {
	_, err := st.conn.Exec(ctx, st.deleteQuery, id, userID)
	return err
}
