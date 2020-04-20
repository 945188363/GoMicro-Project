package Test

import (
	"GoMicro-Project/LoadBalancer"
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
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
		registry.Addrs("192.168.2.1"),
	)
	// 获取服务
	ProdServices, err := consulReg.GetService("ProdService")
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

func TestLB(t *testing.T) {
	LoadBalancer.Test()
}
