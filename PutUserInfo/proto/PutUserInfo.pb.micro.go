// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/PutUserInfo.proto

package PutUserInfo

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

// Api Endpoints for PutUserInfo service

func NewPutUserInfoEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PutUserInfo service

type PutUserInfoService interface {
	PutUserInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (PutUserInfo_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (PutUserInfo_PingPongService, error)
}

type putUserInfoService struct {
	c    client.Client
	name string
}

func NewPutUserInfoService(name string, c client.Client) PutUserInfoService {
	return &putUserInfoService{
		c:    c,
		name: name,
	}
}

func (c *putUserInfoService) PutUserInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PutUserInfo.PutUserInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putUserInfoService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (PutUserInfo_StreamService, error) {
	req := c.c.NewRequest(c.name, "PutUserInfo.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &putUserInfoServiceStream{stream}, nil
}

type PutUserInfo_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type putUserInfoServiceStream struct {
	stream client.Stream
}

func (x *putUserInfoServiceStream) Close() error {
	return x.stream.Close()
}

func (x *putUserInfoServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *putUserInfoServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *putUserInfoServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *putUserInfoServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *putUserInfoService) PingPong(ctx context.Context, opts ...client.CallOption) (PutUserInfo_PingPongService, error) {
	req := c.c.NewRequest(c.name, "PutUserInfo.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &putUserInfoServicePingPong{stream}, nil
}

type PutUserInfo_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type putUserInfoServicePingPong struct {
	stream client.Stream
}

func (x *putUserInfoServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *putUserInfoServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *putUserInfoServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *putUserInfoServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *putUserInfoServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *putUserInfoServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PutUserInfo service

type PutUserInfoHandler interface {
	PutUserInfo(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, PutUserInfo_StreamStream) error
	PingPong(context.Context, PutUserInfo_PingPongStream) error
}

func RegisterPutUserInfoHandler(s server.Server, hdlr PutUserInfoHandler, opts ...server.HandlerOption) error {
	type putUserInfo interface {
		PutUserInfo(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type PutUserInfo struct {
		putUserInfo
	}
	h := &putUserInfoHandler{hdlr}
	return s.Handle(s.NewHandler(&PutUserInfo{h}, opts...))
}

type putUserInfoHandler struct {
	PutUserInfoHandler
}

func (h *putUserInfoHandler) PutUserInfo(ctx context.Context, in *Request, out *Response) error {
	return h.PutUserInfoHandler.PutUserInfo(ctx, in, out)
}

func (h *putUserInfoHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.PutUserInfoHandler.Stream(ctx, m, &putUserInfoStreamStream{stream})
}

type PutUserInfo_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type putUserInfoStreamStream struct {
	stream server.Stream
}

func (x *putUserInfoStreamStream) Close() error {
	return x.stream.Close()
}

func (x *putUserInfoStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *putUserInfoStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *putUserInfoStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *putUserInfoStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *putUserInfoHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.PutUserInfoHandler.PingPong(ctx, &putUserInfoPingPongStream{stream})
}

type PutUserInfo_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type putUserInfoPingPongStream struct {
	stream server.Stream
}

func (x *putUserInfoPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *putUserInfoPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *putUserInfoPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *putUserInfoPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *putUserInfoPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *putUserInfoPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
