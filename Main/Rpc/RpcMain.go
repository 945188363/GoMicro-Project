package main

import (
	"GoMicro-Project/Service/Micros"
	"GoMicro-Project/ServiceImpl"
	"GoMicro-Project/Wrapper"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)

func main() {
	// 注册到mdns
	// mdnsReg := mdns.NewRegistry(
	// 	registry.Addrs("192.168.2.1"),
	// )
	// 注册到consul docker ip 172.17.0.2
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)

	// httpService := web.NewService(
	// 	web.Name("ProdServiceHttp"),
	// 	web.Address(":8012"),
	// 	web.Handler(Routers.NewGinRouter()),
	// 	web.Registry(consulReg),
	// )
	// rpcService := micro.NewService(
	// 	micro.Name("ProdServiceRPC"),
	// 	micro.Address(":8011"),
	// 	micro.Registry(consulReg),
	// 	// micro.Registry(mdnsReg),
	// 	micro.WrapClient(Wrapper.NewLogWrapper),
	// )

	rpcService2 := micro.NewService(
		micro.Name("RpcService"),
		micro.Address(":8011"),
		micro.Registry(consulReg),
		// micro.Registry(mdnsReg),
		micro.WrapClient(Wrapper.NewLogWrapper),
	)
	// httpService.Init()
	// rpcService.Init()
	rpcService2.Init()
	// 注册并启动rpc服务
	// if err := Micros.RegisterProdService1Handler(rpcService.Server(), new(ServiceImpl.ProdService1)); err != nil {
	// 	log.Fatal(err)
	// }
	if err := Micros.RegisterRpcServiceHandler(rpcService2.Server(), new(ServiceImpl.RpcService)); err != nil {
		log.Fatal(err)
	}
	// if err := rpcService.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	if err := rpcService2.Run(); err != nil {
		log.Fatal(err)
	}

	// TestService.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
	// 	w.Write([]byte("hello world"))
	// })
	// 启动http服务
	// httpService.Run()
}
