package hundlers

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Response struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func GetUserName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

func Show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team : "+team+" , member : "+member)
}

func Save(c echo.Context) error {
	// get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b><br>Your email address is : "+email)
}

func SaveUser(c echo.Context) error {
	u := new(User)
	err := c.Bind(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func SendMessage(c echo.Context) error {
	m := new(Message)
	if err := c.Bind(m); err != nil {
		return err
	}
	r := new(Response)
	r.Name = m.Name
	r.Email = m.Email
	r.Message = m.Message
	r.Status = "success!"
	return c.JSON(http.StatusOK, r)
}
