package main

import (
	"html/template"
	"io"

	"github.com/yoshimitsuEgashira/echo-tutorial/twitter-api/handlers"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("./public/views/*.html")),
	}
	e.Renderer = t
	e.POST("/search", handlers.SearchTweet)
	e.GET("/searchTweet", handlers.GetSearch)
	e.Logger.Fatal(e.Start(":1232"))
}
