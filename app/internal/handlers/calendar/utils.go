package calendarhandler

import "time"

type Calendar struct {
    Event_title string 		`json:"title"`
    Detail 		string 		`json:"detail"`
	// try to nest these two structs into one struct (start and end time)
	Start_time  time.Time	`json:"start"`
    End_time    *time.Time	`json:"end"`
}