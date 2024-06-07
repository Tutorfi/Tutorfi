package calendarhandler

import (
    "github.com/labstack/echo/v4"
    "app/internal/storage"
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
    userID := getUserIDFromContext(c)
    if userID == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }

    events, err := h.store.GetEventsByUserID(userID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch events")
    }

    return c.Render(http.StatusOK, "calendar.templ", map[string]interface{}{
        "Title": "Your Calendar",
        "Events": events,
    })
}