// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetSession.proto

package GetSession

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

// Api Endpoints for GetSession service

func NewGetSessionEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GetSession service

type GetSessionService interface {
	GetSession(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetSession_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (GetSession_PingPongService, error)
}

type getSessionService struct {
	c    client.Client
	name string
}

func NewGetSessionService(name string, c client.Client) GetSessionService {
	return &getSessionService{
		c:    c,
		name: name,
	}
}

func (c *getSessionService) GetSession(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetSession.GetSession", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getSessionService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetSession_StreamService, error) {
	req := c.c.NewRequest(c.name, "GetSession.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &getSessionServiceStream{stream}, nil
}

type GetSession_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type getSessionServiceStream struct {
	stream client.Stream
}

func (x *getSessionServiceStream) Close() error {
	return x.stream.Close()
}

func (x *getSessionServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getSessionServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSessionServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSessionServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getSessionService) PingPong(ctx context.Context, opts ...client.CallOption) (GetSession_PingPongService, error) {
	req := c.c.NewRequest(c.name, "GetSession.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &getSessionServicePingPong{stream}, nil
}

type GetSession_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type getSessionServicePingPong struct {
	stream client.Stream
}

func (x *getSessionServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *getSessionServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *getSessionServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSessionServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSessionServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *getSessionServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GetSession service

type GetSessionHandler interface {
	GetSession(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, GetSession_StreamStream) error
	PingPong(context.Context, GetSession_PingPongStream) error
}

func RegisterGetSessionHandler(s server.Server, hdlr GetSessionHandler, opts ...server.HandlerOption) error {
	type getSession interface {
		GetSession(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type GetSession struct {
		getSession
	}
	h := &getSessionHandler{hdlr}
	return s.Handle(s.NewHandler(&GetSession{h}, opts...))
}

type getSessionHandler struct {
	GetSessionHandler
}

func (h *getSessionHandler) GetSession(ctx context.Context, in *Request, out *Response) error {
	return h.GetSessionHandler.GetSession(ctx, in, out)
}

func (h *getSessionHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GetSessionHandler.Stream(ctx, m, &getSessionStreamStream{stream})
}

type GetSession_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type getSessionStreamStream struct {
	stream server.Stream
}

func (x *getSessionStreamStream) Close() error {
	return x.stream.Close()
}

func (x *getSessionStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getSessionStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSessionStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSessionStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *getSessionHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GetSessionHandler.PingPong(ctx, &getSessionPingPongStream{stream})
}

type GetSession_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type getSessionPingPongStream struct {
	stream server.Stream
}

func (x *getSessionPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *getSessionPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *getSessionPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSessionPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSessionPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *getSessionPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
