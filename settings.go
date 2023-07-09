package tests

import (
	"financialManagement/internal/config"
	"financialManagement/internal/database/postgres"
)

var (
	cfg     = config.MustGetConfig("/Users/kare/GolandProjects/financialManagement/config_data/config.yaml")
	STORAGE = postgres.MustNewStorage(cfg.Postgres)
)
