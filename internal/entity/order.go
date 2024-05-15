package entity

type Order struct {
	ID            string `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Email         string `gorm:"column:email;not null"`
	Password      string `gorm:"column:password;not null"`
	Name          string `gorm:"column:name;not null"`
	PickupAddress string `gorm:"column:pickup_address;not null"`
}

func (o *Order) TableName() string {
	return "orders"
}
