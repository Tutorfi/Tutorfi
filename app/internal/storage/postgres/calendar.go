package storage

import (
	"app/internal/models"
	"database/sql"
	_ "database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func (s *PostgresStorage) GetEvent(id int64) (*models.Calendar, error) {
	var event models.Calendar
	err := s.db.QueryRow(`
		SELECT event_title, detail, start_time, end_time
		FROM "event" WHERE "id" = $1`, id).Scan(
			&event.Event_title, &event.Detail, 
			&event.Start_time, & event.End_time)
	if err == sql.ErrNoRows {
		return &models.Calendar{}, nil
	}
	if err != nil {
		fmt.Println("A database Error Occured")
		fmt.Println(err)
		return &models.Calendar{}, err
	}
	return &event, nil
}