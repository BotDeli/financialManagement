package app

import (
	REST_API "financialManagement/internal/HTTPServer/REST-API"
	"financialManagement/internal/config"
	"financialManagement/internal/database/postgres"
	"financialManagement/internal/logger"
)

func StartApplication() {
	cfg := config.MustGetOSConfig()
	logg := logger.MustStartLogger(cfg.Logger)
	store := postgres.MustNewStorage(cfg.Postgres)
	defer logg.Warn("Problems disconnect storage:", postgres.Disconnect(store))
	logg.Error("Problems start server:", REST_API.StartServer(&cfg.HttpServer, store))
}
