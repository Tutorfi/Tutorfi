package storage

import (
	"app/internal/models"
)

type Storage interface {
	GetAccount(email string) (*models.Account, error)
    GetGroups(acc *models.Account) ([]models.Group, error)
    GetAccountSessionId(sessionId string) (*models.Account, error)
	CreateAccount(fname, lname, email, password string) error
    GetPassword(email string) (string, error)
	SetSessionID(id string, sessionid string) error
	DeleteAccount(id string) error
	ResetSessionID(sessionid string) error
}
