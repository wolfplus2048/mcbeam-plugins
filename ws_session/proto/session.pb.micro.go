// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: ws_session/proto/session.proto

package session

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Api Endpoints for Session service

func NewSessionEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Session service

type SessionService interface {
	Send(ctx context.Context, in *Message, opts ...client.CallOption) (*EmptyResponse, error)
	Kick(ctx context.Context, in *KickRequest, opts ...client.CallOption) (*EmptyResponse, error)
	Bind(ctx context.Context, in *SessionStatus, opts ...client.CallOption) (*EmptyResponse, error)
}

type sessionService struct {
	c    client.Client
	name string
}

func NewSessionService(name string, c client.Client) SessionService {
	return &sessionService{
		c:    c,
		name: name,
	}
}

func (c *sessionService) Send(ctx context.Context, in *Message, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "Session.Send", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionService) Kick(ctx context.Context, in *KickRequest, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "Session.Kick", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionService) Bind(ctx context.Context, in *SessionStatus, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "Session.Bind", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Session service

type SessionHandler interface {
	Send(context.Context, *Message, *EmptyResponse) error
	Kick(context.Context, *KickRequest, *EmptyResponse) error
	Bind(context.Context, *SessionStatus, *EmptyResponse) error
}

func RegisterSessionHandler(s server.Server, hdlr SessionHandler, opts ...server.HandlerOption) error {
	type session interface {
		Send(ctx context.Context, in *Message, out *EmptyResponse) error
		Kick(ctx context.Context, in *KickRequest, out *EmptyResponse) error
		Bind(ctx context.Context, in *SessionStatus, out *EmptyResponse) error
	}
	type Session struct {
		session
	}
	h := &sessionHandler{hdlr}
	return s.Handle(s.NewHandler(&Session{h}, opts...))
}

type sessionHandler struct {
	SessionHandler
}

func (h *sessionHandler) Send(ctx context.Context, in *Message, out *EmptyResponse) error {
	return h.SessionHandler.Send(ctx, in, out)
}

func (h *sessionHandler) Kick(ctx context.Context, in *KickRequest, out *EmptyResponse) error {
	return h.SessionHandler.Kick(ctx, in, out)
}

func (h *sessionHandler) Bind(ctx context.Context, in *SessionStatus, out *EmptyResponse) error {
	return h.SessionHandler.Bind(ctx, in, out)
}
