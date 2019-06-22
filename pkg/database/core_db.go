package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Config is a model
type Config struct {
	DriverName   string
	Host         string
	Port         int
	User         string
	Pass         string
	DatabaseName string
	SSLMode      string
}

// CasbinConfig is a model
type CasbinConfig struct {
	DriverName   string
	Host         string
	Port         int
	User         string
	Pass         string
	DatabaseName string
	SSLMode      string
	PolicyFile   string
}

// Database is a interface
type Database interface {
	Connect() *sqlx.DB
}

type coreDatabase struct {
	Config Config
}

// NewCoreDatabase is a instance
func NewCoreDatabase(config Config) Database {
	return &coreDatabase{
		Config: config,
	}
}

func (c *coreDatabase) Connect() *sqlx.DB {
	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Config.Host,
		c.Config.Port,
		c.Config.User,
		c.Config.Pass,
		c.Config.DatabaseName,
		c.Config.SSLMode,
	)
	db, err := sqlx.Connect(c.Config.DriverName, dataSource)
	if err != nil {
		panic(err)
	}
	return db
}
