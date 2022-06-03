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

type Topic struct {
	Id         int
	Titre      string
	contain    string
	nombre_rep int
	id_user    int
}

type Reponse struct {
	Id          int
	TextContent string
	Id_user     int
	Id_topic    int
}

func InitDatabase(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := ` 
	PRAGMA foreign_keys = ON;
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS topics (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		titre TEXT NOT NULL,
		contain TEXT NOT NULL,
		nombre_rep INTEGER,
		id_user INTEGER NOT NULL,
		FOREIGN KEY(id_user) REFERENCES users(id)
		);

	CREATE TABLE IF NOT EXISTS reponse (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		id_user TEXT NOT NULL,
		id_topic INTEGER,
		FOREIGN KEY(id_user) REFERENCES users(id)
		FOREIGN KEY(id_topic) REFERENCES topics(id)

	)
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return db
}

func SelectAllFromTable(db *sql.DB, table string) *sql.Rows {
	query := "SELECT * FROM " + table
	result, _ := db.Query(query)
	return result
}

func InsertIntoUsers(db *sql.DB, name string, email string, password string) (int64, error) {
	result, err := db.Exec(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`, name, email, password)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return -1, nil
	}
	return result.LastInsertId()
}

func InsertIntoTopic(db *sql.DB, titre string, contain string, nombre_rep int, id_user int) (int64, error) {
	result, err := db.Exec(`INSERT INTO topics (titre, contain, nombre_rep, id_user) VALUES (?, ?, ?, ?)`, titre, contain, nombre_rep, id_user)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return -1, nil
	}
	return result.LastInsertId()
}

func InsertIntoReponse(db *sql.DB, content string, id_user int, id_topic int) (int64, error) {
	result, err := db.Exec(`INSERT INTO reponse (content, id_user, id_topic) VALUES (?, ?, ?)`, content, id_user, id_topic)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return -1, nil
	}
	return result.LastInsertId()
}

func Encoding_password(a string) string {
	b := base32.StdEncoding.EncodeToString([]byte(a))
	return string(b)
}

func DisplayUserRows(rows *sql.Rows) {
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
	}
}

func DisplayTopicRows(rows *sql.Rows) {
	for rows.Next() {
		var topic Topic
		err := rows.Scan(&topic.Id, &topic.Titre, &topic.contain, &topic.nombre_rep, &topic.id_user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(topic)
	}
}

func SelectUserById(db *sql.DB, option string, id int) User {
	var user User
	db.QueryRow("SELECT * FROM "+option+" WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	return user
}

func SelectPattern(db *sql.DB, option string, recherche string) *sql.Rows {
	pattern := "SELECT * FROM users WHERE " + option + " LIKE '%" + recherche + "%'"
	result, _ := db.Query(pattern)
	return result
}

func DeleteUsersById(db *sql.DB, table string, id int) {
	db.Exec("DELETE FROM "+table+" WHERE id = ?", id)
}

func DataBase() {
	db := InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	defer db.Close()
	InsertIntoUsers(db, "Mathieu", "m.m@gmail.com", Encoding_password("secret"))
	InsertIntoUsers(db, "Thomas", "t.t@gmail.com", Encoding_password("scret"))
	InsertIntoUsers(db, "Lucas", "l.l@gmail.com", Encoding_password("hello"))
	InsertIntoUsers(db, "vanessa", "vanessa@gmail.com", Encoding_password("world"))
	InsertIntoUsers(db, "com", "vane@gmail.com", Encoding_password("test"))
	InsertIntoUsers(db, "com", "srvvane@gmail.com", Encoding_password("test"))

	// fmt.Println("\n")
	// fmt.Println("dataBase")
	// rows := SelectAllFromTable(db, "users")
	// DisplayUserRows(rows)

	// fmt.Println("\n")
	// fmt.Println("par ID")
	// user2 := SelectUserById(db, "topics", 5)
	// fmt.Println(user2)

	// fmt.Println("\n")
	// fmt.Println("recherche")
	// user := SelectPattern(db, "email", "com")
	// DisplayUserRows(user)

	// DeleteUsersById(db, "topics", 1)

	// fmt.Println("\n")
	// fmt.Println("topics")
	// InsertIntoTopic(db, "test", "j fait un test", 0, 3)
	// topic := SelectAllFromTable(db, "topics")
	// DisplayTopicRows(topic)
}