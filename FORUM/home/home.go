package forum

import (
	"fmt"
	session "forum/FORUM/ACCOUNT"
	database "forum/FORUM/DATABASE"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type Name_Theme struct {
	Theme string
	Name  string
}

var Name_theme Name_Theme

func Mainpage(w http.ResponseWriter, r *http.Request) {
	session := session.Store
	sessions, _ := session.Get(r, "session")
	if sessions.Values["userName"] != nil {
		Name_theme.Name = sessions.Values["userName"].(string)
	}
	var templates = template.Must(template.ParseFiles("././static/home/home.html", "././static/home/header.html"))
	templates.Execute(w, Name_theme)
}

func GetWeeklyTheme() string {
	db := database.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	DbData, err := database.SelectThemes(db)

	if err != nil {
		fmt.Println("Error loeading Themes table : ", err)
	}

	rand.Seed(time.Now().UnixNano())
	randomThemeID := rand.Intn(len(DbData)) + 1

	result, err2 := database.SelectThemeById(db, randomThemeID)

	if err2 != nil {
		fmt.Println("Error searching Theme : ", err)
	}
	fmt.Println("Nouveaux theme !!! : ", result.Theme)
	fmt.Println("Archivage en cours...")
	database.Archive(db)
	return result.Theme
}
