package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
	"html/template"
)

func main() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	// define GraphQL schema using relay library helpers
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	// serve HTTP
	fs := http.FileServer(http.Dir("client/dist"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/graphql", h)
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("client/index.html")
	t.Execute(w, nil)
}
