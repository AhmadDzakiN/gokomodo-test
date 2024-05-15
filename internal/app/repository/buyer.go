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
	GetByID(ctx *gin.Context, id string) (buyer entity.Buyer, err error)
	GetByEmail(ctx *gin.Context, email string) (buyer entity.Buyer, err error)
}

func NewBuyerRepository(db *gorm.DB) *BuyerRepository {
	return &BuyerRepository{
		db: db,
	}
}

func (b *BuyerRepository) GetByID(ctx *gin.Context, id string) (buyer entity.Buyer, err error) {
	query := b.db.WithContext(ctx).First(&buyer, "id = ?", id)
	if query.Error != nil {
		err = query.Error
		return
	}

	return
}

func (s *BuyerRepository) GetByEmail(ctx *gin.Context, email string) (buyer entity.Buyer, err error) {
	query := s.db.WithContext(ctx).First(&buyer, "email = ?", email)
	if query.Error != nil {
		err = query.Error
		return
	}

	return
}
