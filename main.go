package main

import (
	"fmt"
	"html/template"
	"net/http"
	"go-grapgql-practice/schemas"
	"github.com/gorilla/websocket"
	"github.com/graphql-go/handler"
	"go-grapgql-practice/orm"
	"go-grapgql-practice/models"
	"time"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func GetUOM() {
	start := time.Now()
	ids := []int{1, 2, 3}
	var uoms []models.ProductUOM
	c := make(chan models.ProductUOM)
	for _, id := range ids {
		go func(id int) {
			uom := orm.GetUomById(id)
			if uom != nil {
				c <- uom.(models.ProductUOM)
			}
		}(id)
	}
	for len(uoms) < len(ids) {
		select {
		case uom := <-c:
			uoms = append(uoms, uom)
		}
	}
	end := time.Since(start)
	fmt.Printf("%s\n", end)
	fmt.Println(uoms)
}

func main() {
	GetUOM()
	// Schema
	schema, err := schemas.GetSchema()
	if err != nil {
		panic(err)
	}
	// define GraphQL schemas using relay library helpers
	graphqlHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	// serve HTTP
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/dist/", http.StripPrefix("/dist/", fs))
	http.Handle("/api/graphql", graphqlHandler)
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
