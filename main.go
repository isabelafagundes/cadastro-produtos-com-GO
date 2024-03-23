package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"text/template"
	"trabalhando-golang/routes"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	routes.CarregarRotas()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
