package repository

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/entity"
	"gorm.io/gorm"
)

type BuyerRepository struct {
	db *gorm.DB
}

type IBuyerRepository interface {
	Login(ctx *gin.Context, buyer entity.Buyer) (err error)
}

func NewBuyerRepository(db *gorm.DB) *BuyerRepository {
	return &BuyerRepository{
		db: db,
	}
}

func (b *BuyerRepository) Login(ctx *gin.Context, buyer entity.Buyer) (err error) {
	return
}