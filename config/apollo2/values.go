package apollo2

import (
	"encoding/json"
	"github.com/micro/micro/v3/service/config"
	"github.com/zouyx/agollo/v4/storage"
	"time"
)

// value represents a value of any type
type value struct {
	cache *storage.Config
	path string
}

func newValue(c *storage.Config, path string) config.Value {
	return &value{cache: c, path: path}
}
func (v value) Exists() bool {
	panic("implement me")
}

func (v value) Bool(def bool) bool {
	return v.cache.GetBoolValue(v.path, def)
}

func (v value) Int(def int) int {
	return v.cache.GetIntValue(v.path, def)
}

func (v value) String(def string) string {
	return v.cache.GetStringValue(v.path, def)
}

func (v value) Float64(def float64) float64 {
	return v.cache.GetFloatValue(v.path, def)
}

func (v value) Duration(def time.Duration) time.Duration {

	t := v.cache.GetStringValue(v.path, "")

	value, err := time.ParseDuration(t)
	if err != nil {
		return def
	}
	return value
}

func (v value) StringSlice(def []string) []string {
	return v.cache.GetStringSliceValue(v.path)
}

func (v value) StringMap(def map[string]string) map[string]string {
	panic("implement me")
}

func (v value) Scan(val interface{}) error {
	data, err := v.cache.GetCache().Get(v.path)
	if err != nil {
		return err
	}
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func (v value) Bytes() []byte {
	data, err := v.cache.GetCache().Get(v.path)
	if err != nil {
		return nil
	}
	ret, _ := data.([]byte)
	return ret
}

