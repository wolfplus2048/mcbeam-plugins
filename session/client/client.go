package client

import (
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/context"
	"github.com/micro/micro/v3/service/logger"
	"github.com/wolfplus2048/mcbeam-plugins/session/v3"
	"github.com/wolfplus2048/mcbeam-plugins/session/v3/codec/proto"
	pb "github.com/wolfplus2048/mcbeam-plugins/session/v3/proto"
)

type srv struct {
	opts session.Options
	gate pb.McbGateService
	data map[string]interface{}
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
func (s *srv) UID() string {
	return s.opts.Uid
}
func (s *srv) Bind(uid string) error {
	if uid == "" {
		return session.ErrIllegalUID
	}

	if s.UID() != "" {
		return session.ErrSessionAlreadyBound
	}
	s.opts.Uid = uid
	sessionData := &pb.Session{
		Id:  s.opts.Sid,
		Uid: s.opts.Uid,
	}
	rsp, err := s.gate.Bind(context.DefaultContext, sessionData, client.WithServerUid(s.opts.Fid), client.WithAuthToken())
	logger.Infof("bind: rsp:%v, err:%v", rsp, err)
	return err
}
func (s *srv) Kick() error {
	if s.UID() == "" {
		return session.ErrNoUIDBind
	}
	_, err := s.gate.Kick(context.DefaultContext, &pb.KickMsg{UserId: s.UID()}, client.WithServerUid(s.opts.Fid), client.WithAuthToken())
	return err
}
func (s *srv) PushSession() error {
	return nil
}
func (s *srv) Push(route string, v interface{}) error {
	b, err := proto.Marshal(v)
	if err != nil {
		return err
	}
	push := &pb.PushMsg{
		Route: route,
		Uid:   s.UID(),
		Data:  b,
	}
	rsp, err := s.gate.Push(context.DefaultContext, push, client.WithServerUid(s.opts.Fid), client.WithAuthToken())
	logger.Infof("bind: rsp:%v, err:%v", rsp, err)

	return err
}

func NewSession(opts ...session.Option) session.Session {
	s := &srv{
		data: make(map[string]interface{}),
		gate: pb.NewMcbGateService("gate", client.DefaultClient),
	}
	for _, o := range opts {
		o(&s.opts)
	}
	return s
}
