// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetSmsCd.proto

package GetSmsCd

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

// Api Endpoints for GetSmsCd service

func NewGetSmsCdEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GetSmsCd service

type GetSmsCdService interface {
	GetSmsCd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetSmsCd_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (GetSmsCd_PingPongService, error)
}

type getSmsCdService struct {
	c    client.Client
	name string
}

func NewGetSmsCdService(name string, c client.Client) GetSmsCdService {
	return &getSmsCdService{
		c:    c,
		name: name,
	}
}

func (c *getSmsCdService) GetSmsCd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetSmsCd.GetSmsCd", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getSmsCdService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetSmsCd_StreamService, error) {
	req := c.c.NewRequest(c.name, "GetSmsCd.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &getSmsCdServiceStream{stream}, nil
}

type GetSmsCd_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type getSmsCdServiceStream struct {
	stream client.Stream
}

func (x *getSmsCdServiceStream) Close() error {
	return x.stream.Close()
}

func (x *getSmsCdServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getSmsCdServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSmsCdServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSmsCdServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getSmsCdService) PingPong(ctx context.Context, opts ...client.CallOption) (GetSmsCd_PingPongService, error) {
	req := c.c.NewRequest(c.name, "GetSmsCd.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &getSmsCdServicePingPong{stream}, nil
}

type GetSmsCd_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type getSmsCdServicePingPong struct {
	stream client.Stream
}

func (x *getSmsCdServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *getSmsCdServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *getSmsCdServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSmsCdServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSmsCdServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *getSmsCdServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GetSmsCd service

type GetSmsCdHandler interface {
	GetSmsCd(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, GetSmsCd_StreamStream) error
	PingPong(context.Context, GetSmsCd_PingPongStream) error
}

func RegisterGetSmsCdHandler(s server.Server, hdlr GetSmsCdHandler, opts ...server.HandlerOption) error {
	type getSmsCd interface {
		GetSmsCd(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type GetSmsCd struct {
		getSmsCd
	}
	h := &getSmsCdHandler{hdlr}
	return s.Handle(s.NewHandler(&GetSmsCd{h}, opts...))
}

type getSmsCdHandler struct {
	GetSmsCdHandler
}

func (h *getSmsCdHandler) GetSmsCd(ctx context.Context, in *Request, out *Response) error {
	return h.GetSmsCdHandler.GetSmsCd(ctx, in, out)
}

func (h *getSmsCdHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GetSmsCdHandler.Stream(ctx, m, &getSmsCdStreamStream{stream})
}

type GetSmsCd_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type getSmsCdStreamStream struct {
	stream server.Stream
}

func (x *getSmsCdStreamStream) Close() error {
	return x.stream.Close()
}

func (x *getSmsCdStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getSmsCdStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSmsCdStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSmsCdStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *getSmsCdHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GetSmsCdHandler.PingPong(ctx, &getSmsCdPingPongStream{stream})
}

type GetSmsCd_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type getSmsCdPingPongStream struct {
	stream server.Stream
}

func (x *getSmsCdPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *getSmsCdPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getSmsCdPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSmsCdPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSmsCdPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *getSmsCdPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
