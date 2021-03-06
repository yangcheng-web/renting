// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/PostHousesImage.proto

package PostHousesImage

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

// Api Endpoints for PostHousesImage service

func NewPostHousesImageEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PostHousesImage service

type PostHousesImageService interface {
	PostHousesImage(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (PostHousesImage_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (PostHousesImage_PingPongService, error)
}

type postHousesImageService struct {
	c    client.Client
	name string
}

func NewPostHousesImageService(name string, c client.Client) PostHousesImageService {
	return &postHousesImageService{
		c:    c,
		name: name,
	}
}

func (c *postHousesImageService) PostHousesImage(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PostHousesImage.PostHousesImage", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postHousesImageService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (PostHousesImage_StreamService, error) {
	req := c.c.NewRequest(c.name, "PostHousesImage.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &postHousesImageServiceStream{stream}, nil
}

type PostHousesImage_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type postHousesImageServiceStream struct {
	stream client.Stream
}

func (x *postHousesImageServiceStream) Close() error {
	return x.stream.Close()
}

func (x *postHousesImageServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postHousesImageServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postHousesImageServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postHousesImageServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *postHousesImageService) PingPong(ctx context.Context, opts ...client.CallOption) (PostHousesImage_PingPongService, error) {
	req := c.c.NewRequest(c.name, "PostHousesImage.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &postHousesImageServicePingPong{stream}, nil
}

type PostHousesImage_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type postHousesImageServicePingPong struct {
	stream client.Stream
}

func (x *postHousesImageServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *postHousesImageServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *postHousesImageServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postHousesImageServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postHousesImageServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *postHousesImageServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PostHousesImage service

type PostHousesImageHandler interface {
	PostHousesImage(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, PostHousesImage_StreamStream) error
	PingPong(context.Context, PostHousesImage_PingPongStream) error
}

func RegisterPostHousesImageHandler(s server.Server, hdlr PostHousesImageHandler, opts ...server.HandlerOption) error {
	type postHousesImage interface {
		PostHousesImage(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type PostHousesImage struct {
		postHousesImage
	}
	h := &postHousesImageHandler{hdlr}
	return s.Handle(s.NewHandler(&PostHousesImage{h}, opts...))
}

type postHousesImageHandler struct {
	PostHousesImageHandler
}

func (h *postHousesImageHandler) PostHousesImage(ctx context.Context, in *Request, out *Response) error {
	return h.PostHousesImageHandler.PostHousesImage(ctx, in, out)
}

func (h *postHousesImageHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.PostHousesImageHandler.Stream(ctx, m, &postHousesImageStreamStream{stream})
}

type PostHousesImage_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type postHousesImageStreamStream struct {
	stream server.Stream
}

func (x *postHousesImageStreamStream) Close() error {
	return x.stream.Close()
}

func (x *postHousesImageStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postHousesImageStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postHousesImageStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postHousesImageStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *postHousesImageHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.PostHousesImageHandler.PingPong(ctx, &postHousesImagePingPongStream{stream})
}

type PostHousesImage_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type postHousesImagePingPongStream struct {
	stream server.Stream
}

func (x *postHousesImagePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *postHousesImagePingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postHousesImagePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postHousesImagePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postHousesImagePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *postHousesImagePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
