package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gokomodo-assignment/internal/payloads"
	"gokomodo-assignment/internal/repository"
)

type BuyerService struct {
	Validator   *validator.Validate
	BuyerRepo   *repository.BuyerRepository
	ProductRepo *repository.ProductRepository
	OrderRepo   *repository.OrderRepository
}

// modify the responses & requests
type IBuyerService interface {
	Login(ctx *gin.Context, request payloads.BuyerLoginRequest) (err error)
	GetProductList(ctx *gin.Context) (err error)
	CreateOrder(ctx *gin.Context) (err error)
	GetOrderList(ctx *gin.Context) (err error)
}

func NewBuyerService(validator *validator.Validate, buyerRepo *repository.BuyerRepository,
	productRepo *repository.ProductRepository, orderRepo *repository.OrderRepository) *BuyerService {
	return &BuyerService{
		Validator:   validator,
		BuyerRepo:   buyerRepo,
		ProductRepo: productRepo,
		OrderRepo:   orderRepo,
	}
}

func (b *BuyerService) Login(ctx *gin.Context, request payloads.BuyerLoginRequest) (err error) {
	return
}

func (b *BuyerService) GetProductList(ctx *gin.Context) (err error) {
	return
}

func (b *BuyerService) CreateOrder(ctx *gin.Context) (err error) {
	return
}

func (b *BuyerService) GetOrderList(ctx *gin.Context) (err error) {
	return
}
