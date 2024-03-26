package utils

import (
	"app/internal/models"
	"app/internal/storage"
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetAccountFromSessionId(s storage.Storage, c echo.Context) (*models.Account, error) {
	cookie, err := c.Cookie("Tutorfi_Account")
	if err != nil {
		return &models.Account{}, err
	}
	acc, err := s.GetAccountSessionId(cookie.Value)
	if err == sql.ErrNoRows {
		return &models.Account{}, err
	}
	if err != nil {
		fmt.Println("A database Error Occured")
		fmt.Println(err)
		return &models.Account{}, err
	}
	return acc, nil
}
