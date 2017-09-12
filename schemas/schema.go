package schemas

import (
	"go-grapgql-practice/orm"
	"github.com/graphql-go/graphql"
	"go-grapgql-practice/models"
	"log"
)

func GetProductList() []models.Product {
	productsList := []models.Product{}
	// query product
	rows := orm.GetProducts()
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var price float32
		err := rows.Scan(&id, &name, &price)
		if err != nil {
			panic(err)
		}
		productsList = append(productsList, models.Product{Id: id, Name: name, Price: price})
	}
	return productsList
}

func GetSchema() (graphql.Schema, error) {
	productType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Product",
		Description: "The Product Type",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.ID,
				Description: "The id of the product",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if product, ok := p.Source.(models.Product); ok {
						return product.Id, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the product",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if product, ok := p.Source.(models.Product); ok {
						return product.Name, nil
					}
					return nil, nil
				},
			},
			"price": &graphql.Field{
				Type:        graphql.Float,
				Description: "The price of the product",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if product, ok := p.Source.(models.Product); ok {
						return product.Price, nil
					}
					return nil, nil
				},
			},
		},
	})

	fields := graphql.Fields{
		"products": &graphql.Field{
			Type: graphql.NewList(productType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return GetProductList(), nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schemas, error: %v", err)
	}
	return schema, err
}
