package wrapper

import (
	"context"
	"github.com/micro/micro/v3/service/context/metadata"
	"github.com/micro/micro/v3/service/server"
	"github.com/wolfplus2048/mcbeam-plugins/ws_session/v3"
	cli "github.com/wolfplus2048/mcbeam-plugins/ws_session/v3/client"
)

func SessionHandler() server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			sessionID, sok := metadata.Get(ctx, "Micro-SessionID")
			serverID, fok := metadata.Get(ctx, "Micro-ServerID")
			if !sok || !fok{
				return h(ctx, req, rsp)
			}
			s := cli.NewSession(session.ServerID(serverID), session.SessionID(sessionID))
			ctx = context.WithValue(ctx, session.SessionCtxKey{}, s)
			return h(ctx, req, rsp)
		}
	}
}
