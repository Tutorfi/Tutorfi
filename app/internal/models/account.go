package models


type Account struct {
	ID string
	SessionID *string //Session id is null when accounts are created, and only set upon login, so we keep it as a *string instead of a string
	Firstname string
	Lastname string
	Email string
	Password string
}


