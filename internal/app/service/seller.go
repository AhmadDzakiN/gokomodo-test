package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gokomodo-assignment/internal/app/payloads"
	"gokomodo-assignment/internal/app/repository"
)

type SellerService struct {
	Validator   *validator.Validate
	SellerRepo  *repository.SellerRepository
	ProductRepo *repository.ProductRepository
	OrderRepo   *repository.OrderRepository
}

// modify the responses & requests
type ISellerService interface {
	Login(ctx *gin.Context, request payloads.SellerLoginRequest) (err error)
	GetProductList(ctx *gin.Context) (err error)
	CreateProduct(ctx *gin.Context) (err error)
	AcceptOrder(ctx *gin.Context) (err error)
	GetOrderList(ctx *gin.Context) (err error)
}

func NewSellerService(validator *validator.Validate, sellerRepo *repository.SellerRepository,
	productRepo *repository.ProductRepository, orderRepo *repository.OrderRepository) *SellerService {
	return &SellerService{
		Validator:   validator,
		SellerRepo:  sellerRepo,
		ProductRepo: productRepo,
		OrderRepo:   orderRepo,
	}
}

func (s *SellerService) Login(ctx *gin.Context, request payloads.SellerLoginRequest) (err error) {
	return
}

func (s *SellerService) GetProductList(ctx *gin.Context) (err error) {
	return
}

func (s *SellerService) CreateProduct(ctx *gin.Context) (err error) {
	return
}

func (s *SellerService) AcceptOrder(ctx *gin.Context) (err error) {
	return
}

func (s *SellerService) GetOrderList(ctx *gin.Context) (err error) {
	return
}
