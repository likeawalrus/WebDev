package main

import (
	"html/template"
	"log"
	"os"
	"strings"
  "net/http"
	"encoding/base64"
	"encoding/json"
)

type Page struct {
	Title string
	Body  string
}

type User struct {
	Name string
	Age string
}

func main() {
	var err error

	tpl := template.New("tpl.gohtml")
	tpl = tpl.Funcs(template.FuncMap{
		"uppercase": func(str string) string {
			return strings.ToUpper(str)
		},
	})
	tpl, err = tpl.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	name := req.FormValue("name")
	age := req.FormValue("age")
	user := User{
		Name:     name,
		Age:      age,
	}
	jsonstuff, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("error: ", err)
	}
	base64stuff := base64.URLEncoding.EncodeToString(jsonstuff)
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			// Secure: true,
			Name:     "session-fino",
			Value:    id.String()+name+age+base64stuff,
			HttpOnly: true,
		}
	}
	err = tpl.Execute(os.Stdout, Page{
		Title: "My Title 2",
		Body:  `hello world <script>alert("Yow!");</script>`,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
