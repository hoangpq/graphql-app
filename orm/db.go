package orm

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-grapgql-practice/configs"
	"fmt"
	"go-grapgql-practice/models"
)

func GetConnection() (*sql.DB, error) {
	config, _ := configs.GetDatabaseConfig()
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Dbname,
	)
	return sql.Open("postgres", connString)
}

func GetProducts() []models.Product {
	db, err := GetConnection()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	var products []models.Product
	rows, err := db.Query("SELECT id, name, list_price AS price FROM product_template")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var (
		id    int
		name  string
		price float32
	)
	for rows.Next() {
		rows.Scan(&id, &name, &price)
		products = append(products, models.Product{id, name, price})
	}
	return products
}

func GetUomById(uomId int) interface{} {
	db, err := GetConnection()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	var (
		id   int
		name string
	)
	err = db.QueryRow("SELECT id, name FROM product_uom WHERE id = $1", uomId).Scan(&id, &name)
	if err == sql.ErrNoRows || err != nil {
		return nil
	}
	return models.ProductUOM{id, name}
}

func GetProductById(productId int) interface{} {
	db, err := GetConnection()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	var (
		id    int
		name  string
		price float32
	)
	err = db.QueryRow(`
		  SELECT id, name, list_price AS price
		  FROM product_template
		  WHERE id = $1`, productId).Scan(&id, &name, &price)
	if err == sql.ErrNoRows || err != nil {
		return nil
	}
	return models.Product{id, name, price}
}

func GetUOMByProductID(productId int) interface{} {
	db, err := GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var (
		id   int
		name string
	)
	// db.ExecContext(context.Background(), "SELECT pg_sleep(5)")
	err = db.QueryRow(`
		SELECT pu.id AS id, pu.name AS name
		FROM product_template tmpl
		LEFT JOIN product_uom pu ON tmpl.uom_id = pu.id
		WHERE tmpl.id = $1;`, productId).Scan(&id, &name)
	if err == sql.ErrNoRows || err != nil {
		return nil
	}
	return models.ProductUOM{id, name}
}
