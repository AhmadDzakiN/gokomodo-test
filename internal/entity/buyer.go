package entity

import "time"

type Buyer struct {
	ID              string    `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Email           string    `gorm:"column:email;not null"`
	Password        string    `gorm:"column:password;not null"`
	Name            string    `gorm:"column:name;not null"`
	ShippingAddress string    `gorm:"column:shipping_address;not null"`
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt       time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null"`
}

func (b *Buyer) TableName() string {
	return "buyers"
}
