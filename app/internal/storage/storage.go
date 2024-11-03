package storage

import (
	"app/internal/models"
)

type Storage interface {
	GetAccount(email string) (*models.Account, error)
    CreateGroups(account *models.Account, group *models.Group) error 
    GetGroups(acc *models.Account) ([]models.Group, error)
    GetAccountSessionId(sessionId string) (*models.Account, error)
	CreateAccount(fname, lname, email, username, password string) error
    GetPassword(email string) (string, error)
	SetSessionID(id string, sessionid string) error
    UpdateEmail(id string, email string) error
    UpdateUsername(id string, username string) error
	DeleteAccount(id string) error
	ResetSessionID(sessionid string) error
}
