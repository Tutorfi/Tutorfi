package storage

import (
	_ "database/sql"
)

type PostgresError struct {
	message  string
	sqlError error //In the event we need to access the raw error.
}

func (se *PostgresError) Error() string {
	return se.message
}

func (se *PostgresError) SQLError() error {
	return se.sqlError
}
