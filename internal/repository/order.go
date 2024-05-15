package repository

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/entity"
	"gokomodo-assignment/internal/payloads"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

type IOrderRepository interface {
	Accept(ctx *gin.Context) (err error)
	GetList(ctx *gin.Context, params payloads.GetOrderListParams) (orders []entity.Order, err error)
	Create(ctx *gin.Context, order entity.Order) (err error)
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) Accept(ctx *gin.Context) (err error) {
	return
}

func (o *OrderRepository) GetList(ctx *gin.Context, params payloads.GetOrderListParams) (orders []entity.Order, err error) {
	return
}

func (o *OrderRepository) Create(ctx *gin.Context, order entity.Order) (err error) {
	return
}
