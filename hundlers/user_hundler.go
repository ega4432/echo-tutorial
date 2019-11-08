package hundlers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseUsers struct {
	Users []ResponseUser `json:"users"`
}

type ResponseUser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
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

// for MySQL
var (
	tableName = "users"
	seq       = 1
	conn, _   = dbr.Open("mysql", "root:@tcp(127.0.0.1:3306)/echo-tutorial", nil)
	sess      = conn.NewSession(nil)
)

func GetUsers(c echo.Context) error {
	var users ResponseUsers
	_, err := sess.Select("*").From(tableName).Load(&users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Printf("SELECT result %d\n", err)
	r := ResponseUsers{
		Users: nil,
	}
	return c.JSON(http.StatusOK, r)
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
