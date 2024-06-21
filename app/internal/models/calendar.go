package models

import "time"

type Event struct {
	Id          int64
	Account_id  string
	Event_title string
	Detail      string
	Start_time  time.Time
	End_time    *time.Time
}