package models

import "time"

// import "time"

type Schedule struct {
	ID        int64
	OrganizationID int64
    name string
    ScheduledAt time.Time
    time_created time.Time
}

// type UserSchedule struct {
// 	ID        int64
// 	AccountID string
// 	data      struct {
// 		Time time.Time
// 	}
// }
