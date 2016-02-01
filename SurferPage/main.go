package main

import (
	"net/http"
	"html/template"
	"log"
)

func makesite(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
}
func main() {

	http.HandleFunc("/", makesite)
	http.Handle("/pic/", http.StripPrefix("/pic", http.FileServer(http.Dir("./pic"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	http.ListenAndServe(":8040", nil)
}