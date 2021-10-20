package session

import (
	"context"
	"errors"
)

var (
	ErrIllegalUID          = errors.New("illegal uid")
	ErrSessionAlreadyBound = errors.New("session is already bound to an uid")
	ErrNoUIDBind           = errors.New("you have to bind an UID to the session to do that")
)

type Session interface {
	Init(opts ...Option)
	Options() Options
	Send(route string, v interface{}) error
	Bind(status map[string]string) error
	Kick() error
	String() string
}

type SessionCtxKey struct{}

func GetSessionFromCtx(ctx context.Context) Session {
	s, ok := ctx.Value(SessionCtxKey{}).(Session)
	if ok {
		return s
	}
	return nil
}
