package postgre

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type DbConnection struct {
	*sqlx.DB
}

func NewConnectionDb(connStr string) (*DbConnection, error) {
	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return &DbConnection{db}, nil
}
