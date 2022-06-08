// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/museum/anchor.proto

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

// Client API for AnchorService service

type AnchorService interface {
	AddOne(ctx context.Context, in *ReqAnchorAdd, opts ...client.CallOption) (*ReplyAnchorOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAnchorOne, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAnchorList, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyAnchorList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateBase(ctx context.Context, in *ReqAnchorUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
}

type anchorService struct {
	c    client.Client
	name string
}

func NewAnchorService(name string, c client.Client) AnchorService {
	return &anchorService{
		c:    c,
		name: name,
	}
}

func (c *anchorService) AddOne(ctx context.Context, in *ReqAnchorAdd, opts ...client.CallOption) (*ReplyAnchorOne, error) {
	req := c.c.NewRequest(c.name, "AnchorService.AddOne", in)
	out := new(ReplyAnchorOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAnchorOne, error) {
	req := c.c.NewRequest(c.name, "AnchorService.GetOne", in)
	out := new(ReplyAnchorOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "AnchorService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorService) Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAnchorList, error) {
	req := c.c.NewRequest(c.name, "AnchorService.Search", in)
	out := new(ReplyAnchorList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyAnchorList, error) {
	req := c.c.NewRequest(c.name, "AnchorService.GetListByFilter", in)
	out := new(ReplyAnchorList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "AnchorService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorService) UpdateBase(ctx context.Context, in *ReqAnchorUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "AnchorService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anchorService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "AnchorService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AnchorService service

type AnchorServiceHandler interface {
	AddOne(context.Context, *ReqAnchorAdd, *ReplyAnchorOne) error
	GetOne(context.Context, *RequestInfo, *ReplyAnchorOne) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	Search(context.Context, *RequestInfo, *ReplyAnchorList) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyAnchorList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateBase(context.Context, *ReqAnchorUpdate, *ReplyInfo) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
}

func RegisterAnchorServiceHandler(s server.Server, hdlr AnchorServiceHandler, opts ...server.HandlerOption) error {
	type anchorService interface {
		AddOne(ctx context.Context, in *ReqAnchorAdd, out *ReplyAnchorOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyAnchorOne) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		Search(ctx context.Context, in *RequestInfo, out *ReplyAnchorList) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyAnchorList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateBase(ctx context.Context, in *ReqAnchorUpdate, out *ReplyInfo) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
	}
	type AnchorService struct {
		anchorService
	}
	h := &anchorServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AnchorService{h}, opts...))
}

type anchorServiceHandler struct {
	AnchorServiceHandler
}

func (h *anchorServiceHandler) AddOne(ctx context.Context, in *ReqAnchorAdd, out *ReplyAnchorOne) error {
	return h.AnchorServiceHandler.AddOne(ctx, in, out)
}

func (h *anchorServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyAnchorOne) error {
	return h.AnchorServiceHandler.GetOne(ctx, in, out)
}

func (h *anchorServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.AnchorServiceHandler.RemoveOne(ctx, in, out)
}

func (h *anchorServiceHandler) Search(ctx context.Context, in *RequestInfo, out *ReplyAnchorList) error {
	return h.AnchorServiceHandler.Search(ctx, in, out)
}

func (h *anchorServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyAnchorList) error {
	return h.AnchorServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *anchorServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.AnchorServiceHandler.GetStatistic(ctx, in, out)
}

func (h *anchorServiceHandler) UpdateBase(ctx context.Context, in *ReqAnchorUpdate, out *ReplyInfo) error {
	return h.AnchorServiceHandler.UpdateBase(ctx, in, out)
}

func (h *anchorServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.AnchorServiceHandler.UpdateByFilter(ctx, in, out)
}
