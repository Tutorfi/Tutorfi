package storage

import (
	"app/internal/models"
	"fmt"
)

func (s *PostgresStorage) GetGroups(account *models.Account) ([]models.Group, error) {
    groups := make([]models.Group, 0, 10)
    rows, err := s.db.Query(`
    SELECT "id","organization_id","name"
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
        err = rows.Scan(&group.Id, &group.OrganizationId, &group.Name)
        if err != nil {
            return nil, err
        }
        groups = append(groups, group)
    }

    return groups, nil
}
