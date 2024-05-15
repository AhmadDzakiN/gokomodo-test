package repository

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/entity"
	"gorm.io/gorm"
)

type SellerRepository struct {
	db *gorm.DB
}

type ISellerRepository interface {
	Login(ctx *gin.Context, seller entity.Seller) (err error)
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{
		db: db,
	}
}

func (s *SellerRepository) Login(ctx *gin.Context, seller entity.Seller) (err error) {
	return
}
