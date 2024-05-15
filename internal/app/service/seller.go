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
	SellerRepo  *repository.SellerRepository
	ProductRepo *repository.ProductRepository
	OrderRepo   *repository.OrderRepository
}

// modify the responses & requests
type ISellerService interface {
	Login(ctx *gin.Context) (err error)
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

func (s *SellerService) Login(ctx *gin.Context) (err error) {
	var request payloads.SellerLoginRequest
	if err = ctx.ShouldBindJSON(&request); err != nil {
		log.Err(err).Msg("Invalid request format")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": err.Error()})
		return
	}

	seller, err := s.SellerRepo.GetByEmail(ctx, request.Email)
	if err != nil {
		log.Err(err).Msgf("Failed to get Seller by Email %s", request.Email)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "status_code": http.StatusNotFound, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": err.Error()})
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

func (s *SellerService) GetProductList(ctx *gin.Context) (err error) {
	nextToken := strings.TrimSpace(ctx.Query("next"))
	params := payloads.GetProductListParams{
		LastValue: pagination.ParseGetListPageToken(nextToken),
		Limit:     constant.GetItemListLimit,
		SellerID:  "802988d9-054f-4a96-af17-bb8bffdea411", // change this to get from jwt
	}

	productList, err := s.ProductRepo.GetList(ctx, params)
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

func (s *SellerService) CreateProduct(ctx *gin.Context) (err error) {
	var request payloads.CreateProductRequest
	if err = ctx.ShouldBindJSON(&request); err != nil {
		log.Err(err).Msg("Invalid request format")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": err.Error()})
		return
	}

	newProduct := entity.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		SellerID:    "802988d9-054f-4a96-af17-bb8bffdea411", //get from jwt
	}

	err = s.ProductRepo.Create(ctx, &newProduct)
	if err != nil {
		log.Err(err).Msgf("Failed to create new Product for Seller %s", "802988d9-054f-4a96-af17-bb8bffdea411") //get from jWT
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	response := payloads.CreateProductResponse{
		ID:          newProduct.ID,
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		SellerID:    "802988d9-054f-4a96-af17-bb8bffdea411", //get from jwt
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK, "data": response})
	return
}

func (s *SellerService) AcceptOrder(ctx *gin.Context) (err error) {
	orderIDStr := ctx.Param("order_id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		log.Err(err).Msg("Invalid order_id format")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "status_code": http.StatusBadRequest, "error": err.Error()})
		return
	}

	err = s.OrderRepo.Accept(ctx, orderID, "802988d9-054f-4a96-af17-bb8bffdea411") //get selleriD from JWT
	if err != nil {
		log.Err(err).Msgf("Failed to get Accept Order for ID %d", orderID)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "status_code": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "status_code": http.StatusOK})
	return
}

func (s *SellerService) GetOrderList(ctx *gin.Context) (err error) {
	nextToken := strings.TrimSpace(ctx.Query("next"))
	params := payloads.GetOrderListParams{
		LastValue: pagination.ParseGetListPageToken(nextToken),
		Limit:     constant.GetItemListLimit,
	}

	//get role and userid from jwt

	orderList, err := s.OrderRepo.GetList(ctx, params)
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
