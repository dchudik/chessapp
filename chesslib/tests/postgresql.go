package test

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	// for migration
	// "github.com/golang-migrate/migrate"
	// "github.com/golang-migrate/migrate/database/postgres"
	// migrate "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"

	// "github.com/golang-migrate/migrate/v4/source/pkger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var psql *sqlx.DB

type DatabaseClient struct {
	Host             string
	Port             string
	User             string
	Password         string
	Database         string
	ConnectionString string
	SSL              bool
}

var onceDB sync.Once

func connectPostgresOnce() {
	onceDB.Do(connectPostgres)
}

var connString string

// TODO: use rise to bundle migrations with the compiled binary
// Setup first migrates then init DB
func connectPostgres() {
	var err error
	var dc DatabaseClient
	if dc.Host = os.Getenv("POSTGRES_HOST"); dc.Host == "" {
		log.Fatal("DatabaseClient: No set POSTGRES_HOST")
	}

	if dc.Port = os.Getenv("POSTGRES_PORT"); dc.Port == "" {
		log.Fatal("DatabaseClient: No set POSTGRES_PORT")
	}

	if dc.User = os.Getenv("POSTGRES_USER"); dc.User == "" {
		log.Fatal("DatabaseClient: No set POSTGRES_USER")
	}

	if dc.Password = os.Getenv("POSTGRES_PASSWORD"); dc.Password == "" {
		log.Fatal("DatabaseClient: No set POSTGRES_PASSWORD")
	}

	if dc.Database = os.Getenv("POSTGRES_DATABASE"); dc.Database == "" {
		log.Fatal("DatabaseClient: No set POSTGRES_DATABASE")
	}

	if dc.SSL, err = strconv.ParseBool(os.Getenv("POSTGRES_SSL")); err != nil {
		log.Fatal(err)
	}
	dc.ConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dc.User, dc.Password, dc.Host, dc.Port, dc.Database)
	if dc.SSL {
		dc.ConnectionString += "?sslmode=verify-full&sslrootcert=pg.crt"
	} else {
		dc.ConnectionString += "?sslmode=disable"
	}

	connString = dc.ConnectionString

	psql, err = sqlx.Connect("postgres", connString)

	for err != nil {
		psql, err = sqlx.Connect("postgres", connString)
		time.Sleep(time.Second * 5)
	}

	log.Println("DatabaseClient: connected")
}
