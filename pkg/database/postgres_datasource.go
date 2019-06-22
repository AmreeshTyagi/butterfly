package database

// PostgresDataSource is a interface
type PostgresDataSource interface {
}

type postgresDataSource struct {
}

// NewPostgresDataSource is a instance
func NewPostgresDataSource() PostgresDataSource {
	return &postgresDataSource{}
}
