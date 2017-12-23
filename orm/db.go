package orm

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-grapgql-practice/configs"
	"fmt"
	"go-grapgql-practice/models"
)

func getConnection() (*sql.DB, error) {
	config, _ := configs.GetDatabaseConfig()
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Dbname,
	)
	return sql.Open("postgres", connString)
}

func GetProducts() *sql.Rows {
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT id, name, list_price AS price FROM product_template")
	if err != nil {
		panic(err)
	}
	return rows
}

func _GetUOMByProductID(productId int) models.ProductUOM {
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	c := make(chan models.ProductUOM)
	go func() {
		query := fmt.Sprintf("SELECT pu.id as id, pu.name as name "+
			"FROM product_template tmpl INNER JOIN product_uom pu ON tmpl.uom_id = pu.id "+
			"WHERE tmpl.id = %d", productId)
		row := db.QueryRow(query)
		var id int
		var name string
		if row.Scan(&id, &name) != nil {
			c <- models.ProductUOM{}
		}
		c <- models.ProductUOM{id, name}
	}()
	return <-c
}

func GetUOMByProductID(productId int) models.ProductUOM {
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	query := fmt.Sprintf("SELECT pu.id as id, pu.name as name "+
		"FROM product_template tmpl INNER JOIN product_uom pu ON tmpl.uom_id = pu.id "+
		"WHERE tmpl.id = %d", productId)
	row := db.QueryRow(query)
	var id int
	var name string
	if row.Scan(&id, &name) != nil {
		return models.ProductUOM{}
	}
	return models.ProductUOM{id, name}
}
