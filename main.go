package main

import (
	"fmt"
	"go-grapgql-practice/configs"
	"go-grapgql-practice/schemas"
	"html/template"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/graphql-go/handler"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	data, _ := configs.GetDatabaseConfig()
	fmt.Println(data)
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
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/dist/", http.StripPrefix("/dist/", fs))
	http.Handle("/api/graphql", h)
	http.HandleFunc("/", serveVueTemplate)
	http.HandleFunc("/graphql", serveTemplate)
	http.HandleFunc("/ws", wshandler)
	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("dist/graphiql.html")
	t.Execute(w, nil)
}

func serveVueTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("dist/index.html")
	t.Execute(w, nil)
}
