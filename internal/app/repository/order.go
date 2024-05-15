package repository

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/app/constant"
	"gokomodo-assignment/internal/app/entity"
	"gokomodo-assignment/internal/app/payloads"
	"gorm.io/gorm"
	"time"
)

type OrderRepository struct {
	db *gorm.DB
}

type IOrderRepository interface {
	Accept(ctx *gin.Context) (err error)
	GetList(ctx *gin.Context, params payloads.GetOrderListParams) (orders []entity.Order, err error)
	Create(ctx *gin.Context, order *entity.Order) (err error)
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
	query := o.db.WithContext(ctx).Table("orders o").
		Select("o.id, o.buyer_id, o.seller_id, o.source_address, o.destination_address," +
			"o.items, o.quantity, o.price, o.total_price, o.status")

	if params.UserID != "" && params.Role == constant.SellerRole {
		query = query.Where("o.seller_id = ?", params.UserID)
	} else if params.UserID != "" && params.Role == constant.BuyerRole {
		query = query.Where("o.buyer_id = ?", params.UserID)
	}

	if params.LastValue > 0 {
		lastUpdated := time.Unix(int64(params.LastValue), 0)
		query = query.Where("o.updated_at > ?", lastUpdated)
	}

	query.Limit(params.Limit)
	query.Find(&orders)
	if query.Error != nil {
		err = query.Error
		return
	}

	return
}

func (o *OrderRepository) Create(ctx *gin.Context, order *entity.Order) (err error) {
	query := o.db.WithContext(ctx).Create(order)
	if query.Error != nil || query.RowsAffected < 1 {
		err = query.Error
		return
	}
	return
}
