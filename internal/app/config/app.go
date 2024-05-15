package config

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gokomodo-assignment/internal/app/delivery/http/handler"
	"gokomodo-assignment/internal/app/delivery/http/route"
	"gokomodo-assignment/internal/app/repository"
	"gokomodo-assignment/internal/app/service"
	"gorm.io/gorm"
)

type BootstrapAppConfig struct {
	DB        *gorm.DB
	Validator *validator.Validate
	Config    *viper.Viper
}

func BootstrapApp(config *BootstrapAppConfig) (app *gin.Engine) {
	buyerRepository := repository.NewBuyerRepository(config.DB)
	orderRepository := repository.NewOrderRepository(config.DB)
	productRepository := repository.NewProductRepository(config.DB)
	sellerRepository := repository.NewSellerRepository(config.DB)

	buyerService := service.NewBuyerService(config.Validator, buyerRepository, productRepository, orderRepository, sellerRepository)
	sellerService := service.NewSellerService(config.Validator, sellerRepository, productRepository, orderRepository)

	buyerHandler := handler.NewBuyerHandler(buyerService)
	sellerHandler := handler.NewSellerHandler(sellerService)

	app = route.Router(&route.RouteConfig{
		SellerHandler: sellerHandler,
		BuyerHandler:  buyerHandler,
	})

	return
}
