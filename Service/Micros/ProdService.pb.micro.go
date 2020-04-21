// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: ProdService

package Micros

import (
	"GoMicro-Project/Service/Model"
	"fmt"
	"github.com/golang/protobuf/proto"
	"math"
)

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ProdService1 service

type ProdService1Service interface {
	GetProdList(ctx context.Context, in *Model.ProdRequest1, opts ...client.CallOption) (*Model.ProdResponse1, error)
}

type prodService1Service struct {
	c    client.Client
	name string
}

func NewProdService1Service(name string, c client.Client) ProdService1Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "Service"
	}
	return &prodService1Service{
		c:    c,
		name: name,
	}
}

func (c *prodService1Service) GetProdList(ctx context.Context, in *Model.ProdRequest1, opts ...client.CallOption) (*Model.ProdResponse1, error) {
	req := c.c.NewRequest(c.name, "ProdService1.GetProdList", in)
	out := new(Model.ProdResponse1)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProdService1 service

type ProdService1Handler interface {
	GetProdList(context.Context, *Model.ProdRequest1, *Model.ProdResponse1) error
}

func RegisterProdService1Handler(s server.Server, hdlr ProdService1Handler, opts ...server.HandlerOption) error {
	type prodService1 interface {
		GetProdList(ctx context.Context, in *Model.ProdRequest1, out *Model.ProdResponse1) error
	}
	type ProdService1 struct {
		prodService1
	}
	h := &prodService1Handler{hdlr}
	return s.Handle(s.NewHandler(&ProdService1{h}, opts...))
}

type prodService1Handler struct {
	ProdService1Handler
}

func (h *prodService1Handler) GetProdList(ctx context.Context, in *Model.ProdRequest1, out *Model.ProdResponse1) error {
	return h.ProdService1Handler.GetProdList(ctx, in, out)
}