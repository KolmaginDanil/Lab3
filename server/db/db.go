package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/url"
)

type Connection struct {
	DbName         string
	User, Password string
	Host           string
	DisableSSL     bool
}

func (c *Connection) ConnectionURL() string {
	dbUrl := &url.URL{
		Scheme: "postgres",
		Host:   c.Host,
		User:   url.UserPassword(c.User, c.Password),
		Path:   c.DbName,
	}
	if c.DisableSSL {
		dbUrl.RawQuery = url.Values{
			"sslmode": []string{"disable"},
		}.Encode()
	}
	return dbUrl.String()
}

func (c *Connection) Open() (*sql.DB, error) {
	return sql.Open("postgres", c.ConnectionURL())
}
