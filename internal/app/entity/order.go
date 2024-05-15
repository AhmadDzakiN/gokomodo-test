package entity

import "time"

type Order struct {
	ID                 uint64    `gorm:"column:id;primaryKey"`
	BuyerID            string    `gorm:"column:buyer_id;not null"`
	SellerID           string    `gorm:"column:seller_id;not null"`
	SourceAddress      string    `gorm:"column:source_address;not null"`
	DestinationAddress string    `gorm:"column:destination_address;not null"`
	Items              uint64    `gorm:"column:items;not null"`
	Quantity           uint      `gorm:"column:quantity;not null"`
	Price              uint64    `gorm:"column:price;not null"`
	TotalPrice         uint64    `gorm:"column:total_price;not null"`
	Status             string    `gorm:"column:status;not null"`
	CreatedAt          time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt          time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null"`
}

func (o *Order) TableName() string {
	return "orders"
}
