package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gokomodo-assignment/internal/app/delivery/http/handler"
	"net/http"
)

type RouteConfig struct {
	SellerHandler *handler.SellerHandler
	BuyerHandler  *handler.BuyerHandler
}

func Router(cfg *RouteConfig) (router *gin.Engine) {
	corsConfig := cors.Config{
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization", "Content-Length"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}

	corsOpt := cors.New(corsConfig)
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	router.Use(corsOpt)
	router.Use(gin.Recovery())
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	seller := router.Group("/seller")
	{
		seller.POST("/login", cfg.SellerHandler.Login)
		seller.GET("/products", cfg.SellerHandler.GetProductList)
		seller.POST("/products", cfg.SellerHandler.CreateProduct)
		seller.GET("/orders", cfg.SellerHandler.GetOrderList)
		seller.PUT("/orders/:order_id", cfg.SellerHandler.AcceptOrder)
	}

	buyer := router.Group("/buyer")
	{
		buyer.POST("/login", cfg.BuyerHandler.Login)
		buyer.GET("/products", cfg.BuyerHandler.GetProductList)
		buyer.POST("/orders", cfg.BuyerHandler.CreateOrder)
		buyer.GET("/orders", cfg.BuyerHandler.GetOrderList)
	}

	return
}
