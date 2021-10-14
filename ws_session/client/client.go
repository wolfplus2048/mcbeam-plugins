package client

import (
	"github.com/golang/protobuf/proto"
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/context"
	"github.com/micro/micro/v3/service/logger"
	"github.com/wolfplus2048/mcbeam-plugins/ws_session/v3"
	pb "github.com/wolfplus2048/mcbeam-plugins/ws_session/v3/proto"
)

type srv struct {
	opts session.Options
	session pb.SessionService
}

func (s *srv) Init(opts ...session.Option) {
	for _, o := range opts {
		o(&s.opts)
	}
}
func (s *srv) Options() session.Options {
	return s.opts
}
func (s *srv) String() string {
	return "session"
}
func (s *srv) SessionID() string {
	return s.opts.SessionID
}

func (s *srv) Kick() error {
	_, err := s.session.Kick(context.DefaultContext, &pb.KickRequest{Sid: s.opts.SessionID}, client.WithServerUid(s.opts.ServerID), client.WithAuthToken())
	if err != nil {
		logger.Infof("session kick, err:%v", err)
	}
	return err
}

func (s *srv) Send(route string, v interface{}) error {
	b, err := proto.Marshal(v.(proto.Message))
	if err != nil {
		return err
	}
	req := &pb.Message{
		Sid:   s.opts.SessionID,
		Route: route,
		Body:  b,
	}
	_, err = s.session.Send(context.DefaultContext, req, client.WithServerUid(s.opts.ServerID), client.WithAuthToken())
	if err != nil {
		logger.Infof("session send, err:%v", err)
	}
	return err
}

func NewSession(opts ...session.Option) session.Session {
	s := &srv{
		session: pb.NewSessionService("websocket", client.DefaultClient),
	}
	for _, o := range opts {
		o(&s.opts)
	}
	return s
}
