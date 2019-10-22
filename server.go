package main

import (
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

const helloText string = `Hello, World from echo!
echo is high performance, extensible, minimalist Go web framework`

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, helloText)
}

func main() {
	// Echo Instance
	e := echo.New()

	// Route
	e.GET("/", hello)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
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
	return c.String(http.StatusOK, "team,:"+team+", member:"+member)
}

func save(c echo.Context) error {
	// get name
	name := c.FormValue("name")
	// get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
