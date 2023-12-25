package account

import (
	"database/sql"
)

type account struct {
	ID int 
}

type accountModel struct {
	db *sql.DB
}

func NewAccountModel(db *sql.DB) *accountModel {
	return &accountModel{
		db:db,
	}
}

func (m *accountModel) createAccount(id int) (string, error) {

	return "Added account", nil
}

func (m *accountModel) checkAccount(user string, pass string) (bool, error) {

	return false, nil
}