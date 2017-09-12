package orm

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-grapgql-practice/configs"
	"fmt"
)

func GetProducts() *sql.Rows {
	config, _ := configs.GetDatabaseConfig()
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Dbname,
	)

	db, err := sql.Open("postgres", connString)
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
