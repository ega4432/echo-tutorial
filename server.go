package main

import (
	"github.com/labstack/echo"
	"net/http"
)

const helloText string = `Hello, World from echo!
echo is high performance, extensible, minimalist Go web framework`

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

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, helloText)
}

func main() {
	// Echo Instance
	e := echo.New()

	// Route

	// GET
	e.GET("/", hello)
	e.GET("/users/:id", getUser)
	e.GET("/users/:name", getUserName)
	e.GET("/show", show)

	// POST
	e.POST("/save", save)
	e.POST("/users", saveUser)
	e.POST("/send", sendMessage)

	// ToDo: Others ( PUT/DELETE )
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func getUserName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team : "+team+" , member : "+member)
}

func save(c echo.Context) error {
	// get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b><br>Your email address is : "+email)
}

func saveUser(c echo.Context) error {
	u := new(User)
	err := c.Bind(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func sendMessage(c echo.Context) error {
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
