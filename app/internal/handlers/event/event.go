package eventhandler

import (
	"app/internal/storage"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type eventHadler struct {
	store storage.Storage
}

func New(store storage.Storage) *eventHadler {
	return &eventHadler{store: store}
}

func (h *eventHadler) HandleGetEvent(c echo.Context) error {
	// Assuming you can extract userID from context or session
	cookie, err := c.Cookie("Tutorfi_Account")
	if err != nil {
		return c.Redirect(302, "/login")
	}
	sessionId := cookie.Value
	account, err := h.store.GetAccountSessionId(sessionId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch session ID")
	}

	events, err := h.store.GetEventsByUserID(account.Id)
    if err != sql.ErrNoRows {
        return echo.NewHTTPError(http.StatusOK, "No events")
    }
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, "Unable to fetch events")
	}

	var calendarEvents []event

	for _, e := range events {
		thing := event{
			Event_title: e.Event_title,
			Detail:      e.Detail,
			Start_time:  e.Start_time,
			End_time:    e.End_time,
		}
		calendarEvents = append(calendarEvents, thing)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Title":  "Your Calendar",
		"Events": calendarEvents,
	})
}
