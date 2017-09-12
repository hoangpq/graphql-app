package orm

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func GetProducts() *sql.Rows {
	db, err := sql.Open("postgres", "postgres://odoo:odoo@localhost:5432/odoo10?sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		rows, err := db.Query("SELECT id, name, list_price as price FROM product_template")
		if err != nil {
			panic(err)
		}
		return rows
	}
	return nil
}
