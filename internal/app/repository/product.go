package repository

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/app/entity"
	"gokomodo-assignment/internal/app/payloads"
	"gorm.io/gorm"
	"time"
)

type ProductRepository struct {
	db *gorm.DB
}

type IProductRepository interface {
	GetList(ctx *gin.Context, params payloads.GetProductListParams) (products []entity.Product, err error)
	Create(ctx *gin.Context, product *entity.Product) (err error)
	GetByID(ctx *gin.Context, id uint64) (product entity.Product, err error)
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) GetList(ctx *gin.Context, params payloads.GetProductListParams) (products []entity.Product, err error) {
	query := p.db.WithContext(ctx).Table("products p").
		Select("p.id, p.name, p.description, p.price, p.seller_id")

	if params.SellerID != "" {
		query = query.Where("p.seller_id = ?", params.SellerID)
	}

	if params.LastValue > 0 {
		lastUpdated := time.Unix(int64(params.LastValue), 0)
		query = query.Where("p.updated_at > ?", lastUpdated)
	}

	query.Limit(params.Limit)
	query.Find(&products)
	if query.Error != nil {
		err = query.Error
		return
	}

	return
}

func (p *ProductRepository) Create(ctx *gin.Context, product *entity.Product) (err error) {
	return
}

func (p *ProductRepository) GetByID(ctx *gin.Context, id uint64) (product entity.Product, err error) {
	query := p.db.WithContext(ctx).First(&product, "id = ?", id)
	if query.Error != nil {
		err = query.Error
		return
	}

	return
}
