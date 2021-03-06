// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetUserInfo.proto

package GetUserInfo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for GetUserInfo service

func NewGetUserInfoEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GetUserInfo service

type GetUserInfoService interface {
	GetUserInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetUserInfo_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (GetUserInfo_PingPongService, error)
}

type getUserInfoService struct {
	c    client.Client
	name string
}

func NewGetUserInfoService(name string, c client.Client) GetUserInfoService {
	return &getUserInfoService{
		c:    c,
		name: name,
	}
}

func (c *getUserInfoService) GetUserInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetUserInfo.GetUserInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getUserInfoService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetUserInfo_StreamService, error) {
	req := c.c.NewRequest(c.name, "GetUserInfo.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &getUserInfoServiceStream{stream}, nil
}

type GetUserInfo_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type getUserInfoServiceStream struct {
	stream client.Stream
}

func (x *getUserInfoServiceStream) Close() error {
	return x.stream.Close()
}

func (x *getUserInfoServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getUserInfoServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getUserInfoServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getUserInfoServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getUserInfoService) PingPong(ctx context.Context, opts ...client.CallOption) (GetUserInfo_PingPongService, error) {
	req := c.c.NewRequest(c.name, "GetUserInfo.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &getUserInfoServicePingPong{stream}, nil
}

type GetUserInfo_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type getUserInfoServicePingPong struct {
	stream client.Stream
}

func (x *getUserInfoServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *getUserInfoServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *getUserInfoServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getUserInfoServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getUserInfoServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *getUserInfoServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GetUserInfo service

type GetUserInfoHandler interface {
	GetUserInfo(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, GetUserInfo_StreamStream) error
	PingPong(context.Context, GetUserInfo_PingPongStream) error
}

func RegisterGetUserInfoHandler(s server.Server, hdlr GetUserInfoHandler, opts ...server.HandlerOption) error {
	type getUserInfo interface {
		GetUserInfo(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type GetUserInfo struct {
		getUserInfo
	}
	h := &getUserInfoHandler{hdlr}
	return s.Handle(s.NewHandler(&GetUserInfo{h}, opts...))
}

type getUserInfoHandler struct {
	GetUserInfoHandler
}

func (h *getUserInfoHandler) GetUserInfo(ctx context.Context, in *Request, out *Response) error {
	return h.GetUserInfoHandler.GetUserInfo(ctx, in, out)
}

func (h *getUserInfoHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GetUserInfoHandler.Stream(ctx, m, &getUserInfoStreamStream{stream})
}

type GetUserInfo_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type getUserInfoStreamStream struct {
	stream server.Stream
}

func (x *getUserInfoStreamStream) Close() error {
	return x.stream.Close()
}

func (x *getUserInfoStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getUserInfoStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getUserInfoStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getUserInfoStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *getUserInfoHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GetUserInfoHandler.PingPong(ctx, &getUserInfoPingPongStream{stream})
}

type GetUserInfo_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type getUserInfoPingPongStream struct {
	stream server.Stream
}

func (x *getUserInfoPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *getUserInfoPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getUserInfoPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getUserInfoPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getUserInfoPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *getUserInfoPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
