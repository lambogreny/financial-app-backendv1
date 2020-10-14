package database

import (
	"io"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	io.Closer
}

type database struct {
	conn *sqlx.DB
}

func (d *database) Close() error {
	return d.conn.Close()
}
