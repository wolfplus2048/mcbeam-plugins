package apollo

import (
	"github.com/micro/micro/v3/service/config"
	agollo "github.com/philchia/agollo/v4"
)


type apollo struct {
	opts config.Options
	client agollo.Client
}

func New(opts ...config.Option) config.Config {
	a := &apollo{}
	for _, o := range opts {
		o(&a.opts)
	}
	a.configure()
	return a
}
func (a apollo) configure() {
	var config *agollo.Conf
	if config, ok := a.opts.Context.Value(appConfigKey{}).(*agollo.Conf); ok {
		a.client = agollo.NewClient(config)

	}
}
func (a apollo) Get(path string, options ...config.Option) (config.Value, error) {
	panic("implement me")
}

func (a apollo) Set(path string, val interface{}, options ...config.Option) error {
	panic("implement me")
}

func (a apollo) Delete(path string, options ...config.Option) error {
	panic("implement me")
}
