package Routers

import (
	"GoMicro-Project/Handlers"
	"GoMicro-Project/Middlewares"
	"GoMicro-Project/ServiceImpl"
	"github.com/gin-gonic/gin"
)

func NewGinRouterWithService(prodService ServiceImpl.ProdService1) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(Middlewares.InitMiddleware(prodService))
	ginGroup := ginRouter.Group("/v1")
	{
		ginGroup.Handle("GET", "/prods", Handlers.GetProdsList)
	}
	ginRouter.Handle("GET", "/users", Handlers.GetUser)

	return ginRouter
}

func NewGinRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginGroup := ginRouter.Group("/v1")
	{
		ginGroup.Handle("GET", "/prods", Handlers.GetProdsList)
	}
	ginRouter.Handle("GET", "/users", Handlers.GetUser)

	return ginRouter
}
