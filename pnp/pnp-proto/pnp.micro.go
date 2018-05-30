// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pnp.proto

/*
Package pnp is a generated protocol buffer package.

It is generated from these files:
	pnp.proto

It has these top-level messages:
	ClientInfo
	CommonClientInfo
	ClientPkgRequest
	CommonServerResponse
	ServerInstructionPayload
	ServerPkgResponse
*/
package pnp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/ZTP/pnp/common/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PnP service

type PnPService interface {
	GetPackages(ctx context.Context, opts ...client.CallOption) (PnP_GetPackagesService, error)
}

type pnPService struct {
	c           client.Client
	serviceName string
}

func PnPServiceClient(serviceName string, c client.Client) PnPService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pnp"
	}
	return &pnPService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *pnPService) GetPackages(ctx context.Context, opts ...client.CallOption) (PnP_GetPackagesService, error) {
	req := c.c.NewRequest(c.serviceName, "PnP.GetPackages", &ClientPkgRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &pnPGetPackagesService{stream}, nil
}

type PnP_GetPackagesService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ClientPkgRequest) error
	Recv() (*ServerPkgResponse, error)
}

type pnPGetPackagesService struct {
	stream client.Stream
}

func (x *pnPGetPackagesService) Close() error {
	return x.stream.Close()
}

func (x *pnPGetPackagesService) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *pnPGetPackagesService) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *pnPGetPackagesService) Send(m *ClientPkgRequest) error {
	return x.stream.Send(m)
}

func (x *pnPGetPackagesService) Recv() (*ServerPkgResponse, error) {
	m := new(ServerPkgResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PnP service

type PnPHandler interface {
	GetPackages(context.Context, PnP_GetPackagesStream) error
}

func RegisterPnPHandler(s server.Server, hdlr PnPHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&PnP{hdlr}, opts...))
}

type PnP struct {
	PnPHandler
}

func (h *PnP) GetPackages(ctx context.Context, stream server.Stream) error {
	return h.PnPHandler.GetPackages(ctx, &pnPGetPackagesStream{stream})
}

type PnP_GetPackagesStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerPkgResponse) error
	Recv() (*ClientPkgRequest, error)
}

type pnPGetPackagesStream struct {
	stream server.Stream
}

func (x *pnPGetPackagesStream) Close() error {
	return x.stream.Close()
}

func (x *pnPGetPackagesStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *pnPGetPackagesStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *pnPGetPackagesStream) Send(m *ServerPkgResponse) error {
	return x.stream.Send(m)
}

func (x *pnPGetPackagesStream) Recv() (*ClientPkgRequest, error) {
	m := new(ClientPkgRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
