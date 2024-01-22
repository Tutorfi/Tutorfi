package models
import (
	"github.com/go-playground/validator/v11"
)

type Account struct {
	ID string `validate:required`
	SessionID *string //Session id is null when accounts are created, and only set upon login, so we keep it as a *string instead of a string
	Firstname string `validate:required`
	Lastname string `validate:required`
	Email string `validate:required`
	Password string `validate:required`
}


