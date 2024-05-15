package entity

type Product struct {
	ID          string `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string `gorm:"column:name;not null"`
	Description string `gorm:"column:description;not null"`
	Price       int64  `gorm:"column:price;not null"`
	SellerID    string `gorm:"column:seller_id;not null"`
}

type GetProductListParams struct {
	LastValue uint64 `json:"-"`
	NextToken string `json:"-"`
	Limit     int    `json:"-"`
	Query     string `json:"-"`
	SellerID  string `json:"-"`
}

func (p *Product) TableName() string {
	return "products"
}
