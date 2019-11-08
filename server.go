package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yoshimitsuEgashira/echo-tutorial/hundlers"
	"html/template"
	"io"
)

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

	// Root Level Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Group Level Middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (b bool, e error) {
		if username == "Joe" && password == "secret" {
			return true, nil
		} else {
			return false, nil
		}
	}))

	// Route
	// GET
	e.GET("/", hundlers.Home)
	e.GET("/users", hundlers.GetUsers)
	e.GET("/users/:name", hundlers.GetUserName)
	e.GET("/show", hundlers.Show)

	// POST
	e.POST("/save", hundlers.Save)
	e.POST("/users", hundlers.SaveUser)
	e.POST("/send", hundlers.SendMessage)

	// ToDo: Others ( PUT/DELETE )
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
