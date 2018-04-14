package database

import "github.com/jmoiron/sqlx"

// Operations is the set of operations that can be done on all pixel.horse data.
type Operations interface {
	UserOperations

	Close() error
}

// PostgresOperations is the wrapper type for the set of operations that can be
// done on data stored in postgres.
type PostgresOperations struct {
	durl string
	db   *sqlx.DB
}
