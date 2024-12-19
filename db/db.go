package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Db *sql.DB
}

// Create a database connection and store
func NewStore(dbName string) (Store, error) {
	// Establish connection to the database
	Db, err := getConnection(dbName)

	if err != nil {
		// Return error and empty store
		return Store{}, err
	}

	// Migrate database if necessary
	err = createMigrations(dbName, Db)
	if err != nil {
		// Return error and empty store
		return Store{}, err
	}

	// Return new database store
	return Store{Db}, nil
}

func getConnection(dbName string) (*sql.DB, error) {
	var (
		err error
		db  *sql.DB
	)

	// Return database if it's aleady intialised
	// QUESTION: How? I initialise the variable just above, how does that work?
	// Does it carry over from repeated calls because it's a pointer?
	if db != nil {
		return db, nil
	}

	// Open database connection, return error if failed
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("🔥 failed to connect to the database: %s", err)
	}

	// Successfully connected
	log.Println("🚀 Connected Successfully to the Database")

	// Return database
	return db, nil
}

func createMigrations(dbName string, db *sql.DB) error {
	stmt := `
		CREATE TABLE IF NOT EXISTS person (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			name TEXT NOT NULL
		);
	`

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `
		CREATE TABLE IF NOT EXISTS task (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_by INTEGER NOT NULL,
			title TEXT NOT NULL,
			assignee INTEGER NOT NULL,
			reviewer INTEGER NOT NULL,
			completed INTEGER DEFAULT(FALSE),
			reviewed INTEGER DEFAULT(FALSE),
			due_on TEXT,
			created_at TEXT DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(created_by) REFERENCES person(id),
			FOREIGN KEY(assignee) REFERENCES person(id),
			FOREIGN KEY(reviewer) REFERENCES person(id)
		);
	`

	_, err = db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}
