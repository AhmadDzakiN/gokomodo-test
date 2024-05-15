package repository

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/app/entity"
	"gorm.io/gorm"
)

type BuyerRepository struct {
	db *gorm.DB
}

type IBuyerRepository interface {
	Login(ctx *gin.Context, buyer entity.Buyer) (err error)
	GetByID(ctx *gin.Context, id string) (buyer entity.Buyer, err error)
}

func NewBuyerRepository(db *gorm.DB) *BuyerRepository {
	return &BuyerRepository{
		db: db,
	}
}

func (b *BuyerRepository) Login(ctx *gin.Context, buyer entity.Buyer) (err error) {
	return
}

func (b *BuyerRepository) GetByID(ctx *gin.Context, id string) (buyer entity.Buyer, err error) {
	query := b.db.WithContext(ctx).First(&buyer, "id = ?", id)
	if query.Error != nil {
		err = query.Error
		return
	}

	return
}
