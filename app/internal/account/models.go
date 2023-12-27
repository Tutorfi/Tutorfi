package account

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type account struct {
	ID int64
	Firstname string
	Lastname string
	Email string
	Password string
}

type accountModel struct {
	db *sql.DB
}

func NewAccountModel(db *sql.DB) *accountModel {
	return &accountModel{
		db: db,
	}
}

func (m *accountModel) createAccount(id int) (bool, error) {
	return false, nil
}

func (m *accountModel) checkAccount(user string, pass string) (bool, error) {
	var ac account
	row := m.db.QueryRow("SELECT * FROM Account WHERE Username = $1", user)
	if row != nil {
		return false, nil
	}
	err := row.Scan(&ac.ID, &ac.Firstname, &ac.Lastname, &ac.Email, &ac.Password)
	if err != nil {
		return false, err
	}

	return false, nil
}
