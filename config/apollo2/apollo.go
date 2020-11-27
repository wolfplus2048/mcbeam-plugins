package apollo2

import (
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
	agollo "github.com/philchia/agollo/v4"
)


type apollo struct {
	opts config.Options
	client agollo.Client
	namespace string
}

func NewConfig(opts ...config.Option) config.Config {
	a := &apollo{}
	for _, o := range opts {
		o(&a.opts)
	}
	a.configure()
	return a
}
func (a *apollo) configure() {
	config, ok := a.opts.Context.Value(appConfigKey{}).(*agollo.Conf)
	if !ok {
		logger.Fatal("load apollo2 config failed")
	}
	a.client = agollo.NewClient(config)
	err := a.client.Start()
	if err != nil {
		logger.Fatal(err)
	}
	if len(config.NameSpaceNames) > 0 {
		a.namespace = config.NameSpaceNames[0]
	} else {
		a.namespace = "application"
	}
}
func (a *apollo) Get(path string, options ...config.Option) (config.Value, error) {
	nullValue := config.NewJSONValue([]byte("null"))
	value := a.client.GetString(path, agollo.WithNamespace(a.namespace))
	if len(value) > 0 {
		return config.NewJSONValue([]byte(value)), nil
	}
	return nullValue, nil
}

func (a *apollo) Set(path string, val interface{}, options ...config.Option) error {
	panic("implement me")
}

func (a *apollo) Delete(path string, options ...config.Option) error {
	panic("implement me")
}
