package models

import ()

type Account struct {
	Id             string
	SessionId      *string //Session id is null when accounts are created, and only set upon login, so we keep it as a *string instead of a string
	OrganizationId string
	Firstname      string
	Lastname       string
	Email          string
	Password       string
}
