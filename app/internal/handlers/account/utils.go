package accounthandler

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
)

type accountAuth struct {
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

/*
Here are the status that can be returned:
- Authorized
- Unauthorized
- Failed (server failed)
- Success
- Invalid (user input bad)
*/
type response struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func fillResponse(status string, msg string) response {
	return response{
		Status: status,
		Msg:    msg,
	}
}

func createCookie(sessionid string, remember bool) *http.Cookie {
	var cookie = new(http.Cookie)
	cookie.Name = "Tutorfi_Account"
	cookie.Value = sessionid
	if remember {
		cookie.Expires = time.Now().AddDate(1, 0, 0)
	} else {
		cookie.Expires = time.Now().AddDate(0, 0, 1)
	}
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Path = "/"
	return cookie
}

func checkFormValue(expression, val string) error {
	regex := expression
	matcher, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println("regex failed") //What to do here?
		return err
	}
	matched := matcher.MatchString(val)
	if !matched {
		return &AccountError{msg: "Invalid form value"}
	}
	return nil
}
