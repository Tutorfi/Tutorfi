package calendarhandler

import (
	"app/internal/storage"
	"github.com/labstack/echo/v4"
	"net/http"
)

type calendarHandler struct {
	store storage.Storage
}

func New(store storage.Storage) *calendarHandler {
	return &calendarHandler{store: store}
}

func (h *calendarHandler) HandleGetCalendar(c echo.Context) error {
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
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch events")
	}

	return c.Render(http.StatusOK, "calendar.templ", map[string]interface{}{
		"Title":  "Your Calendar",
		"Events": events,
	})
}
