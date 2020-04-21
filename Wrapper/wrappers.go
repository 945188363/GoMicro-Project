package Wrapper

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
)

type logWrapper struct {
	client.Client
}

func (this *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("日志测试")
	return this.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}
