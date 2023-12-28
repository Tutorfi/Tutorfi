package storage

import "app/internal/models"

type Storage interface {
	GetAccount(email string) (*models.Account, error)
}