package main

import (
	"gokomodo-assignment/internal/config"
	"gokomodo-assignment/internal/delivery/http/route"
)

func main() {
	cfg := config.NewViperConfig()
	log := config.NewLogger(cfg)
	db, err := config.NewPostgreDatabase(cfg)
	validate := config.NewValidator(cfg)

	app := route.Router()
	port := cfg.GetString("APP_PORT")
	app.Run(":" + port)
}
