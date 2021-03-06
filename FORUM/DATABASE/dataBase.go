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
	Contain    string
	Nombre_rep int
	Id_user    int
	User_Name  string
	archive    int
}

type Reponse struct {
	Id          int
	TextContent string
	Id_user     int
	Id_topic    int
	User_name   string
}

type like struct {
	Id         int
	Id_user    int
	Id_reponse int
}

type Theme struct {
	Id    int
	Theme string
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
		user_name TEXT NOT NULL,
		archive INEGER NOT NULL,
		FOREIGN KEY(id_user) REFERENCES users(id)
		);

	CREATE TABLE IF NOT EXISTS reponses (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		id_topic INTEGER NOT NULL,
		id_user INTEGER NOT NULL,
		user_name TEXT NOT NULL,
		FOREIGN KEY(id_user) REFERENCES users(id),
		FOREIGN KEY(id_topic) REFERENCES topics(id)
	);
	CREATE TABLE IF NOT EXISTS likes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		id_reponse INTEGER,
		id_user TEXT NOT NULL,
		FOREIGN KEY(id_user) REFERENCES users(id)
		FOREIGN KEY(id_reponse) REFERENCES reponses(id)
	);
	CREATE TABLE IF NOT EXISTS themes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		theme TEXT NOT NULL UNIQUE
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
	dbUsers := InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	result, err := db.Exec(`INSERT INTO topics (titre, contain, nombre_rep, id_user, user_name, archive) VALUES (?, ?, ?, ?, ?, 0)`, titre, contain, nombre_rep, id_user, SelectUserById(dbUsers, "users", id_user).Name)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return -1, nil
	}
	return result.LastInsertId()
}

func InsertIntoReponse(db *sql.DB, content string, id_topic int, id_user int) (int64, error) {
	dbUsers := InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	result, err := db.Exec(`INSERT INTO reponses (content, id_topic, id_user, user_name) VALUES (?, ?, ?, ?)`, content, id_topic, id_user, SelectUserById(dbUsers, "users", id_user).Name)
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

func InsertIntoTheme(db *sql.DB, theme string) (int64, error) {
	result, err := db.Exec(`INSERT INTO themes (theme) VALUES (?)`, theme)
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
		err := rows.Scan(&topic.Id, &topic.Titre, &topic.Contain, &topic.Nombre_rep, &topic.Id_user, &topic.User_Name, &topic.archive)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(topic)
	}
}

func DisplayReponsesRows(rows *sql.Rows) []Reponse {
	var reponses []Reponse
	for rows.Next() {
		var reponse Reponse
		err := rows.Scan(&reponse.Id, &reponse.Id_topic, &reponse.Id_user, &reponse.TextContent, &reponse.User_name)
		if err != nil {
			log.Fatal(err)
		}
		reponses = append(reponses, reponse)
	}
	return reponses
}

func SelectThemeById(db *sql.DB, id int) (Theme, error) {
	var theme Theme
	if err := db.QueryRow("SELECT * FROM themes WHERE id = ?", id).Scan(&theme.Id, &theme.Theme); err != nil {
		return theme, err
	}

	return theme, nil
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
	db.Exec("DELETE FROM "+table+" WHERE id = ?;", id)
}

//----------------------------------
func UpDate(db *sql.DB, table string, optionEdited string, new string, id int) {
	if optionEdited != "id" {
		db.Exec("UPDATE "+table+" SET "+optionEdited+" = "+"'"+new+"'"+" WHERE id = ?;", id)
	}
}

//---------------------------------
func Count(db *sql.DB, table string, id_reponse int) int {
	var n int
	result := db.QueryRow("SELECT COUNT(*) FROM "+table+" WHERE id_reponse = ?", id_reponse)
	result.Scan(&n)
	return n
}

func SelectArchiveFromTopic(db *sql.DB, archive int) ([]Topic, error) {
	pattern, err := db.Query("SELECT * FROM topics WHERE archive =?", archive)
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return nil, err
	}

	var topics []Topic
	for pattern.Next() {
		var topic Topic
		if err := pattern.Scan(&topic.Id, &topic.Titre, &topic.Contain, &topic.Nombre_rep, &topic.Id_user, &topic.User_Name, &topic.archive); err != nil {
			return topics, err
		}
		topics = append(topics, topic)
	}
	if err = pattern.Err(); err != nil {
		return topics, err
	}
	return topics, nil
}

func Archive(db *sql.DB) {
	db.Exec("UPDATE topics SET archive = archive + 1")
	db.Exec("UPDATE reponses SET archive = archive + 1")
}

func ArchiveSpecificTopic(db *sql.DB, id int) {
	db.Exec("UPDATE topics SET archive = archive + 1 WHERE id = ?", id)
}

func SelectReponseFromTopic(db *sql.DB) ([]Reponse, error) {
	pattern, err := db.Query("SELECT * FROM reponses")
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return nil, err
	}

	var reponses []Reponse
	for pattern.Next() {
		var reponse Reponse
		if err := pattern.Scan(&reponse.Id, &reponse.TextContent, &reponse.Id_topic, &reponse.Id_user, &reponse.User_name); err != nil {
			return reponses, err
		}
		reponses = append(reponses, reponse)
	}
	if err = pattern.Err(); err != nil {
		return reponses, err
	}
	return reponses, nil
}

func SelectThemes(db *sql.DB) ([]Theme, error) {
	pattern, err := db.Query("SELECT * FROM themes")
	if err != nil {
		log.Printf("ERR: %s\n", err)
		return nil, err
	}

	var themes []Theme
	for pattern.Next() {
		var theme Theme
		if err := pattern.Scan(&theme.Id, &theme.Theme); err != nil {
			return themes, err
		}
		themes = append(themes, theme)
	}
	if err = pattern.Err(); err != nil {
		return themes, err
	}
	return themes, nil
}

func Login(db *sql.DB, email string, password string) User {
	fmt.Println("---------------Login Function---------------")
	var user User
	Epassword := Encoding_password(password)
	db.QueryRow("SELECT * FROM users WHERE email = ? AND password = ?", email, Epassword).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if user == (User{}) {
		fmt.Println("Can't Load User")

	}
	return user
}

func DataBase() {
	db := InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	defer db.Close()

	// InsertIntoUsers(db, "Mathieu", "m.m@gmail.com", Encoding_password("secret"))
	// InsertIntoUsers(db, "Thomas", "t.t@gmail.com", Encoding_password("scret"))
	// InsertIntoUsers(db, "Lucas", "l.l@gmail.com", Encoding_password("hello"))
	// InsertIntoUsers(db, "vanessa", "vanessa@gmail.com", Encoding_password("world"))
	// InsertIntoUsers(db, "com", "vane@gmail.com", Encoding_password("test"))
	// InsertIntoUsers(db, "com", "srvvane@gmail.com", Encoding_password("test"))

	// InsertIntoTopic(db, "test", "j fait un test", 0, 4)

	// InsertIntoReponse(db, "je rep", 5, 1)

	// fmt.Println("USER : ", Login(db, "random.mail@mail.mail", "12345678"))

	// InsertIntoLike(db, 2, 2)
	// InsertIntoLike(db, 2, 2)
	// InsertIntoLike(db, 2, 2)

	fmt.Println(Count(db, "likes", 1))

	// topic := SelectArchiveFromTopic(db, 0)
	// DisplayTopicRows(topic)
	fmt.Println(SelectArchiveFromTopic(db, 1))
	// rows := SelectAllFromTable(db, "users")
	// DisplayUserRows(rows)

	// fmt.Println("\n")
	// fmt.Println("par ID")
	// user2 := SelectUserById(db, "topics", 5)
	// fmt.Println(user2)

	fmt.Println("\n")
	fmt.Println("dataBase")
	rows := SelectAllFromTable(db, "topics")
	DisplayTopicRows(rows)

	fmt.Println("\n")
	fmt.Println("par ID")
	user2 := SelectUserById(db, "name", 4)
	fmt.Println(user2)

	// UpDate(db, "users", "email", "nazi.super.MILF@gmail.con", 1)

	// fmt.Println("\n")
	// fmt.Println("topics")
	// InsertIntoTheme(db, "bikini")
	// InsertIntoTheme(db, "pilon")
	// InsertIntoTheme(db, "portugai")
	fmt.Println(SelectThemes(db))

}
