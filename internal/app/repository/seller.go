package repository

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/app/entity"
	"gorm.io/gorm"
)

type SellerRepository struct {
	db *gorm.DB
}

type ISellerRepository interface {
	Login(ctx *gin.Context, seller entity.Seller) (err error)
	GetByID(ctx *gin.Context, id string) (seller entity.Seller, err error)
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{
		db: db,
	}
}

func (s *SellerRepository) Login(ctx *gin.Context, seller entity.Seller) (err error) {
	return
}

func (s *SellerRepository) GetByID(ctx *gin.Context, id string) (seller entity.Seller, err error) {
	query := s.db.WithContext(ctx).First(&seller, "id = ?", id)
	if query.Error != nil {
		err = query.Error
		return
	}

	return
}
