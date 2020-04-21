package main

import (
	"GoMicro-Project/Routers"
	"GoMicro-Project/Service/Micros"
	"GoMicro-Project/ServiceImpl"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/mdns"
	"github.com/micro/go-micro/web"
)

func main() {
	// 注册到mdns
	mdnsReg := mdns.NewRegistry(
		registry.Addrs("192.168.2.1"),
	)
	// 注册到consul
	// consulReg := consul.NewRegistry(
	// 	registry.Addrs("192.168.2.1"),
	// )

	httpService := web.NewService(
		web.Name("ProdServiceHttp"),
		web.Address(":8888"),
		web.Handler(Routers.NewGinRouter()),
		// web.Registry(consulReg),
	)
	rpcService := micro.NewService(
		micro.Name("ProdServiceRPC"),
		micro.Address(":8881"),
		// micro.Registry(consulReg),
		micro.Registry(mdnsReg),
	)
	// 注册并启动rpc服务
	rpcService.Init()
	Micros.RegisterProdService1Handler(rpcService.Server(), new(ServiceImpl.ProdService1))
	rpcService.Run()

	// TestService.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
	// 	w.Write([]byte("hello world"))
	// })
	// 启动http服务
	httpService.Run()
}
