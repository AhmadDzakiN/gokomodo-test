package repository

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/entity"
	"gokomodo-assignment/internal/payloads"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type IProductRepository interface {
	GetList(ctx *gin.Context, params payloads.GetProductListParams) (products []entity.Product, err error)
	Create(ctx *gin.Context, product *entity.Product) (err error)
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetList(ctx *gin.Context, params payloads.GetProductListParams) (products []entity.Product, err error) {
	//query := pr.db.WithContext(ctx).Table("products p").
	//	Select("p.id, p.name, p.description, p.price, p.seller_id")

	return
}

func (pr *ProductRepository) Create(ctx *gin.Context, product *entity.Product) (err error) {
	return
}
