package apollo

import (
	"context"
	"github.com/micro/micro/v3/service/config"
	c "github.com/zouyx/agollo/v4/env/config"
)

type appConfigKey struct{}

func WithConfig(apollo *c.AppConfig) config.Option {
	return func(o *config.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, appConfigKey{}, apollo)
	}
}
