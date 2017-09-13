package schemas

import (
	"go-grapgql-practice/orm"
	"github.com/graphql-go/graphql"
	"go-grapgql-practice/models"
	"log"
)

var (
	humanType *graphql.Object
	droidType *graphql.Object
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

func GenCharacter() []Character {
	humans := []Character{
		{Name: "Jedi", Starship: "TIE Advanced x1"},
	}
	droids := []Character{
		{Name: "R2-D2", PrimaryFunction: "Astromech"},
	}
	return append(humans, droids...)
}

type Character struct {
	Name            string `json:"name"`
	Starship        string `json:"starship"`
	PrimaryFunction string `json:"primaryFunction"`
}

func GetSchema() (graphql.Schema, error) {

	characterInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Character",
		Description: "A Character in the Star Wars Trilogy",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the character",
			},
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			if character, ok := p.Value.(Character); ok {
				if character.Starship != "" {
					return humanType
				}
				return droidType
			}
			return nil
		},
	})

	humanType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Human",
		Description: "A humanoid creature in the Star Wars universe",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the human",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if character, ok := p.Source.(Character); ok {
						return character.Name, nil
					}
					return nil, nil
				},
			},
			"starship": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The starship of the human",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if character, ok := p.Source.(Character); ok {
						return character.Starship, nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			characterInterface,
		},
	})

	droidType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Droid",
		Description: "A mechanical creature in the Star Wars universe",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the droid",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if character, ok := p.Source.(Character); ok {
						return character.Name, nil
					}
					return nil, nil
				},
			},
			"primaryFunction": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The function of the droid",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if character, ok := p.Source.(Character); ok {
						return character.PrimaryFunction, nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			characterInterface,
		},
	})

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
		"characters": &graphql.Field{
			Type: graphql.NewList(characterInterface),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return GenCharacter(), nil
			},
		},
		"droid": &graphql.Field{
			Type: droidType,
		},
		"human": &graphql.Field{
			Type: humanType,
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
