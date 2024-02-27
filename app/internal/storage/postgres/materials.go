package storage

import (
	"app/internal/models"
	"database/sql"
	_ "github.com/jackc/pgx/v5"
)

func (s *PostgresStorage) GetMaterialsForUser() (*models.Account, error) {

	err := s.db.QueryRow("SELECT * FROM files")

}

func (s *PostgresStorage) GetFilePaths() (*models.Account, error) {

	err := s.db.QueryRow("SELECT * FROM files")

}
