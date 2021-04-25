// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetHouseInfo.proto

package GetHouseInfo

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

// Api Endpoints for GetHouseInfo service

func NewGetHouseInfoEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GetHouseInfo service

type GetHouseInfoService interface {
	GetHouseInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetHouseInfo_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (GetHouseInfo_PingPongService, error)
}

type getHouseInfoService struct {
	c    client.Client
	name string
}

func NewGetHouseInfoService(name string, c client.Client) GetHouseInfoService {
	return &getHouseInfoService{
		c:    c,
		name: name,
	}
}

func (c *getHouseInfoService) GetHouseInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetHouseInfo.GetHouseInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getHouseInfoService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetHouseInfo_StreamService, error) {
	req := c.c.NewRequest(c.name, "GetHouseInfo.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &getHouseInfoServiceStream{stream}, nil
}

type GetHouseInfo_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type getHouseInfoServiceStream struct {
	stream client.Stream
}

func (x *getHouseInfoServiceStream) Close() error {
	return x.stream.Close()
}

func (x *getHouseInfoServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getHouseInfoServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getHouseInfoServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getHouseInfoServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getHouseInfoService) PingPong(ctx context.Context, opts ...client.CallOption) (GetHouseInfo_PingPongService, error) {
	req := c.c.NewRequest(c.name, "GetHouseInfo.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &getHouseInfoServicePingPong{stream}, nil
}

type GetHouseInfo_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type getHouseInfoServicePingPong struct {
	stream client.Stream
}

func (x *getHouseInfoServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *getHouseInfoServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *getHouseInfoServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getHouseInfoServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getHouseInfoServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *getHouseInfoServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GetHouseInfo service

type GetHouseInfoHandler interface {
	GetHouseInfo(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, GetHouseInfo_StreamStream) error
	PingPong(context.Context, GetHouseInfo_PingPongStream) error
}

func RegisterGetHouseInfoHandler(s server.Server, hdlr GetHouseInfoHandler, opts ...server.HandlerOption) error {
	type getHouseInfo interface {
		GetHouseInfo(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type GetHouseInfo struct {
		getHouseInfo
	}
	h := &getHouseInfoHandler{hdlr}
	return s.Handle(s.NewHandler(&GetHouseInfo{h}, opts...))
}

type getHouseInfoHandler struct {
	GetHouseInfoHandler
}

func (h *getHouseInfoHandler) GetHouseInfo(ctx context.Context, in *Request, out *Response) error {
	return h.GetHouseInfoHandler.GetHouseInfo(ctx, in, out)
}

func (h *getHouseInfoHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GetHouseInfoHandler.Stream(ctx, m, &getHouseInfoStreamStream{stream})
}

type GetHouseInfo_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type getHouseInfoStreamStream struct {
	stream server.Stream
}

func (x *getHouseInfoStreamStream) Close() error {
	return x.stream.Close()
}

func (x *getHouseInfoStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getHouseInfoStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getHouseInfoStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getHouseInfoStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *getHouseInfoHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GetHouseInfoHandler.PingPong(ctx, &getHouseInfoPingPongStream{stream})
}

type GetHouseInfo_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type getHouseInfoPingPongStream struct {
	stream server.Stream
}

func (x *getHouseInfoPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *getHouseInfoPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getHouseInfoPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getHouseInfoPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getHouseInfoPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *getHouseInfoPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
