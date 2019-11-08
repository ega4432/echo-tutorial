package main

import (
	"github.com/labstack/echo"
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

	// Route
	// GET
	e.GET("/", hundlers.Home)
	e.GET("/users/:id", hundlers.GetUser)
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
