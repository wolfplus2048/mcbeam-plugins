package apollo

import (
	"fmt"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
	"github.com/zouyx/agollo/v4"
	agoconfig "github.com/zouyx/agollo/v4/env/config"
)

type conf struct {
	opts config.Options
	namespace string
	client *agollo.Client
}

func (c *conf) configure()  {
	agollo.SetLogger(&DefaultLogger{})

	client, err := agollo.StartWithConfig(func() (*agoconfig.AppConfig, error) {
		var apollo, ok = c.opts.Context.Value(appConfigKey{}).(*agoconfig.AppConfig)
		if ok {
			c.namespace = apollo.NamespaceName
			return apollo, nil
		}
		return nil, fmt.Errorf("no apollo config info")
	})

	if err != nil{
		logger.Fatal(err)
	}
	c.client = client
}
func (c *conf) Get(path string, options ...config.Option) (config.Value, error) {

	config := c.client.GetConfig(c.namespace)
	if config == nil {
		return nil, fmt.Errorf("namespace not exists")
	}

	return newValue(config, path), nil

}

func (c *conf) Set(path string, val interface{}, options ...config.Option) error {
	panic("implement me")
}

func (c *conf) Delete(path string, options ...config.Option) error {
	panic("implement me")
}
func NewConfig(opts ...config.Option) config.Config {
	conf := &conf{}
	for _, o := range opts {
		o(&conf.opts)
	}
	conf.configure()
	return conf
}
