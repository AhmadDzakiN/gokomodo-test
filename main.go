package main

import (
	"gokomodo-assignment/internal/config"
)

func main() {
	cfg := config.NewViperConfig()
	log := config.NewLogger(cfg)
	validator := config.NewValidator(cfg)
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
	log.Info().Msgf("Start Toko API at port %s", port)
	app.Run(":" + port)
}
