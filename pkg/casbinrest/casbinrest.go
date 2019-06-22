package casbinrest

import (
	"fmt"
	"strings"

	"github.com/casbin/casbin"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

type (
	// Config defines the config for CasbinAuth middleware.
	Config struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// Enforcer CasbinAuth main rule.
		// Required.
		Enforcer *casbin.Enforcer
	}
)

var (
	// DefaultConfig is the default CasbinAuth middleware config.
	DefaultConfig = Config{
		Skipper: middleware.DefaultSkipper,
	}
)

// Middleware returns a CasbinAuth middleware.
//
// For valid credentials it calls the next handler.
// For missing or invalid credentials, it sends "401 - Unauthorized" response.
func Middleware(ce *casbin.Enforcer) echo.MiddlewareFunc {
	c := DefaultConfig
	c.Enforcer = ce
	return MiddlewareWithConfig(c)
}

// MiddlewareWithConfig returns a CasbinAuth middleware with config.
// See `Middleware()`.
func MiddlewareWithConfig(config Config) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) || config.CheckPermission(c) {
				return next(c)
			}

			return echo.ErrForbidden
		}
	}
}

// GetRole gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *Config) GetRole(c echo.Context) string {
	reqToken := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		// Error: Bearer token not in proper format
		return ""
	}
	reqToken = strings.TrimSpace(splitToken[1])
	fmt.Println(reqToken)

	return reqToken
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *Config) CheckPermission(c echo.Context) bool {
	user := a.GetRole(c)
	method := c.Request().Method
	path := c.Request().URL.Path
	return a.Enforcer.Enforce(user, path, method)
}