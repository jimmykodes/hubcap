package dao

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/multierr"
)

type stmt int

type statements map[stmt]*sqlx.Stmt

func (s statements) Close() error {
	var err error
	for _, stmt := range s {
		err = multierr.Append(err, stmt.Close())
	}
	return err
}
