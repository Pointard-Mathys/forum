package forum

import (
	"database/sql"
	"encoding/base32"
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

func initDatabase(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := ` 
	CREATE TABLE IF NOT EXISTS users (
		id_user	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
		);
	CREATE TABLE IF NOT EXISTS topic (
			id_topic	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			titre TEXT NOT NULL,
			contain TEXT NOT NULL,
			id_user INTEGER NOT NULL,
			FOREIGN KEY(id_user) REFERENCES users(id_user)
		);
		`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}

	fmt.Println(selectAllFromTable(db, "topic"))

	return db
}

func selectAllFromTable(db *sql.DB, table string) *sql.Rows {
	query := "SELECT * FROM " + table
	result, _ := db.Query(query)
	return result
}

func insertIntoUsers(db *sql.DB, name string, email string, password string) (int64, error) {
	result, err := db.Exec(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`, name, email, password)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return -1, nil
	}

	return result.LastInsertId()
}

func encoding_password(a string) string {
	b := base32.StdEncoding.EncodeToString([]byte(a))
	return string(b)
}

func insertIntoTopic(db *sql.DB, titre string, contain string, id_user int) (int64, error) {
	result, err := db.Exec(`INSERT INTO topic (titre, contain, id_user) VALUES (?, ?, ?)`, titre, contain, id_user)
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

func selectUserById(db *sql.DB, id int) User {
	var user User
	db.QueryRow(`SELECT * FROM users WHERE id = ?`, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	return user
}

func selectPattern(db *sql.DB, option string, recherche string) *sql.Rows {
	pattern := "SELECT * FROM users WHERE " + option + " LIKE '%" + recherche + "%'"
	result, _ := db.Query(pattern)
	return result
}

func DataBase() {
	db := initDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	defer db.Close()
	insertIntoUsers(db, "Mathieu", "m.m@gmail.com", encoding_password("secret"))
	insertIntoUsers(db, "Thomas", "t.t@gmail.com", encoding_password("scret"))
	insertIntoUsers(db, "Lucas", "l.l@gmail.com", encoding_password("hello"))
	insertIntoUsers(db, "vanessa", "vanessa@gmail.com", encoding_password("world"))
	insertIntoUsers(db, "com", "vane@gmail.com", encoding_password("test"))
	insertIntoUsers(db, "com", "srvvane@gmail.com", encoding_password("test"))

	insertIntoTopic(db, "test", "j fait un test", 1)

	fmt.Println("\n")
	fmt.Println("dataBase")
	rows := selectAllFromTable(db, "topic")
	displayUserRows(rows)

	fmt.Println("\n")
	fmt.Println("par ID")
	user2 := selectUserById(db, 4)
	fmt.Println(user2)

	fmt.Println("\n")
	fmt.Println("par Name")
	user := selectPattern(db, "email", "com")
	displayUserRows(user)
}
