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
	GetByID(ctx *gin.Context, id string) (seller entity.Seller, err error)
	GetByEmail(ctx *gin.Context, email string) (seller entity.Seller, err error)
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{
		db: db,
	}
}

func (s *SellerRepository) GetByID(ctx *gin.Context, id string) (seller entity.Seller, err error) {
	query := s.db.WithContext(ctx).First(&seller, "id = ?", id)
	if query.Error != nil {
		err = query.Error
		return
	}

	return
}

func (s *SellerRepository) GetByEmail(ctx *gin.Context, email string) (seller entity.Seller, err error) {
	query := s.db.WithContext(ctx).First(&seller, "email = ?", email)
	if query.Error != nil {
		err = query.Error
		return
	}

	return
}
