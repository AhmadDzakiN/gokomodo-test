package handler

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/app/service"
	"net/http"
)

type SellerHandler struct {
	SellerService *service.SellerService
}

func NewSellerHandler(svc *service.SellerService) *SellerHandler {
	return &SellerHandler{SellerService: svc}
}

func (s *SellerHandler) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": "ok"}})
}

func (s *SellerHandler) GetProductList(ctx *gin.Context) {

}

func (s *SellerHandler) CreateProduct(ctx *gin.Context) {

}

func (s *SellerHandler) AcceptOrder(ctx *gin.Context) {

}

func (s *SellerHandler) GetOrderList(ctx *gin.Context) {

}
