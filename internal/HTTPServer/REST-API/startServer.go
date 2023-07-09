package REST_API

import (
	"financialManagement/internal/config"
	"financialManagement/internal/database/postgres"
	"github.com/gin-gonic/gin"
)

func StartServer(cfg *config.HTTPServerConfig, s *postgres.Storage) error {
	router := gin.Default()
	router.GET("/", SS)
	err := router.Run(cfg.Address)
	return err
}

func SS(ctx *gin.Context) {
	ctx.Writer.Write([]byte("Hello World!"))
}
