package apollo

import (
	"github.com/micro/micro/v3/service/config"
	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/env/config"
)

type conf struct {
	client *agollo.Client
}

func (c conf) Get(path string, options ...config.Option) (config.Value, error) {
	panic("implement me")
}

func (c conf) Set(path string, val interface{}, options ...config.Option) error {
	panic("implement me")
}

func (c conf) Delete(path string, options ...config.Option) error {
	panic("implement me")
}
func newConfig(ns string) config.Config {
	return &conf{}
}
