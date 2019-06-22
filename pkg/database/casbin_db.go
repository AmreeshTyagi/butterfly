package database

import (
	"fmt"
	"log"

	"github.com/casbin/casbin"
	sqlxadapter "github.com/memwey/casbin-sqlx-adapter"
)

// CasbinDatabase is a interface
type CasbinDatabase interface {
	NewEnforcer() *casbin.Enforcer
}

type casbindatabase struct {
	Config CasbinConfig
}

// NewCasbinDatabase is a instance
func NewCasbinDatabase(config CasbinConfig) CasbinDatabase {
	return &casbindatabase{
		Config: config,
	}
}

func (c *casbindatabase) NewEnforcer() *casbin.Enforcer {
	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Config.Host,
		c.Config.Port,
		c.Config.User,
		c.Config.Pass,
		c.Config.DatabaseName,
		c.Config.SSLMode,
	)
	log.Println(dataSource)
	a := sqlxadapter.NewAdapter(c.Config.DriverName, dataSource)
	ce := casbin.NewEnforcer(c.Config.PolicyFile, a)
	ce.EnableAutoSave(true)
	return ce
}
