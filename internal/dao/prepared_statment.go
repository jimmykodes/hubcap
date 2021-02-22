package dao

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/multierr"
)

type stmt int

type statements map[stmt]*sqlx.Stmt

func prepareStatements(db *sqlx.DB, queries map[stmt]string) (statements, error) {
	stmts := make(statements, 0)
	for s, query := range queries {
		preparedStmt, err := db.Preparex(query)
		if err != nil {
			return nil, err
		}
		stmts[s] = preparedStmt
	}
	return stmts, nil
}

func (s statements) Close() error {
	var err error
	for _, stmt := range s {
		err = multierr.Append(err, stmt.Close())
	}
	return err
}
