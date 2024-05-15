package entity

type Buyer struct {
	ID              string `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Email           string `gorm:"column:email;not null"`
	Password        string `gorm:"column:password;not null"`
	Name            string `gorm:"column:name;not null"`
	ShippingAddress string `gorm:"column:shipping_address;not null"`
}

func (b *Buyer) TableName() string {
	return "buyers"
}
