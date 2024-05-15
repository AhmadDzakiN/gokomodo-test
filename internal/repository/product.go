package repository

import (
	"context"
	"gokomodo-assignment/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetList(ctx context.Context, params entity.GetProductListParams) (products []entity.Product, err error) {
	query := pr.db.WithContext(ctx).Table("products p").
		Select("p.id, p.name, p.description, p.price, p.seller_id")

	return
}
