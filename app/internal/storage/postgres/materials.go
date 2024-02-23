package storage

import (
	"app/internal/models"
	_ "database/sql"
	// "fmt"
	_ "github.com/lib/pq"
)

func (s *PostgresStorage) GetMaterials() (*models.Account, error) {

	err := s.db.QueryRow("SELECT * FROM files")
	if err != nil {
		return nil, nil
	}
	return nil, nil
}