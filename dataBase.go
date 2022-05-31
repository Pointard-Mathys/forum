package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func selectAllFromTable(db *sql.DB, table string) *sql.Rows {
	query := "SELECT * FROM " + table
	result, _ := db.Query(query)
	return result
}

func initDatabase(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := ` 
	PRAGMA foreign_keys = ON;
		CREATE TABLE IF NOT EXISTS users (
			id	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
		`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return db
}

func insertIntoUsers(db *sql.DB, name string, email string, password string) (int64, error) {
	result, err := db.Exec(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`, name, email, password)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return -1, nil
	}

	return result.LastInsertId()
}

func displayUserRows(rows *sql.Rows) {
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
	}
}

// func selectUserById(db *sql.DB, id int) struct {
//     var user User
// 	db.QueryRow(`SELECT * FROM types WHERE id = ?`, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
// 	return u

// }
func main() {
	db := initDatabase("user.db")
	defer db.Close()
	insertIntoUsers(db, "Mathieu", "m.m@gmail.com", "abcde")
	insertIntoUsers(db, "Thomas", "t.t@gmail.com", "fghij")
	insertIntoUsers(db, "Lucas", "l.l@gmail.com", "klmno")
	insertIntoUsers(db, "vanessa", "vanessa@gmail.com", "coeur")

	// user := selectTypesById(db, 4)
	// fmt.Println(user)

	rows := selectAllFromTable(db, "users")
	displayUserRows(rows)
	// fmt.Println(rows)

}
