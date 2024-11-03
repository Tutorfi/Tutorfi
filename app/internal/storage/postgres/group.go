package storage

import (
	"app/internal/models"
	"crypto/rand"
	"fmt"
)

func tokenGenerator() (string, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "error", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (s *PostgresStorage) GetGroups(account *models.Account) ([]models.Group, error) {
	groups := make([]models.Group, 0, 10)
	rows, err := s.db.Query(`
    SELECT "id","group_id","organization_id","name","created_at"
    FROM "group" 
    WHERE 
    "id" = (
        SELECT "group_id" 
        FROM "group_account" 
        WHERE "account_id" = $1
    )`, account.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		group := models.Group{}
		err = rows.Scan(&group.Id, &group.GroupID, &group.OrganizationId, &group.Name, &group.CreatedAt)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func (s *PostgresStorage) CreateGroups(account *models.Account, group *models.Group) error {

	token, err := tokenGenerator()
	if err != nil {
		fmt.Println("Can't generate token")
		fmt.Println(err)
		return err
	}

	_, err = s.db.Exec(`
    INSERT "group" ("group_id","organization_id", "name") VALUES ($1,$2,$3);
    `, token, account.OrganizationId, group.Name)
	if err != nil {
		fmt.Println("Insert error occured")
		fmt.Println(err)
		return err
	}

	rows := s.db.QueryRow(`
    SELECT "id" FROM "group" WHERE "group_id" = $1
    `, token)
	group_id := new(int32)
	err = rows.Scan(&group_id)
	if err != nil {
		fmt.Println("No rows exist even though created")
		fmt.Println(err)
		return err
	}

	_, err = s.db.Exec(`
    INSERT "group_account" ("group_id", "account_id") VALUES ($1,$2)
    `, group_id, account.Id)
	if err != nil {
		fmt.Println("Insert error occured")
		fmt.Println(err)
		return err
	}

	return nil
}
