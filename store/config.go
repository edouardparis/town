package store

import "fmt"

// Config are a database connection options.
type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	SSLMode  string `json:"ssl_mode"`
}

// PostgresURI returns the postgres connection URI.
func (c Config) PostgresURI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s;timezone=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.SSLMode,
		"UTC")
}
