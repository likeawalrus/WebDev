package main

import (
	"fmt"
	"html/template"
	"net/http"
	"io/ioutil"

)

func uploadTextFile(res http.ResponseWriter, req *http.Request) {
	mytemp, err := template.ParseFiles("index.gohtml")
	if req.Method == "POST" {
		theFile, _, err := req.FormFile("name")
		defer theFile.Close()
		contents, err := ioutil.ReadAll(inFile)
		err = mytemp.Execute(res, string(contents))
	}
}


func main() {
	http.HandleFunc("/", fileUp)
	http.ListenAndServe(":8080", nil)
}
