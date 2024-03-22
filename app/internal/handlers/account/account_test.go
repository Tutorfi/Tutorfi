package accounthandler

import (
	"app/internal/storage/postgres"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAccountCreation(t *testing.T) {
	db, err := storage.ConnectPgsqlTest()
	if err != nil {
		t.Fatalf(err.Error())
	}
	pgstore := storage.NewPostgresStorage(db)
	account := New(pgstore)
	SampleUsers := [][]string{
		{"Test", "Test", "asdf@gmail.com", "password123"},
		{"", "", "", ""},
	}
	f := make(url.Values)
	e := echo.New()
	for _, element := range SampleUsers {
		f.Set("fname", element[0])
		f.Set("lname", element[1])
		f.Set("email", element[2])
		f.Set("password", element[3])
		req := httptest.NewRequest(http.MethodPost, "/create-account", strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if err = account.CreateAccount(c); err != nil {
			t.Errorf("Account creation test failed: %s", err.Error())
		}
	}
	t.Logf("Account creation test completed")
}
