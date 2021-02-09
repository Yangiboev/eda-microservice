package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "gitlab.udevs.io/macbro/mb_corporate_service/api/v1"
	"gitlab.udevs.io/macbro/mb_corporate_service/config"
	"gitlab.udevs.io/macbro/mb_corporate_service/pkg/logger"
	"gitlab.udevs.io/macbro/mb_corporate_service/storage"
)

type RouterOptions struct {
	Log     logger.Logger
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cfg := cors.DefaultConfig()

	cfg.AllowHeaders = append(cfg.AllowHeaders, "*")
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true

	router.Use(cors.New(cfg))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:     opt.Log,
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")

	apiV1.Use()
	{
		apiV1.GET("/company/:id", handlerV1.Get)
		apiV1.GET("/company", handlerV1.GetAll)
	}

	return router
}
