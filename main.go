package main

import (
	"GoMicro-Project/Service/ProdService"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// 注册到mdns
	// consulReg :=mdns.NewRegistry(
	// 	registry.Addrs("192.168.2.1"),
	// )
	// 注册到consul
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.2.1"),
	)

	ginRouter := gin.Default()
	ginGroup := ginRouter.Group("/v1")
	{
		ginGroup.Handle("GET", "/prods", func(context *gin.Context) {
			context.JSON(200, ProdService.NewProdList(5))
		})
	}
	ginRouter.Handle("GET", "/users", func(context *gin.Context) {
		context.String(200, "users api")
	})

	service := web.NewService(
		web.Name("ProdService"),
		web.Address(":8888"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)
	// Service.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
	// 	w.Write([]byte("hello world"))
	// })
	service.Run()
}
