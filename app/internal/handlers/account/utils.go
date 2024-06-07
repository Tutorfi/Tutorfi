package accounthandler

import (
	"net/http"
	"time"
)


func createCookie(sessionid string) *http.Cookie {
	var cookie = new(http.Cookie)
	cookie.Name = "Tutorfi_Account"
	cookie.Value = sessionid
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Path = "/"
	return cookie
}
