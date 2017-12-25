package models

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type ProductUOM struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UOMMessage struct {
	Uom       ProductUOM `json:"uom"`
	ProductId int        `json:"product_id"`
}
