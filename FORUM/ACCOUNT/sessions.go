package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"unicode"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template
var db *sql.DB

// var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
// don't store key in source code, pass in via a environment variable to avoid accidentally commit it with code
// func NewCookieStore(keyPairs ...[]byte) *CookieStore
var store = sessions.NewCookieStore([]byte("CEDRICTROPBEAU"))

// func main() {
// 	tpl, _ = template.ParseGlob("templates/*.html")
// 	var err error
// 	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	http.HandleFunc("/login", loginHandler)
// 	http.HandleFunc("/loginauth", loginAuthHandler)
// 	http.HandleFunc("/logout", logoutHandler)
// 	http.HandleFunc("/register", registerHandler)
// 	http.HandleFunc("/registerauth", registerAuthHandler)
// 	http.HandleFunc("/about", aboutHandler)
// 	http.HandleFunc("/", indexHandler)

// 	http.ListenAndServe("localhost:8080", context.ClearHandler(http.DefaultServeMux))
// }

func LoginAuthHandler(w http.ResponseWriter, r *http.Request, BddID int, BddName string) {
	fmt.Println("--------------loginAuth--------------")
	session, _ := store.Get(r, "session")
	// CECI DOIT ETRE EGALE AU USER ID DE LA BDD C'EST LOGIQUE CONNARD
	session.Values["userName"] = BddName
	session.Values["userID"] = BddID
	// SAUVEGARDE DE LA SESSION
	session.Save(r, w)
	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------INDEX--------------")
	session, _ := store.Get(r, "session")
	// ICI ON VERIFIE QUE LE MEC EST BIEN LOG IN SINON ON RENVOIE SUR LA PAGE LOGIN C'EST PLUTOT SYMPA ("ok" EST UN BOOL)
	_, ok := session.Values["userID"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound) // http.StatusFound is 302
		return
	}
	tpl.ExecuteTemplate(w, "index.html", "Logged In")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------ABOUT--------------")
	session, _ := store.Get(r, "session")
	// ICI ON VERIFIE QUE LE MEC EST BIEN LOG IN SINON ON RENVOIE SUR LA PAGE LOGIN C'EST PLUTOT SYMPA ("ok" EST UN BOOL)
	_, ok := session.Values["userID"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound) // http.StatusFound is 302
		return
	}
	tpl.ExecuteTemplate(w, "about.html", "Logged In")
}

// A VOIR COMMENT UTILISER CETTE MERDE J'AI ENVIE DE CREVER
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------LOGOUT--------------")
	session, _ := store.Get(r, "session")
	fmt.Println("ICI ON SUPPRIME")
	fmt.Println(session.Values["userID"])
	//SUPPRIME L'ID ACTUEL PUIS REFRESH LA SESSION
	delete(session.Values, "userID")
	fmt.Println("check si c'est delet en dessous")
	fmt.Println(session.Values["userID"])
	session.Save(r, w)
	tpl.ExecuteTemplate(w, "login.html", "Logged Out")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------REGISTER--------------")
	tpl.ExecuteTemplate(w, "register.html", nil)
}

// BON LA C'EST TRES CHIANT MAIS EN GROS C'EST POUR CREER DES NOUVEAUX UTILISATEURS MAIS VAZY FLEMME ON FERA AUTRE CHOSE
func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	/*
		1. check username criteria
		2. check password criteria
		3. check if username is already exists in database
		4. create bcrypt hash from password
		5. insert username and password hash in database
		(email validation will be in another video)
	*/
	fmt.Println("*****registerAuthHandler running*****")
	r.ParseForm()
	username := r.FormValue("username")
	// check username for only alphaNumeric characters
	var nameAlphaNumeric = true
	for _, char := range username {
		// func IsLetter(r rune) bool, func IsNumber(r rune) bool
		// if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}
	// check username pswdLength
	var nameLength bool
	if 5 <= len(username) && len(username) <= 50 {
		nameLength = true
	}
	// check password criteria
	password := r.FormValue("password")
	fmt.Println("password:", password, "\npswdLength:", len(password))
	// variables that must pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
		// func IsLower(r rune) bool
		case unicode.IsLower(char):
			pswdLowercase = true
		// func IsUpper(r rune) bool
		case unicode.IsUpper(char):
			pswdUppercase = true
		// func IsNumber(r rune) bool
		case unicode.IsNumber(char):
			pswdNumber = true
		// func IsPunct(r rune) bool, func IsSymbol(r rune) bool
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		// func IsSpace(r rune) bool, type rune = int32
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	if 11 < len(password) && len(password) < 60 {
		pswdLength = true
	}
	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces || !nameAlphaNumeric || !nameLength {
		tpl.ExecuteTemplate(w, "register.html", "please check username and password criteria")
		return
	}
	// check if username already exists for availability
	stmt := "SELECT UserID FROM bcrypt WHERE username = ?"
	row := db.QueryRow(stmt, username)
	var uID string
	err := row.Scan(&uID)
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, err:", err)
		tpl.ExecuteTemplate(w, "register.html", "username already taken")
		return
	}
	// create hash from password
	var hash []byte
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
		return
	}
	fmt.Println("hash:", hash)
	fmt.Println("string(hash):", string(hash))
	// func (db *DB) Prepare(query string) (*Stmt, error)
	var insertStmt *sql.Stmt
	insertStmt, err = db.Prepare("INSERT INTO bcrypt (Username, Hash) VALUES (?, ?);")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
		return
	}
	defer insertStmt.Close()
	var result sql.Result
	//  func (s *Stmt) Exec(args ...interface{}) (Result, error)
	result, err = insertStmt.Exec(username, hash)
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("lastIns:", lastIns)
	fmt.Println("err:", err)
	if err != nil {
		fmt.Println("error inserting new user")
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
		return
	}
	fmt.Fprint(w, "congrats, your account has been successfully created")
}
