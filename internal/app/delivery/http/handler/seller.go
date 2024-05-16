package handler

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/app/service"
)

type SellerHandler struct {
	SellerService service.ISellerService
}

func NewSellerHandler(svc *service.SellerService) *SellerHandler {
	return &SellerHandler{SellerService: svc}
}

func (s *SellerHandler) Login(ctx *gin.Context) {
	s.SellerService.Login(ctx)
}

func (s *SellerHandler) GetProductList(ctx *gin.Context) {
	s.SellerService.GetProductList(ctx)
}

func (s *SellerHandler) CreateProduct(ctx *gin.Context) {
	s.SellerService.CreateProduct(ctx)
}

func (s *SellerHandler) AcceptOrder(ctx *gin.Context) {
	s.SellerService.AcceptOrder(ctx)
}

func (s *SellerHandler) GetOrderList(ctx *gin.Context) {
	s.SellerService.GetOrderList(ctx)
}
