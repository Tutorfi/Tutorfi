/*
Contains the base handlers for scheduling and creating events.
*/
package schedulerhandler

import (
	"app/internal/storage"
	"fmt"

	"github.com/labstack/echo/v4"
)

type SchedulerHandler struct {
	store storage.Storage
}

func New(store storage.Storage) *SchedulerHandler {
	return &SchedulerHandler{
		store: store,
	}
}

func (handle *SchedulerHandler) Schedule(c echo.Context) error {
	fmt.Println("Scheduler")
	return nil
}
