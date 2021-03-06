// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetUserHouses.proto

package GetUserHouses

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

// Api Endpoints for GetUserHouses service

func NewGetUserHousesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GetUserHouses service

type GetUserHousesService interface {
	GetUserHouses(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetUserHouses_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (GetUserHouses_PingPongService, error)
}

type getUserHousesService struct {
	c    client.Client
	name string
}

func NewGetUserHousesService(name string, c client.Client) GetUserHousesService {
	return &getUserHousesService{
		c:    c,
		name: name,
	}
}

func (c *getUserHousesService) GetUserHouses(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetUserHouses.GetUserHouses", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getUserHousesService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetUserHouses_StreamService, error) {
	req := c.c.NewRequest(c.name, "GetUserHouses.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &getUserHousesServiceStream{stream}, nil
}

type GetUserHouses_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type getUserHousesServiceStream struct {
	stream client.Stream
}

func (x *getUserHousesServiceStream) Close() error {
	return x.stream.Close()
}

func (x *getUserHousesServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getUserHousesServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getUserHousesServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getUserHousesServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getUserHousesService) PingPong(ctx context.Context, opts ...client.CallOption) (GetUserHouses_PingPongService, error) {
	req := c.c.NewRequest(c.name, "GetUserHouses.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &getUserHousesServicePingPong{stream}, nil
}

type GetUserHouses_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type getUserHousesServicePingPong struct {
	stream client.Stream
}

func (x *getUserHousesServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *getUserHousesServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *getUserHousesServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getUserHousesServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getUserHousesServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *getUserHousesServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GetUserHouses service

type GetUserHousesHandler interface {
	GetUserHouses(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, GetUserHouses_StreamStream) error
	PingPong(context.Context, GetUserHouses_PingPongStream) error
}

func RegisterGetUserHousesHandler(s server.Server, hdlr GetUserHousesHandler, opts ...server.HandlerOption) error {
	type getUserHouses interface {
		GetUserHouses(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type GetUserHouses struct {
		getUserHouses
	}
	h := &getUserHousesHandler{hdlr}
	return s.Handle(s.NewHandler(&GetUserHouses{h}, opts...))
}

type getUserHousesHandler struct {
	GetUserHousesHandler
}

func (h *getUserHousesHandler) GetUserHouses(ctx context.Context, in *Request, out *Response) error {
	return h.GetUserHousesHandler.GetUserHouses(ctx, in, out)
}

func (h *getUserHousesHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GetUserHousesHandler.Stream(ctx, m, &getUserHousesStreamStream{stream})
}

type GetUserHouses_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type getUserHousesStreamStream struct {
	stream server.Stream
}

func (x *getUserHousesStreamStream) Close() error {
	return x.stream.Close()
}

func (x *getUserHousesStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getUserHousesStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getUserHousesStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getUserHousesStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *getUserHousesHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GetUserHousesHandler.PingPong(ctx, &getUserHousesPingPongStream{stream})
}

type GetUserHouses_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type getUserHousesPingPongStream struct {
	stream server.Stream
}

func (x *getUserHousesPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *getUserHousesPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getUserHousesPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getUserHousesPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getUserHousesPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *getUserHousesPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
