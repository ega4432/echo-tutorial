# echo-tutorial

<img height="80" style="max-width:100%;" src="https://camo.githubusercontent.com/5e68fe65866b89b9e6dc9975ddef1339c92b7534/68747470733a2f2f63646e2e6c6162737461636b2e636f6d2f696d616765732f6563686f2d6c6f676f2e737667">

### Official Document

- [Echo - High performance, minimalist Go web framework](https://echo.labstack.com/)
- [labstack/echo - GitHub](https://github.com/labstack/echo)

### Practiced Feature Overview

- Optimized HTTP router which smartly prioritize routes
- coming soon ...


### Quick Start

Installation

```bash
$ go get -u github.com/labstack/echo/...
```

Create `server.go`

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

Start server

```bash
$ go run server.go
```

Browse to http://localhost:1323 and you should see Hello, World! on the page.
