package service

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the Postgres driver.

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"

	schema "github.com/eveld/graphql/postgres"
)

// NewDatabase connects to the Postgres databae.
func NewDatabase(DBhost string, DBport int64, DBuser string, DBpassword string, DBname string, DBmaxopenconns int, DBconnmaxlifetime time.Duration) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DBhost, DBport, DBuser, DBpassword, DBname))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Printf("Retrying database connection in 5 seconds")
		time.Sleep(time.Duration(5) * time.Second)
		return NewDatabase(DBhost, DBport, DBuser, DBpassword, DBname, DBmaxopenconns, DBconnmaxlifetime)
	}

	db.SetMaxOpenConns(DBmaxopenconns)
	db.SetConnMaxLifetime(DBconnmaxlifetime)

	return db.Unsafe(), nil
}

// Migrate handles database migrations.
func Migrate(db *sqlx.DB) {
	source := bindata.Resource(schema.AssetNames(),
		func(name string) ([]byte, error) {
			return schema.Asset(name)
		})

	data, err := bindata.WithInstance(source)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrations, err := migrate.NewWithInstance("go-bindata", data, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	version, dirty, err := migrations.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatal(err)
	}
	log.Printf(fmt.Sprintf("Database version: %d (dirty: %v)", version, dirty))
	log.Printf("Starting database migrations")

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Printf("Finished database migrations")
}
