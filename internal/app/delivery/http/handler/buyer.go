package handler

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/app/service"
)

type BuyerHandler struct {
	BuyerService *service.BuyerService
}

func NewBuyerHandler(svc *service.BuyerService) *BuyerHandler {
	return &BuyerHandler{BuyerService: svc}
}

func (b *BuyerHandler) Login(ctx *gin.Context) {

}

func (b *BuyerHandler) GetProductList(ctx *gin.Context) {
	b.BuyerService.GetProductList(ctx)
}

func (b *BuyerHandler) CreateOrder(ctx *gin.Context) {
	b.BuyerService.CreateOrder(ctx)
}

func (b *BuyerHandler) GetOrderList(ctx *gin.Context) {
	b.BuyerService.GetOrderList(ctx)
}
