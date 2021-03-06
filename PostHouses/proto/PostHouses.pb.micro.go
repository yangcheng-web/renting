// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/PostHouses.proto

package PostHouses

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

// Api Endpoints for PostHouses service

func NewPostHousesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PostHouses service

type PostHousesService interface {
	PostHouses(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (PostHouses_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (PostHouses_PingPongService, error)
}

type postHousesService struct {
	c    client.Client
	name string
}

func NewPostHousesService(name string, c client.Client) PostHousesService {
	return &postHousesService{
		c:    c,
		name: name,
	}
}

func (c *postHousesService) PostHouses(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PostHouses.PostHouses", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postHousesService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (PostHouses_StreamService, error) {
	req := c.c.NewRequest(c.name, "PostHouses.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &postHousesServiceStream{stream}, nil
}

type PostHouses_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type postHousesServiceStream struct {
	stream client.Stream
}

func (x *postHousesServiceStream) Close() error {
	return x.stream.Close()
}

func (x *postHousesServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postHousesServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postHousesServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postHousesServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *postHousesService) PingPong(ctx context.Context, opts ...client.CallOption) (PostHouses_PingPongService, error) {
	req := c.c.NewRequest(c.name, "PostHouses.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &postHousesServicePingPong{stream}, nil
}

type PostHouses_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type postHousesServicePingPong struct {
	stream client.Stream
}

func (x *postHousesServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *postHousesServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *postHousesServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postHousesServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postHousesServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *postHousesServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PostHouses service

type PostHousesHandler interface {
	PostHouses(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, PostHouses_StreamStream) error
	PingPong(context.Context, PostHouses_PingPongStream) error
}

func RegisterPostHousesHandler(s server.Server, hdlr PostHousesHandler, opts ...server.HandlerOption) error {
	type postHouses interface {
		PostHouses(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type PostHouses struct {
		postHouses
	}
	h := &postHousesHandler{hdlr}
	return s.Handle(s.NewHandler(&PostHouses{h}, opts...))
}

type postHousesHandler struct {
	PostHousesHandler
}

func (h *postHousesHandler) PostHouses(ctx context.Context, in *Request, out *Response) error {
	return h.PostHousesHandler.PostHouses(ctx, in, out)
}

func (h *postHousesHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.PostHousesHandler.Stream(ctx, m, &postHousesStreamStream{stream})
}

type PostHouses_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type postHousesStreamStream struct {
	stream server.Stream
}

func (x *postHousesStreamStream) Close() error {
	return x.stream.Close()
}

func (x *postHousesStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postHousesStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postHousesStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postHousesStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *postHousesHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.PostHousesHandler.PingPong(ctx, &postHousesPingPongStream{stream})
}

type PostHouses_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type postHousesPingPongStream struct {
	stream server.Stream
}

func (x *postHousesPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *postHousesPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postHousesPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postHousesPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postHousesPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *postHousesPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
