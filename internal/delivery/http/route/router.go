package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() (router *gin.Engine) {
	corsConfig := cors.Config{
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization", "Content-Length"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}

	corsOpt := cors.New(corsConfig)
	router = gin.Default()
	router.Use(corsOpt)
	router.Use(gin.Recovery())

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	seller := router.Group("/seller")
	{
		seller.POST("/login")
		seller.GET("/products")
		seller.POST("/product")
		seller.GET("/orders")
		seller.PUT("/order")
	}

	buyer := router.Group("/buyer")
	{
		buyer.POST("/login")
		buyer.GET("/products")
		buyer.POST("/create")
		buyer.GET("/orders")
	}

	return
}
