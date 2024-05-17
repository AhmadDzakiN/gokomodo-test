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
	"gokomodo-assignment/internal/pkg/jwt"
	"gokomodo-assignment/internal/pkg/pagination"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type BuyerService struct {
	Validator   *validator.Validate
	BuyerRepo   repository.IBuyerRepository
	ProductRepo repository.IProductRepository
	OrderRepo   repository.IOrderRepository
	SellerRepo  repository.ISellerRepository
}

type IBuyerService interface {
	Login(ctx *gin.Context)
	GetProductList(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	GetOrderList(ctx *gin.Context)
}

func NewBuyerService(validator *validator.Validate, buyerRepo repository.IBuyerRepository,
	productRepo repository.IProductRepository, orderRepo repository.IOrderRepository, sellerRepo repository.ISellerRepository) *BuyerService {
	return &BuyerService{
		Validator:   validator,
		BuyerRepo:   buyerRepo,
		ProductRepo: productRepo,
		OrderRepo:   orderRepo,
		SellerRepo:  sellerRepo,
	}
}

func (b *BuyerService) Login(ctx *gin.Context) {
	var request payloads.BuyerLoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Err(err).Msg("Empty or invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Empty or invalid request"})
		return
	}

	if err := b.Validator.Struct(request); err != nil {
		log.Err(err).Msg("Empty or invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Empty or invalid request"})
		return
	}

	buyer, err := b.BuyerRepo.GetByEmail(ctx, request.Email)
	if err != nil {
		log.Err(err).Msgf("Failed to get Buyer by Email %s", request.Email)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": "Buyer not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
		return
	}

	//err = bcrypt.CompareHashAndPassword([]byte(buyer.Password), []byte(request.Password))
	//if err != nil {
	//	log.Err(err).Msgf("Invalid or wrong password")
	//	ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Invalid or wrong password"})
	//	return
	//}

	token, err := jwt.CreateToken(buyer.ID, buyer.Name, constant.BuyerRole)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Failed to create Buyer token"})
		return
	}

	response := payloads.BuyerLoginResponse{
		Email: buyer.Email,
		Name:  buyer.Name,
		Token: token,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "data": response})

	return
}

func (b *BuyerService) GetProductList(ctx *gin.Context) {
	nextToken := strings.TrimSpace(ctx.Query("next"))
	params := payloads.GetProductListParams{
		LastValue: pagination.ParseGetListPageToken(nextToken),
		Limit:     constant.GetItemListLimit,
	}

	productList, err := b.ProductRepo.GetList(ctx, params)
	if err != nil {
		log.Err(err).Msgf("Failed to get Product List")
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
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
			UpdatedAt:   product.UpdatedAt.Unix(),
		})
	}

	newToken := pagination.CreateGetListPageToken(response, constant.GetItemListLimit)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "next_token": newToken, "data": response})
	return
}

func (b *BuyerService) CreateOrder(ctx *gin.Context) {
	var request payloads.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Err(err).Msg("Empty or invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Empty or invalid request"})
		return
	}

	if err := b.Validator.Struct(request); err != nil {
		log.Err(err).Msg("Empty or invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Empty or invalid request"})
		return
	}

	product, err := b.ProductRepo.GetByID(ctx, request.Items)
	if err != nil {
		log.Err(err).Msgf("Failed to get Product by ID %d", request.Items)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": "Product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
		return
	}

	ogTotalPrice := uint64(request.Quantity) * product.Price
	if uint64(request.Quantity)*request.Price != ogTotalPrice || request.TotalPrice != ogTotalPrice {
		log.Warn().Msg("Total price input is not same with the original total price")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Empty or invalid request"})
		return
	}

	tokenClaims, err := jwt.GetTokenClaims(ctx)
	if err != nil {
		log.Err(err).Msg("Invalid user token")
		ctx.JSON(http.StatusForbidden, gin.H{"status": "error", "status_code": http.StatusUnauthorized, "error": "Access forbidden"})
		ctx.Abort()
		return
	}

	buyerID := tokenClaims.UserID
	buyer, err := b.BuyerRepo.GetByID(ctx, buyerID)
	if err != nil {
		log.Err(err).Msgf("Failed to get Buyer by ID %s", buyerID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": "Buyer not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
		return
	}

	seller, err := b.SellerRepo.GetByID(ctx, product.SellerID)
	if err != nil {
		log.Err(err).Msgf("Failed to get Seller by ID %s", product.SellerID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": "Seller not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
		return
	}

	newOrder := entity.Order{
		BuyerID:            buyer.ID,
		SellerID:           product.SellerID,
		SourceAddress:      seller.PickupAddress,
		DestinationAddress: buyer.ShippingAddress,
		Items:              product.ID,
		Quantity:           request.Quantity,
		Price:              request.Price,
		TotalPrice:         request.TotalPrice,
		Status:             constant.OrderStatusPending,
	}

	err = b.OrderRepo.Create(ctx, &newOrder)
	if err != nil {
		log.Err(err).Msgf("Failed to create new Order for Buyer %s", buyer.ID)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
		return
	}

	response := payloads.CreateOrderResponse{
		ID:         newOrder.ID,
		Items:      newOrder.Items,
		Quantity:   newOrder.Quantity,
		Price:      newOrder.Price,
		TotalPrice: newOrder.TotalPrice,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "data": response})
	return
}

func (b *BuyerService) GetOrderList(ctx *gin.Context) {
	nextToken := strings.TrimSpace(ctx.Query("next"))
	params := payloads.GetOrderListParams{
		LastValue: pagination.ParseGetListPageToken(nextToken),
		Limit:     constant.GetItemListLimit,
	}

	tokenClaims, err := jwt.GetTokenClaims(ctx)
	if err != nil {
		log.Err(err).Msg("Invalid user token")
		ctx.JSON(http.StatusForbidden, gin.H{"status": "error", "status_code": http.StatusForbidden, "error": "Access forbidden"})
		ctx.Abort()
		return
	}

	params.UserID = tokenClaims.UserID
	params.Role = tokenClaims.Role
	orderList, err := b.OrderRepo.GetList(ctx, params)
	if err != nil {
		log.Err(err).Msg("Failed to get Order List")
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
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
			UpdatedAt:          order.UpdatedAt.Unix(),
		})
	}

	newToken := pagination.CreateGetListPageToken(response, constant.GetItemListLimit)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "next_token": newToken, "data": response})
	return
}
