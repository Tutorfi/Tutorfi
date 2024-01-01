package pages

import(
	"fmt"
	"html/template"
	"github.com/labstack/echo/v4"
)
func login (c echo.Context) error {
	fmt.Println("Got a GET request")
	if len(c.Cookies()) == 0{
		temp, e := template.ParseFiles("./public/templates/login.html")
		if e != nil {
			fmt.Println("Error parsing template")
			fmt.Println(e)
			return e
		}
		tmpl := template.Must(temp, e)
		return tmpl.Execute(c.Response().Writer, nil)	 
	}
	return nil
	
}