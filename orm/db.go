package orm

import (
	"fmt"
	_ "github.com/lib/pq"
	"go-grapgql-practice/configs"
	"go-grapgql-practice/models"
	"github.com/jinzhu/gorm"
)

func GetConnection() (*gorm.DB) {
	config, _ := configs.GetDatabaseConfig()
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Dbname,
	)
	db, err := gorm.Open("postgres", connString)
	db.LogMode(true)
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func GetProducts() []models.Product {
	db := GetConnection()
	defer db.Close()
	var products []models.Product
	db.Debug().Select([]string{"id", "name", "list_price"}).Find(&products)
	return products
}

func GetUomById(uomId int) interface{} {
	db := GetConnection()
	defer db.Close()
	var uom models.ProductUOM
	db.Debug().Where("id = ?", uomId).Select([]string{"id", "name"}).First(&uom)
	return uom
}

func GetProductById(productId int) interface{} {
	db := GetConnection()
	defer db.Close()
	var product models.Product
	db.Debug().Where("id = ?", productId).Select([]string{"id", "name", "list_price"}).First(&product)
	return product
}

func GetUOMByProductID(productId int) interface{} {
	db := GetConnection()
	defer db.Close()
	uom := models.ProductUOM{}
	db.Debug().
		Table("product_template tmpl").
		Select("uom.id, uom.name").
		Joins("left join product_uom uom on tmpl.uom_id = uom.id").
		Where("tmpl.id = ?", productId).
		Find(&uom)
	return uom
}

func GetProductCount() (int) {
	db := GetConnection()
	defer db.Close()
	var count int
	db.Debug().
		Model(&models.Product{}).
		Count(&count)
	return count
}
