package storage

import "app/internal/models"

type Storage interface {
	GetAccount(email string) (*models.Account, error)
	CreateAccount(fname, lname, email, password string) (error)
	GetPassword(email string) (string, error)
}