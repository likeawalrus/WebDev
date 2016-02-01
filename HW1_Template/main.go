package main

import (
	"log"
	"os"
	"text/template"
)

type thing struct {
	Name string
	IsHuman  bool
}

func main() {
	test1 := thing{
		Name: "Dog",
		IsHuman: false,
	}

	tpl, err := template.ParseFiles("template.gohtml.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, test1)
	if err != nil {
		log.Fatalln(err)
	}
}