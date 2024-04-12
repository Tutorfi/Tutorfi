package storage

import (
	_ "database/sql"

	_ "github.com/lib/pq"
)

// func (s *Storage) GetScheduler() (models.Scheduler, error) {

// }

func (s *PostgresStorage) CreateTag(name string) error {
	//var acc models.Account
	res := s.db.QueryRow(`INSERT INTO "tags" VALUES ($1, DEFAULT, DEFAULT, DEFAULT)`, name)
	return res.Err()
}

// func (s *PostgresStorage) SetGrouptoTag(groupid string) error {
// }