package accounthandler

import (
	"app/internal/storage/postgres"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"github.com/labstack/echo/v4"
)

// Ok most of this is invalid bc of a change that I made, but i'm keeping it around for integration tests
func TestEmail(t *testing.T) {
	//https://gist.github.com/cjaoude/fd9910626629b53c4d25
	validEmails := []string{"email@example.com",
		"firstname.lastname@example.com",
		"email@subdomain.example.com",
		"firstname+lastname@example.com",
		"email@123.123.123.123",
		"email@[123.123.123.123]",

		"1234567890@example.com",
		"email@example-one.com",
		"_______@example.com",
		"email@example.name",
		"email@example.museum",
		"email@example.co.jp",
		"firstname-lastname@example.com"}

	invalidEmails := []string{
		"plainaddress",
		"#@%^%#$@#$@#.com",
		"@example.com",
		"Joe Smith <email@example.com>",
		"email.example.com",
		"email@example@example.com",
		".email@example.com",
		"email.@example.com",
		"email..email@example.com",
		"あいうえお@example.com",
		"email@example.com (Joe Smith)",
		"email@example",
		"email@-example.com",
		"email@example.web",
		"email@111.222.333.44444",
		"email@example..com",
		"Abc..123@example.com",
		`“email”@example.com`}
	if len(validEmails) == 1 || len(invalidEmails) == 1 {
		t.Logf("Dummy")
	}
	//t.Logf("Email test completed")
}


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