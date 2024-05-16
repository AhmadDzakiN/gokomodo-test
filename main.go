package main

import (
	"gokomodo-assignment/internal/app/config"
)

func main() {
	cfg := config.NewViperConfig()
	log := config.NewLogger(cfg)
	validator := config.NewValidator()
	db, err := config.NewPostgreDatabase(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start, error connect to DB Postgre")
		return
	}

	app := config.BootstrapApp(&config.BootstrapAppConfig{
		DB:        db,
		Validator: validator,
		Config:    cfg,
	})

	port := cfg.GetString("APP_PORT")
	appName := cfg.GetString("APP_NAME")
	log.Info().Msgf("Start %s at port %s", appName, port)
	app.Run(":" + port)
}
