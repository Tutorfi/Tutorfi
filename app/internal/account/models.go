package account

import (
	"database/sql"
	"fmt"
)

type account struct {
	ID int
}

type accountModel struct {
	db *sql.DB
}

func NewAccountModel(db *sql.DB) *accountModel {
	insert := account{53}
	fmt.Println(insert.ID)
	return &accountModel{
		db: db,
	}
}

func (m *accountModel) createAccount(id int) (bool, error) {
	insert := account{53}
	fmt.Println(insert.ID)
	return false, nil
}

func (m *accountModel) checkAccount(user string, pass string) (bool, error) {
	// Add check here
	return false, nil
}
