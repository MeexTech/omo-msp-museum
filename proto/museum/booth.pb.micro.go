// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/museum/booth.proto

package omo_msp_museum

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Client API for BoothService service

type BoothService interface {
	AddOne(ctx context.Context, in *ReqBoothAdd, opts ...client.CallOption) (*ReplyBoothInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyBoothInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyBoothList, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyBoothList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateBase(ctx context.Context, in *ReqBoothUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
}

type boothService struct {
	c    client.Client
	name string
}

func NewBoothService(name string, c client.Client) BoothService {
	return &boothService{
		c:    c,
		name: name,
	}
}

func (c *boothService) AddOne(ctx context.Context, in *ReqBoothAdd, opts ...client.CallOption) (*ReplyBoothInfo, error) {
	req := c.c.NewRequest(c.name, "BoothService.AddOne", in)
	out := new(ReplyBoothInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boothService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyBoothInfo, error) {
	req := c.c.NewRequest(c.name, "BoothService.GetOne", in)
	out := new(ReplyBoothInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boothService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "BoothService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boothService) Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyBoothList, error) {
	req := c.c.NewRequest(c.name, "BoothService.Search", in)
	out := new(ReplyBoothList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boothService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyBoothList, error) {
	req := c.c.NewRequest(c.name, "BoothService.GetListByFilter", in)
	out := new(ReplyBoothList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boothService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "BoothService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boothService) UpdateBase(ctx context.Context, in *ReqBoothUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "BoothService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boothService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "BoothService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BoothService service

type BoothServiceHandler interface {
	AddOne(context.Context, *ReqBoothAdd, *ReplyBoothInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyBoothInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	Search(context.Context, *RequestInfo, *ReplyBoothList) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyBoothList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateBase(context.Context, *ReqBoothUpdate, *ReplyInfo) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
}

func RegisterBoothServiceHandler(s server.Server, hdlr BoothServiceHandler, opts ...server.HandlerOption) error {
	type boothService interface {
		AddOne(ctx context.Context, in *ReqBoothAdd, out *ReplyBoothInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyBoothInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		Search(ctx context.Context, in *RequestInfo, out *ReplyBoothList) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyBoothList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateBase(ctx context.Context, in *ReqBoothUpdate, out *ReplyInfo) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
	}
	type BoothService struct {
		boothService
	}
	h := &boothServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BoothService{h}, opts...))
}

type boothServiceHandler struct {
	BoothServiceHandler
}

func (h *boothServiceHandler) AddOne(ctx context.Context, in *ReqBoothAdd, out *ReplyBoothInfo) error {
	return h.BoothServiceHandler.AddOne(ctx, in, out)
}

func (h *boothServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyBoothInfo) error {
	return h.BoothServiceHandler.GetOne(ctx, in, out)
}

func (h *boothServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.BoothServiceHandler.RemoveOne(ctx, in, out)
}

func (h *boothServiceHandler) Search(ctx context.Context, in *RequestInfo, out *ReplyBoothList) error {
	return h.BoothServiceHandler.Search(ctx, in, out)
}

func (h *boothServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyBoothList) error {
	return h.BoothServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *boothServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.BoothServiceHandler.GetStatistic(ctx, in, out)
}

func (h *boothServiceHandler) UpdateBase(ctx context.Context, in *ReqBoothUpdate, out *ReplyInfo) error {
	return h.BoothServiceHandler.UpdateBase(ctx, in, out)
}

func (h *boothServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.BoothServiceHandler.UpdateByFilter(ctx, in, out)
}
