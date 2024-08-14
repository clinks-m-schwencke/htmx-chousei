package lib

import (
	"database/sql"
	// "fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func OpenDatabase() {
	var err error
	// Db, err = sql.Open("sqlite3", "db/test.db")
	Db, err = sql.Open("sqlite3", "db/task.db")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// sts := `
	// DROP TABLE IF EXISTS cars;
	// CREATE TABLE cars(id INTEGER PRIMARY KEY, name TEXT, price INT);
	// INSERT INTO cars(name, price) VALUES('Audi', 52642);
	// INSERT INTO cars(name, price) VALUES('Mercedes', 57127);
	// INSERT INTO cars(name, price) VALUES('Skoda', 9000);
	// INSERT INTO cars(name, price) VALUES('Volvo', 29000);
	// INSERT INTO cars(name, price) VALUES('Bentley', 350000);
	// INSERT INTO cars(name, price) VALUES('Citreon', 21000);
	// INSERT INTO cars(name, price) VALUES('Hummer', 41400);
	// INSERT INTO cars(name, price) VALUES('Volkswagen', 21600);
	// `
	//
	// _, err = db.Exec(sts)

	// rows, err := Db.Query("SELECT * FROM cars")
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	//
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	var price int
	//
	// 	err = rows.Scan(&id, &name, &price)
	//
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%d %s %d\n", id, name, price)
	// }

	// fmt.Println("table cars created")
}

// func Query(sql string) {
// 	rows, err := Db.Query(sql)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	result := []
// 	for rows.Next() {
// 		err = rows.Columns
// 	}
// }

// func createTables() {
// 	sql := `CREATE TABLE person (
// 	person_id INTEGER PRIMARY KEY,
// 	name TEXT NOT NULL,
// 	email TEXT NOT NULL,
// 	password TEXT NOT NULL
// 	);`
//
// 	_, err := db.Exec(sql)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	sql = `CREATE TABLE task (
// 	task_id INTEGER PRIMARY KEY,
// 	title TEXT NOT NULL,
// 	author INTEGER NOT NULL,
// 	assigned INTEGER,
// 	due_at DATETIME
// 	)`
// }

// func execSql() error {
//
// 	sql := `CREATE TABLE events (
// 		id INTEGER PRIMARY KEY,
// 		source TEXT NOT NULL,
// 		payload JSON NOT NULL
// 	);`
//
// 	_, err := db.Exec(sql)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
