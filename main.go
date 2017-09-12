package main

import (
	"github.com/graphql-go/handler"
	"net/http"
	"html/template"
	"fmt"
	"go-grapgql-practice/schemas"
)

func main() {
	// Schema
	schema, err := schemas.GetSchema()
	if err != nil {
		panic(err)
	}
	// define GraphQL schemas using relay library helpers
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	// serve HTTP
	fs := http.FileServer(http.Dir("client/dist"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/graphql", h)
	http.HandleFunc("/", serveTemplate)
	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("client/index.html")
	t.Execute(w, nil)
}
