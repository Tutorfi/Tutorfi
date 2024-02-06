package accounthandler

import (
	//"app/internal/models"
	//"database/sql"
	//_ "github.com/lib/pq"
	"app/internal/app"
	"app/internal/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"

	//"fmt"
	//"errors"
	"database/sql"

	"github.com/ory/dockertest/v3"
	
)

//Ok most of this is invalid bc of a change that I made, but i'm keeping it around for integration tests
func TestEmail(t *testing.T){
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
		`“email”@example.com`,}
	if(len(validEmails) == 1 || len(invalidEmails) == 1){
		t.Logf("Dummy")
	}
	//t.Logf("Email test completed")
}
//Is there a point to testing names?
func TestName(t *testing.T){
	nameRegex := `^[A-Za-z\x{00C0}-\x{00FF}][A-Za-z\x{00C0}-\x{00FF}\'\-]+([\ A-Za-z\x{00C0}-\x{00FF}][A-Za-z\x{00C0}-\x{00FF}\'\-]+)*`
	validNames := []string{
		"John",
		"Smith",
	}

	invalidNames := []string{
		"12345",
		"<Script>",
	}
	for _, element := range validNames{
		if err := checkFormValue(nameRegex, element); err != nil{
			t.Errorf("Valid names check failed with %s, got %s, expected nil.", element, err.Error())
		}
	}
	for _, element := range invalidNames{
		if err := checkFormValue(nameRegex, element); err == nil{
			t.Errorf("Valid names check failed with %s, got nil, expected %s.", element, err.Error())
		}
	}
	t.Logf("Name test completed")
}
//Options - create a new docker image in every main testing file.
//Pass down the same docker container (which is what I want to do)
//Ok nvm we can just connect to the database each time.

func TestAccountCreation(t *testing.T){
	db, err := app.ConnectPgsqlTest()
	pgstore := storage.NewPostgresStorage(db)
	SampleUsers := `{
		"User1": {
			"Firstname":"Test",
			"Lastname":"Test",
			"Email":"asdf@gmail.com",
			"Password":"password123"
		}
	}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/create-account", strings.NewReader(SampleUsers))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	accHandler := accounthandler.New(pgstore)
	if err = accHandler.CreateAccount(c); err != nil{
		t.Errorf("Account creation test failed: %s", err.Error())
	}
	t.Logf("Account creation test completed")
}
func TestAccountLogin(t *testing.T){

}
// func TestFormValid(t *testing.T){
// 	if errs := TestEmail(); errs != nil{
// 		fmt.Println("email test failed")
// 		fmt.Println(errs)
// 	}
// }
//In the future add fuzzing somewhere
//https://github.com/stretchr/testify???