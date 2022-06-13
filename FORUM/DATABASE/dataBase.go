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
	archive    int
}

type Reponse struct {
	Id          int
	TextContent string
	Id_user     int
	Id_topic    int
}

type like struct {
	Id         int
	Id_user    int
	Id_reponse int
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
		archive INEGER NOT NULL,
		FOREIGN KEY(id_user) REFERENCES users(id)
		);

	CREATE TABLE IF NOT EXISTS reponses (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		id_topic INTEGER NOT NULL,
		id_user INTEGER NOT NULL,
		FOREIGN KEY(id_user) REFERENCES users(id)
		FOREIGN KEY(id_topic) REFERENCES topics(id)
	);
	CREATE TABLE IF NOT EXISTS likes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		id_reponse INTEGER,
		id_user TEXT NOT NULL,
		FOREIGN KEY(id_user) REFERENCES users(id)
		FOREIGN KEY(id_reponse) REFERENCES reponses(id)
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

//inserce user
func InsertIntoUsers(db *sql.DB, name string, email string, password string) (int64, error) {
	result, err := db.Exec(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`, name, email, password)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return -1, nil
	}
	return result.LastInsertId()
}

func Login(db *sql.DB, email string, password string) User {
	var user User
	Epassword := Encoding_password(password)
	db.QueryRow("SELECT * FROM users WHERE email = ? AND password = ?", email, Epassword).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	return user
}

// _-_-_-_-_-_-_-_-_-_-_-_-

func InsertIntoTopic(db *sql.DB, titre string, contain string, nombre_rep int, id_user int, archive int) (int64, error) {
	result, err := db.Exec(`INSERT INTO topics (titre, contain, nombre_rep, id_user, archive) VALUES (?, ?, ?, ?, ?)`, titre, contain, nombre_rep, id_user, archive)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return -1, nil
	}
	return result.LastInsertId()
}

func InsertIntoReponse(db *sql.DB, content string, id_topic int, id_user int) (int64, error) {
	result, err := db.Exec(`INSERT INTO reponses (content, id_topic, id_user) VALUES (?, ?, ?)`, content, id_topic, id_user)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return -1, nil
	}
	return result.LastInsertId()
}

func InsertIntoLike(db *sql.DB, id_reponse int, id_user int) (int64, error) {
	result, err := db.Exec(`INSERT INTO likes (id_reponse, id_user) VALUES (?, ?)`, id_reponse, id_user)
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
		err := rows.Scan(&topic.Id, &topic.Titre, &topic.contain, &topic.nombre_rep, &topic.id_user, &topic.archive)
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

func UpDate(db *sql.DB, option string, jacklyne string, new string, id int) {
	if jacklyne != "id" {
		db.Exec("UPDATE "+option+" SET "+jacklyne+" = "+"'"+new+"'"+" WHERE id = ?;", id)
	}
}

func Count(db *sql.DB, table string, id_reponse int) int {
	var n int
	result := db.QueryRow("SELECT COUNT(*) FROM "+table+" WHERE id_reponse = ?", id_reponse)
	result.Scan(&n)
	return n
}

// func SelectArchiveFromTopic(db *sql.DB, archive int) *sql.Rows {
// 	var topic Topic
// 	pattern := db.QueryRow("SELECT * FROM topics WHERE archive =?", archive)
// 	pattern.Scan(&topic.Id, &topic.Titre, &topic.contain, &topic.nombre_rep, &topic.id_user, &topic.archive)
// 	result, _ := db.Query(pattern)
// 	return result
// }

func SelectArchiveFromTopic(db *sql.DB, archive int) ([]Topic, error) {
	pattern, err := db.Query("SELECT * FROM topics WHERE archive =?", archive)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return nil, err
	}

	var topics []Topic
	for pattern.Next() {
		var topic Topic
		if err := pattern.Scan(&topic.Id, &topic.Titre, &topic.contain, &topic.nombre_rep, &topic.id_user, &topic.archive); err != nil {
			return topics, err
		}
		topics = append(topics, topic)
	}
	if err = pattern.Err(); err != nil {
		return topics, err
	}
	return topics, nil
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

	// InsertIntoTopic(db, "test", "j fait un test", 0, 4, 1)

	// InsertIntoReponse(db, "je rep", 1, 6)

	// InsertIntoLike(db, 2, 2)
	// InsertIntoLike(db, 2, 2)
	// InsertIntoLike(db, 2, 2)

	// fmt.Println(Count(db, "likes", 1))

	// topic := SelectArchiveFromTopic(db, 0)
	// DisplayTopicRows(topic)
	// rows := SelectAllFromTable(db, "users")
	// DisplayUserRows(rows)

	// fmt.Println("\n")
	// fmt.Println("par ID")
	// user2 := SelectUserById(db, "topics", 5)
	// fmt.Println(user2)

	// fmt.Println("\n")
	// fmt.Println("dataBase")
	// rows := SelectAllFromTable(db, "topics")
	// DisplayTopicRows(rows)

	// fmt.Println("\n")
	fmt.Println("par ID")
	user2 := SelectUserById(db, "users", 4)
	fmt.Println(user2)

	// UpDate(db, "users", "email", "nazi.super.MILF@gmail.con", 1)

	fmt.Println("\n")
	userTest := Login(db, "vanessa@gmail.com", "worl")
	fmt.Println(userTest)
}
