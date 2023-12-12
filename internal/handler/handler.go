package handler

import (
	"github.com/YuStep/wbl0/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	//router.LoadHTMLGlob("assets/*")

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	{
		order := v1.Group("/order")
		{
			order.GET("", h.getAllOrders)
			order.GET("/:id", h.getOrderById)
		}
	}
	return router
}
