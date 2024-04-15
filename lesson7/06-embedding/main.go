package main

import (
	"embed"
	"log"
	"os"
	"text/template"
)

type Person struct {
	FirstName, LastName string
	Age                 int
}

type Class struct {
	Grade    int
	Teacher  Person
	Students []Person
}

//go:embed template.tpl
var tplData embed.FS

func main() {
	tpl, err := template.New("template.tpl").Funcs(template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}).ParseFS(tplData, "*.tpl")
	if err != nil {
		log.Fatalf("parse template file error: %s", err)
	}

	data := loadData()

	if err = tpl.ExecuteTemplate(os.Stdout, "template.tpl", data); err != nil {
		log.Fatalf("execute template error: %s", err)
	}
}
