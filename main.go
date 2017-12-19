package main

import (
	"github.com/graphql-go/handler"
	"github.com/gorilla/websocket"
	"net/http"
	"html/template"
	"fmt"
	"go-grapgql-practice/schemas"
	"go-grapgql-practice/configs"
	"encoding/json"
	"time"
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
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/graphql", h)
	http.HandleFunc("/json", handleJSONRequest)
	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/ws", wshandler)
	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func handleJSONRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	data := make(map[string]string)
	data["text"] = "Hello, World"
	jsonStr, _ := json.Marshal(data)
	time.Sleep(5 * time.Second)
	w.Write(jsonStr)
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
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}
