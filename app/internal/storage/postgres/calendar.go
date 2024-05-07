package storage

import (
	"app/internal/models"
	"fmt"
)

// GetEventsByUserID fetches all events for a specific user from the database
func (s *PostgresStorage) GetEventsByUserID(userID int64) ([]models.Calendar, error) {
    var events []models.Calendar
    query := `SELECT id, event_title, detail, start_time, end_time
              FROM "event"
              WHERE account_id = $1`
    rows, err := s.db.Query(query, userID)
    if err != nil {
        fmt.Println("Database error occurred:", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var event models.Calendar
        if err := rows.Scan(&event.Id, &event.Event_title, &event.Detail, &event.Start_time, &event.End_time); err != nil {
            fmt.Println("Error scanning rows:", err)
            return nil, err
        }
        events = append(events, event)
    }

    if err = rows.Err(); err != nil {
        fmt.Println("Row handling error:", err)
        return nil, err
    }

    return events, nil
}