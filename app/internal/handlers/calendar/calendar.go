package calendarhandler

import (
	"app/internal/public/views/calendar"
	"app/internal/storage"
	"fmt"
	"github.com/labstack/echo/v4"
)

type Calendar struct {
	store storage.Storage
}

func New(store storage.Storage) *Calendar {
	return &Calendar{
		store: store,
	}
}

// func (handle *calendarHandler) calendarting(){
// 	//Return a template
// }
