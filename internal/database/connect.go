package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/namsral/flag"
	"github.com/pkg/errors"
	//"github.com/sirupsen/logrus"
)

var (
	databaseURL    = flag.String("databse-url", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable", "Database URL.")
	databseTimeout = flag.Int64("database-timeout-ms", 000, "")
)


//Connect creates a new database connection
func Connect() (*sqlx.DB, error) {
	//connect to database.
	dbURL := *databaseURL


	logrus.WithField("url", dbURL).Debug("Connecting to database.")

	conn, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return nil, errors.Wrap(err, "Could not connect to database")
	}

	conn.SetMaxOpenConns(32)

	//Check if database running
	if err := waitForDB(conn.DB); err != nil {
		return nil, err
	}

	//Migrate database schema
	/*
	if err := migrateDb(conn.DB); err != nil {
		return nil, errors.Wrap(err, "cloud not migrate database")
	}

	 */

	//conn *sqlx.DBsqlx.Open("postgres", dbURL)

	return conn, nil
}

//New Creates a new database
func New() (Database, error) {
	conn, err := Connect()
	if err != nil {
		return nil, err
	}

	d := &database{
		conn: conn,
	}

	return d, nil
}

func waitForDB(conn *sql.DB) error {
	ready := make(chan struct{})
	go func() {
		for {
			if err := conn.Ping(); err != nil {
				close(ready)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {
	case <-ready:
		return nil
	case <-time.After(time.Duration(*databseTimeout) * time.Second):
		return errors.New("database not ready")
	}
}
