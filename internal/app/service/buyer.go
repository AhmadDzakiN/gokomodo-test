package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gokomodo-assignment/internal/app/constant"
	"gokomodo-assignment/internal/app/entity"
	"gokomodo-assignment/internal/app/payloads"
	"gokomodo-assignment/internal/app/repository"
	"gokomodo-assignment/internal/pkg/pagination"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type BuyerService struct {
	Validator   *validator.Validate
	BuyerRepo   *repository.BuyerRepository
	ProductRepo *repository.ProductRepository
	OrderRepo   *repository.OrderRepository
	SellerRepo  *repository.SellerRepository
}

// modify the responses & requests
type IBuyerService interface {
	Login(ctx *gin.Context, request payloads.BuyerLoginRequest) (err error)
	GetProductList(ctx *gin.Context) (err error)
	CreateOrder(ctx *gin.Context) (err error)
	GetOrderList(ctx *gin.Context) (err error)
}

func NewBuyerService(validator *validator.Validate, buyerRepo *repository.BuyerRepository,
	productRepo *repository.ProductRepository, orderRepo *repository.OrderRepository, sellerRepo *repository.SellerRepository) *BuyerService {
	return &BuyerService{
		Validator:   validator,
		BuyerRepo:   buyerRepo,
		ProductRepo: productRepo,
		OrderRepo:   orderRepo,
		SellerRepo:  sellerRepo,
	}
}

func (b *BuyerService) Login(ctx *gin.Context, request payloads.BuyerLoginRequest) (err error) {
	return
}

func (b *BuyerService) GetProductList(ctx *gin.Context) (err error) {
	nextToken := strings.TrimSpace(ctx.Query("next"))
	params := payloads.GetProductListParams{
		LastValue: pagination.ParseGetListPageToken(nextToken),
		Limit:     constant.GetItemListLimit,
	}

	productList, err := b.ProductRepo.GetList(ctx, params)
	if err != nil {
		log.Err(err).Msgf("Failed to get Product List")
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	var response []payloads.GetProductListResponse
	for _, product := range productList {
		response = append(response, payloads.GetProductListResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Seller:      product.SellerID,
		})
	}

	newToken := pagination.CreateGetListPageToken(response, constant.GetItemListLimit)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "next_token": newToken, "data": response})
	return
}

func (b *BuyerService) CreateOrder(ctx *gin.Context) (err error) {
	var request payloads.CreateOrderRequest
	if err = ctx.ShouldBindJSON(&request); err != nil {
		log.Err(err).Msg("Invalid request format")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": err.Error()})
		return
	}

	product, err := b.ProductRepo.GetByID(ctx, request.Items)
	if err != nil {
		log.Err(err).Msgf("Failed to get Product by ID %d", request.Items)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	buyer, err := b.BuyerRepo.GetByID(ctx, "397a0a54-76a0-4eb7-b61b-ade1ee8676d9") // ambil dari jwt
	if err != nil {
		log.Err(err).Msgf("Failed to get Buyer by ID %s", "397a0a54-76a0-4eb7-b61b-ade1ee8676d9") // ambil dari jwt
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	seller, err := b.SellerRepo.GetByID(ctx, product.SellerID)
	if err != nil {
		log.Err(err).Msgf("Failed to get Seller by ID %s", product.SellerID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	order := entity.Order{
		BuyerID:            buyer.ID, //get from jwt
		SellerID:           product.SellerID,
		SourceAddress:      seller.PickupAddress,
		DestinationAddress: buyer.ShippingAddress,
		Items:              product.ID,
		Quantity:           request.Quantity,
		Price:              request.Price,
		TotalPrice:         request.TotalPrice,
		Status:             constant.OrderStatusPending,
	}

	err = b.OrderRepo.Create(ctx, &order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	response := payloads.CreateOrderResponse{
		ID:         order.ID,
		Items:      order.Items,
		Quantity:   order.Quantity,
		Price:      order.Price,
		TotalPrice: order.TotalPrice,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "data": response})
	return
}

func (b *BuyerService) GetOrderList(ctx *gin.Context) (err error) {
	nextToken := strings.TrimSpace(ctx.Query("next"))
	params := payloads.GetOrderListParams{
		LastValue: pagination.ParseGetListPageToken(nextToken),
		Limit:     constant.GetItemListLimit,
	}

	//get role and userid from jwt

	orderList, err := b.OrderRepo.GetList(ctx, params)
	if err != nil {
		log.Err(err).Msg("Failed to get Order List")
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	var response []payloads.GetOrderListResponse
	for _, order := range orderList {
		response = append(response, payloads.GetOrderListResponse{
			ID:                 order.ID,
			BuyerID:            order.BuyerID,
			SellerID:           order.SellerID,
			SourceAddress:      order.SourceAddress,
			DestinationAddress: order.DestinationAddress,
			Items:              order.Items,
			Quantity:           order.Quantity,
			Price:              order.Price,
			TotalPrice:         order.TotalPrice,
			Status:             order.Status,
		})
	}

	newToken := pagination.CreateGetListPageToken(response, constant.GetItemListLimit)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "next_token": newToken, "data": response})
	return
}
