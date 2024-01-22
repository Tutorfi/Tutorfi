package models
import (
	_ "github.com/go-playground/validator/v10"
)

type Account struct {
	ID string `validate:"required"`
	SessionID *string //Session id is null when accounts are created, and only set upon login, so we keep it as a *string instead of a string
	Firstname string `validate:"required"`
	Lastname string `validate:"required"`
	Email string `validate:"required,email"`
	Password string `validate:"required"`
}


