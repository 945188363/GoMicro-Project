package Test

import (
	"GoMicro-Project/LoadBalancer"
	"GoMicro-Project/Service/Micros"
	"GoMicro-Project/Service/Model"
	"GoMicro-Project/Wrapper"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myHttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

// 直接根据地址调用http服务
func CallHttpAPI(addr string, path string, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

// 拉去服务注册中心服务调用http服务
func CallHttpAPI2(path string, method string) (string, error) {
	// 获取consul注册地址
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	// 获取服务
	ProdServices, err := consulReg.GetService("ProdServiceHttp")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	// 轮询选择服务
	next := selector.RoundRobin(ProdServices)
	node, err := next()
	// 访问服务
	req, _ := http.NewRequest(method, "http://"+node.Address+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

// 使用micro插件调用http服务
func CallHttpAPI3(s selector.Selector, path string, method string) {
	myClient := myHttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)

	// req := myClient.NewRequest("ProdService","v1/prods",map[string]string{})
	req := myClient.NewRequest("ProdServiceHttp", "v1/prods", Model.ProdRequest{Size: 3})
	// var resp map[string]interface{}
	var resp Model.ProdResponse
	err := myClient.Call(context.Background(), req, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Data)
}

// 调用rpc服务
func CallRpcAPI() (string, error) {
	// 获取consul注册地址
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	// 获取服务
	prodServiceClient := micro.NewService(
		micro.Name("ProdService.client"),
		micro.WrapClient(Wrapper.NewLogWrapper),
		micro.Registry(consulReg),
	)
	prodServiceClient.Init()
	prodService1 := Micros.NewProdService1Service("ProdServiceRPC", prodServiceClient.Client())
	var req Model.ProdRequest1
	req.Size = 2
	prodResponse1, err := prodService1.GetProdList(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return prodResponse1.String(), nil
}

// 使用micro插件调用rpc服务
func CallRpcAPI2(s selector.Selector) {
	// 获取服务
	prodServiceClient := micro.NewService(
		micro.Name("ProdService.client"),
		micro.Selector(s),
	)
	prodService1 := Micros.NewProdService1Service("ProdServiceRPC", prodServiceClient.Client())
	var req Model.ProdRequest1
	req.Size = 2
	prodResponse1, err := prodService1.GetProdList(context.Background(), &req)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(prodResponse1.Data)
}

func TestAPI(t *testing.T) {
	res, err := CallHttpAPI("localhost:8888", "/v1/prods", "GET")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}

func TestAPI2(t *testing.T) {
	res, err := CallHttpAPI2("/v1/prods", "GET")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}

func TestAPI3(t *testing.T) {
	// 获取consul注册地址
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.Random),
	)
	CallHttpAPI3(mySelector, "/v1/prods", "GET")
}

func TestAPI4(t *testing.T) {
	CallRpcAPI()
}

func TestAPI5(t *testing.T) {
	// 获取consul注册地址
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.Random),
	)
	CallRpcAPI2(mySelector)
}

func TestLB(t *testing.T) {
	LoadBalancer.Test()
}
