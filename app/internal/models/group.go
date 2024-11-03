package models

import "time"

type Group struct {
	Id             int64
	GroupID        string
	OrganizationId int64
	Name           string
	CreatedAt      time.Time
}
