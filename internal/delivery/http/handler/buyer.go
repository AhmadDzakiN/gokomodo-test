package handler

import (
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/service"
)

type BuyerHandler struct {
	BuyerService *service.BuyerService
}

func NewBuyerHandler(svc *service.BuyerService) *BuyerHandler {
	return &BuyerHandler{BuyerService: svc}
}

func (h *BuyerHandler) Login(ctx *gin.Context) {

}

func (h *BuyerHandler) GetProductList(ctx *gin.Context) {

}

func (h *BuyerHandler) CreateOrder(ctx *gin.Context) {

}

func (h *BuyerHandler) GetOrderList(ctx *gin.Context) {

}
