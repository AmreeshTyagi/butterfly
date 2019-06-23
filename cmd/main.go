package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/prongbang/butterfly/pkg/di"
	"github.com/prongbang/casbinrest"
)

type mockDataSource struct {
}

// NewMockDataSource is the instance
func NewMockDataSource() casbinrest.DataSource {
	return &mockDataSource{}
}

const mockAdminToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

func (r *mockDataSource) GetRoleByToken(reqToken string) string {
	role := "anonymous"
	if reqToken == mockAdminToken {
		role = "admin"
	}
	return role
}

func main() {
	mockSource := NewMockDataSource()

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

	// myRes := ce.GetPolicy()
	// log.Print("Policy: ", myRes)

	e := echo.New()
	e.Debug = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(casbinrest.Middleware(ce, mockSource))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
