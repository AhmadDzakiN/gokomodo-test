package entity

type Seller struct {
	ID                 string `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	BuyerID            string `gorm:"column:buyer_id;not null"`
	SellerID           string `gorm:"column:seller_id;not null"`
	SourceAddress      string `gorm:"column:source_address;not null"`
	DestinationAddress string `gorm:"column:destination_address;not null"`
	Items              string `gorm:"column:items;not null"`
	Quantity           int    `gorm:"column:quantity;not null"`
	Price              int64  `gorm:"column:price;not null"`
	TotalPrice         int64  `gorm:"column:total_price;not null"`
	Status             string `gorm:"column:status;not null"`
}

func (s *Seller) TableName() string {
	return "sellers"
}
