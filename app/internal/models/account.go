package models


type Account struct {
	ID string
	SessionID *string
	Firstname string
	Lastname string
	Email string
	Password string
}

func (a *Account) verify () error{
	return nil
}

