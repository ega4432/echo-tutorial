package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
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

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo Instance
	e := echo.New()

	// Template
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	// Route

	// GET
	e.GET("/", home)
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

// Handler
func home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
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
