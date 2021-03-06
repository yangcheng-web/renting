// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetImageCd.proto

package GetImageCd

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

// Api Endpoints for GetImageCd service

func NewGetImageCdEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GetImageCd service

type GetImageCdService interface {
	GetImageCd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetImageCd_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (GetImageCd_PingPongService, error)
}

type getImageCdService struct {
	c    client.Client
	name string
}

func NewGetImageCdService(name string, c client.Client) GetImageCdService {
	return &getImageCdService{
		c:    c,
		name: name,
	}
}

func (c *getImageCdService) GetImageCd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetImageCd.GetImageCd", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getImageCdService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetImageCd_StreamService, error) {
	req := c.c.NewRequest(c.name, "GetImageCd.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &getImageCdServiceStream{stream}, nil
}

type GetImageCd_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type getImageCdServiceStream struct {
	stream client.Stream
}

func (x *getImageCdServiceStream) Close() error {
	return x.stream.Close()
}

func (x *getImageCdServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getImageCdServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getImageCdServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getImageCdServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getImageCdService) PingPong(ctx context.Context, opts ...client.CallOption) (GetImageCd_PingPongService, error) {
	req := c.c.NewRequest(c.name, "GetImageCd.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &getImageCdServicePingPong{stream}, nil
}

type GetImageCd_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type getImageCdServicePingPong struct {
	stream client.Stream
}

func (x *getImageCdServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *getImageCdServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *getImageCdServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getImageCdServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getImageCdServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *getImageCdServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GetImageCd service

type GetImageCdHandler interface {
	GetImageCd(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, GetImageCd_StreamStream) error
	PingPong(context.Context, GetImageCd_PingPongStream) error
}

func RegisterGetImageCdHandler(s server.Server, hdlr GetImageCdHandler, opts ...server.HandlerOption) error {
	type getImageCd interface {
		GetImageCd(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type GetImageCd struct {
		getImageCd
	}
	h := &getImageCdHandler{hdlr}
	return s.Handle(s.NewHandler(&GetImageCd{h}, opts...))
}

type getImageCdHandler struct {
	GetImageCdHandler
}

func (h *getImageCdHandler) GetImageCd(ctx context.Context, in *Request, out *Response) error {
	return h.GetImageCdHandler.GetImageCd(ctx, in, out)
}

func (h *getImageCdHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GetImageCdHandler.Stream(ctx, m, &getImageCdStreamStream{stream})
}

type GetImageCd_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type getImageCdStreamStream struct {
	stream server.Stream
}

func (x *getImageCdStreamStream) Close() error {
	return x.stream.Close()
}

func (x *getImageCdStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getImageCdStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getImageCdStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getImageCdStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *getImageCdHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GetImageCdHandler.PingPong(ctx, &getImageCdPingPongStream{stream})
}

type GetImageCd_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type getImageCdPingPongStream struct {
	stream server.Stream
}

func (x *getImageCdPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *getImageCdPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getImageCdPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getImageCdPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getImageCdPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *getImageCdPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
