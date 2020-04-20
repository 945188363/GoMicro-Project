package main

import (
	"GoMicro-Project/Service/ProdService"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"math/rand"
	"sort"
	"time"
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

// func main()  {
// 	Test()
// }

type IpTable struct {
	Ip             string
	Port           int
	Weight         int
	CurrentWeight  int
	CurrentConnect int
}

func NewIpTable(ip string, port int, weight int, currentWeight int, connect int) *IpTable {
	return &IpTable{
		Ip:             ip,
		Port:           port,
		Weight:         weight,
		CurrentWeight:  currentWeight,
		CurrentConnect: connect,
	}
}

func Test() {
	ipTables := make([]*IpTable, 0)
	ipTables = append(ipTables, NewIpTable("192.168.1.1", 8888, 1, 0, 0))
	ipTables = append(ipTables, NewIpTable("192.168.1.2", 8888, 2, 0, 0))
	ipTables = append(ipTables, NewIpTable("192.168.1.3", 8888, 7, 0, 0))
	for {
		fmt.Println(WeightedLeastConnection(ipTables))
	}
}

// 标记轮询的地址，使用当前轮询地址
var roundRobinIndex = 0

// 轮询
func RoundRobin(ips []*IpTable) *IpTable {
	roundRobinIndex++
	if roundRobinIndex >= len(ips) {
		roundRobinIndex = roundRobinIndex % len(ips)
	}
	return ips[roundRobinIndex]
}

// 标记轮询的地址，使用当前轮询地址
var weightRoundRobinIndex = 0

// 加权轮询
func WeightedRoundRobin(ips []*IpTable) *IpTable {
	var totalWeight = 0
	for _, Ip := range ips {
		totalWeight += Ip.Weight
	}
	nums := weightRoundRobinIndex % totalWeight
	weightRoundRobinIndex++

	// 根据ip切片顺序来排权重区间
	for _, Ip := range ips {
		// 落在当前区间返回
		if Ip.Weight >= nums {
			return Ip
		}
		// 如果未落在当前区间，则减掉当前区间
		nums -= Ip.Weight
	}
	return nil

	// var index = -1
	// var totalWeight = 0
	//
	// for i,Ip := range ips {
	// 	Ip.CurrentWeight += Ip.Weight
	// 	totalWeight += Ip.Weight
	//
	// 	if index == -1 || ips[index].CurrentWeight < Ip.CurrentWeight {
	// 		index = i
	// 	}
	// }
	//
	// ips[index].CurrentWeight -= totalWeight
	// return ips[index]
}

// 完全随机
func Random(ips []*IpTable) *IpTable {
	// 使用系统时间作为种子
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(ips))
	return ips[index]
}

// 加权随机
func WeightedRandom(ips []*IpTable) *IpTable {
	var totalWeight = 0
	for _, Ip := range ips {
		totalWeight += Ip.Weight
	}
	rand.Seed(time.Now().UnixNano())
	nums := rand.Intn(totalWeight)
	// 根据ip切片顺序来排权重区间
	for _, Ip := range ips {
		// 落在当前区间返回
		if Ip.Weight >= nums {
			return Ip
		}
		// 如果未落在当前区间，则减掉当前区间
		nums -= Ip.Weight
	}
	return nil
}

func Hash(ips []*IpTable) {

}

// 最少连接
func LeastConnection(ips []*IpTable) *IpTable {
	// 根据连接数升序排序
	sort.Slice(ips, func(i, j int) bool {
		return ips[i].CurrentConnect < ips[j].CurrentConnect
	})
	// 连接加一并返回当前最少连接的ip
	ips[0].CurrentConnect++
	return ips[0]
}

// 加权最少连接
func WeightedLeastConnection(ips []*IpTable) *IpTable {
	var index = 0
	for i, Ip := range ips {
		// 优先级相关性:连接负相关，权重正相关 if W(i)/C(i) > W(j)/C(j)  return ip[i]
		if Ip.Weight*ips[index].CurrentConnect >= Ip.CurrentConnect*ips[index].Weight {
			index = i
		}
	}
	// 连接加一并返回当前最少连接的ip
	ips[index].CurrentConnect++
	return ips[index]
}

func LeastResponseTime(ips []*IpTable) {

}
