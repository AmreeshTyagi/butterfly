package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/prongbang/butterfly/pkg/casbinrest"
	"github.com/prongbang/butterfly/pkg/di"
)

func main() {

	rc := di.ProvideRedisClient()
	tokenID, _ := rc.Get("token_id").Result()
	log.Println(tokenID)

	ce := di.ProvideEnforcer()
	// e.AddPolicy("admin", "/*", "*")
	ce.AddPolicy("anonymous", "/login", "GET")
	// e.AddPolicy("anonymous", "/login", "POST")
	// e.AddPolicy("member", "/logout", "GET")
	// e.AddPolicy("member", "/member", "*")
	ce.LoadPolicy()
	ce.GetPolicy()

	role := "anonymous"
	path := "/login"
	method := "POST"
	result := ce.Enforce(role, path, method)
	log.Println(result)

	myRes := ce.GetPolicy()
	log.Print("Policy: ", myRes)

	e := echo.New()
	e.Debug = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(casbinrest.Middleware(ce))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
