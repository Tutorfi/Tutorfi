package storage

import (
	"app/internal/models"

)

type MemoryStorage struct {}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) GetAccount(email string) (*models.Account, error) {

	return &models.Account{}, nil
}