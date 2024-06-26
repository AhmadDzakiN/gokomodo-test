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
	"strconv"
	"strings"
)

type SellerService struct {
	Validator   *validator.Validate
	SellerRepo  repository.ISellerRepository
	ProductRepo repository.IProductRepository
	OrderRepo   repository.IOrderRepository
}

type ISellerService interface {
	Login(ctx *gin.Context)
	GetProductList(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	AcceptOrder(ctx *gin.Context)
	GetOrderList(ctx *gin.Context)
}

func NewSellerService(validator *validator.Validate, sellerRepo repository.ISellerRepository,
	productRepo repository.IProductRepository, orderRepo repository.IOrderRepository) *SellerService {
	return &SellerService{
		Validator:   validator,
		SellerRepo:  sellerRepo,
		ProductRepo: productRepo,
		OrderRepo:   orderRepo,
	}
}

func (s *SellerService) Login(ctx *gin.Context) {
	var request payloads.SellerLoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Err(err).Msg("Invalid request format")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Empty or invalid request"})
		return
	}

	if err := s.Validator.Struct(request); err != nil {
		log.Err(err).Msg("Empty or invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Empty or invalid request"})
		return
	}

	seller, err := s.SellerRepo.GetByEmail(ctx, request.Email)
	if err != nil {
		log.Err(err).Msgf("Failed to get Seller by Email %s", request.Email)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": "Seller not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
		return
	}

	//err = bcrypt.CompareHashAndPassword([]byte(seller.Password), []byte(request.Password))
	//if err != nil {
	//	log.Err(err).Msgf("Invalid or wrong password")
	//	ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Invalid or wrong password"})
	//	return
	//}

	token, err := jwt.CreateToken(seller.ID, seller.Name, constant.SellerRole)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Failed to create Seller token"})
		return
	}

	response := payloads.SellerLoginResponse{
		Email: seller.Email,
		Name:  seller.Name,
		Token: token,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "data": response})
	return
}

func (s *SellerService) GetProductList(ctx *gin.Context) {
	nextToken := strings.TrimSpace(ctx.Query("next"))
	params := payloads.GetProductListParams{
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

	params.SellerID = tokenClaims.UserID
	productList, err := s.ProductRepo.GetList(ctx, params)
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

func (s *SellerService) CreateProduct(ctx *gin.Context) {
	var request payloads.CreateProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Err(err).Msg("Empty or invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Empty or invalid request"})
		return
	}

	if err := s.Validator.Struct(request); err != nil {
		log.Err(err).Msg("Empty or invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Empty or invalid request"})
		return
	}

	newProduct := entity.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}

	tokenClaims, err := jwt.GetTokenClaims(ctx)
	if err != nil {
		log.Err(err).Msg("Invalid user token")
		ctx.JSON(http.StatusForbidden, gin.H{"status": "error", "status_code": http.StatusUnauthorized, "error": "Access forbidden"})
		ctx.Abort()
		return
	}

	newProduct.SellerID = tokenClaims.UserID
	err = s.ProductRepo.Create(ctx, &newProduct)
	if err != nil {
		log.Err(err).Msgf("Failed to create new Product for Seller %s", newProduct.SellerID)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
		return
	}

	response := payloads.CreateProductResponse{
		ID:          newProduct.ID,
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		SellerID:    newProduct.SellerID,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "data": response})
	return
}

func (s *SellerService) AcceptOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("order_id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		log.Err(err).Msg("Empty or invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": "Empty or invalid request"})
		return
	}

	tokenClaims, err := jwt.GetTokenClaims(ctx)
	if err != nil {
		log.Err(err).Msg("Invalid user token")
		ctx.JSON(http.StatusForbidden, gin.H{"status": "error", "status_code": http.StatusForbidden, "error": "Access forbidden"})
		ctx.Abort()
		return
	}

	order, err := s.OrderRepo.GetByID(ctx, orderID)
	if err != nil {
		log.Err(err).Msgf("Failed to get order by ID %d", orderID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": "Order not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
		return
	}

	if order.SellerID != tokenClaims.UserID {
		log.Warn().Msgf("No product found for this seller. SellerID: %s", tokenClaims.UserID)
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": "No product found for this seller"})
		return
	}

	err = s.OrderRepo.Accept(ctx, orderID, tokenClaims.UserID)
	if err != nil {
		log.Err(err).Msgf("Failed to get Accept Order for ID %d", orderID)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": "Ups, something wrong with the server. Please try again later"})
		return
	}

	response := payloads.AcceptOrderResponse{
		ID:    order.ID,
		Items: order.Items,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "data": response})
	return
}

func (s *SellerService) GetOrderList(ctx *gin.Context) {
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
	orderList, err := s.OrderRepo.GetList(ctx, params)
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
