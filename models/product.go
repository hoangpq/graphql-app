package models

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `gorm:"field:list_price",json:"price"`
}

func (Product) TableName() (string) {
	return "product_template"
}

type ProductUOM struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (ProductUOM) TableName() (string) {
	return "product_uom"
}

type UOMMessage struct {
	Uom       ProductUOM `json:"uom"`
	ProductId int        `json:"product_id"`
}
