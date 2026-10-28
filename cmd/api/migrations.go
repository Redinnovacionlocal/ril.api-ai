package main

import (
	"database/sql"
	"errors"
	"io/fs"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func runMigrations(databaseURL string) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Error initializing driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Error initializing migrations:", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		if errors.Is(err, fs.ErrNotExist) {
			version, _, _ := m.Version()
			log.Printf("WARNING: la DB está en la versión %d, que no existe en migrations/; se omiten las migraciones", version)
			return
		}
		log.Fatal("Error running migrations:", err)
	}

	log.Println("Migrations applied successfully")
}
