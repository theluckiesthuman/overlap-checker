package main

import (
	"github.com/labstack/echo/v4"
	"github.com/theluckiesthuman/overlap-checker/internal/adapter"
)

func main() {
	e := echo.New()
	adapter.RegisterOverlapHandlers(e)
	e.Logger.Fatal(e.Start(":8080"))
}
