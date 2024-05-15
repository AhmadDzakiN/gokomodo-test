package entity

import "time"

type Product struct {
	ID          uint64    `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name;not null"`
	Description string    `gorm:"column:description;not null"`
	Price       int64     `gorm:"column:price;not null"`
	SellerID    string    `gorm:"column:seller_id;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null"`
}

func (p *Product) TableName() string {
	return "products"
}
