// Package etcd is an etcd implementation of lock
package etcd

import (
	"context"
	"crypto/tls"
	"errors"
	"go.uber.org/zap"
	"net"
	"path"
	"strings"
	gosync "sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	cc "github.com/coreos/etcd/clientv3/concurrency"
	"github.com/micro/micro/v3/service/sync"
)

type etcdSync struct {
	options sync.Options
	path    string
	client  *clientv3.Client

	mtx   gosync.Mutex
	locks map[string]*etcdLock
}

type etcdLock struct {
	s *cc.Session
	m *cc.Mutex
}

type etcdLeader struct {
	opts sync.LeaderOptions
	s    *cc.Session
	e    *cc.Election
	id   string
}

func (e *etcdSync) Leader(id string, opts ...sync.LeaderOption) (sync.Leader, error) {
	var options sync.LeaderOptions
	for _, o := range opts {
		o(&options)
	}

	// make path
	path := path.Join(e.path, strings.Replace(e.options.Prefix+id, "/", "-", -1))

	s, err := cc.NewSession(e.client)
	if err != nil {
		return nil, err
	}

	l := cc.NewElection(s, path)

	if err := l.Campaign(context.TODO(), id); err != nil {
		return nil, err
	}

	return &etcdLeader{
		opts: options,
		e:    l,
		id:   id,
	}, nil
}

func (e *etcdLeader) Status() chan bool {
	ch := make(chan bool, 1)
	ech := e.e.Observe(context.Background())

	go func() {
		for r := range ech {
			if string(r.Kvs[0].Value) != e.id {
				ch <- true
				close(ch)
				return
			}
		}
	}()

	return ch
}

func (e *etcdLeader) Resign() error {
	return e.e.Resign(context.Background())
}

func (e *etcdSync) Init(opts ...sync.Option) error {
	for _, o := range opts {
		o(&e.options)
	}
	return nil
}

func (e *etcdSync) Options() sync.Options {
	return e.options
}

func (e *etcdSync) Lock(id string, opts ...sync.LockOption) error {
	var options sync.LockOptions
	for _, o := range opts {
		o(&options)
	}

	// make path
	path := path.Join(e.path, strings.Replace(e.options.Prefix+id, "/", "-", -1))

	var sopts []cc.SessionOption
	if options.TTL > 0 {
		sopts = append(sopts, cc.WithTTL(int(options.TTL.Seconds())))
	}

	s, err := cc.NewSession(e.client, sopts...)
	if err != nil {
		return err
	}

	m := cc.NewMutex(s, path)

	if err := m.Lock(context.TODO()); err != nil {
		return err
	}

	e.mtx.Lock()
	e.locks[id] = &etcdLock{
		s: s,
		m: m,
	}
	e.mtx.Unlock()
	return nil
}

func (e *etcdSync) Unlock(id string) error {
	e.mtx.Lock()
	defer e.mtx.Unlock()
	v, ok := e.locks[id]
	if !ok {
		return errors.New("lock not found")
	}
	err := v.m.Unlock(context.Background())
	delete(e.locks, id)
	return err
}

func (e *etcdSync) String() string {
	return "etcd"
}

func newClient(e *etcdSync) (*clientv3.Client, error) {
	config := clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	}

	if e.options.Timeout == 0 {
		e.options.Timeout = 5 * time.Second
	}

	if e.options.Secure || e.options.TLSConfig != nil {
		tlsConfig := e.options.TLSConfig
		if tlsConfig == nil {
			tlsConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}

		config.TLS = tlsConfig
	}

	if e.options.Context != nil {
		u, ok := e.options.Context.Value(authKey{}).(*authCreds)
		if ok {
			config.Username = u.Username
			config.Password = u.Password
		}
		cfg, ok := e.options.Context.Value(logConfigKey{}).(*zap.Config)
		if ok && cfg != nil {
			config.LogConfig = cfg
		}
	}

	var cAddrs []string

	for _, address := range e.options.Nodes {
		if len(address) == 0 {
			continue
		}
		addr, port, err := net.SplitHostPort(address)
		if ae, ok := err.(*net.AddrError); ok && ae.Err == "missing port in address" {
			port = "2379"
			addr = address
			cAddrs = append(cAddrs, net.JoinHostPort(addr, port))
		} else if err == nil {
			cAddrs = append(cAddrs, net.JoinHostPort(addr, port))
		}
	}

	// if we got addrs then we'll update
	if len(cAddrs) > 0 {
		config.Endpoints = cAddrs
	}

	// check if the endpoints have https://
	if config.TLS != nil {
		for i, ep := range config.Endpoints {
			if !strings.HasPrefix(ep, "https://") {
				config.Endpoints[i] = "https://" + ep
			}
		}
	}

	cli, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}

	return cli, nil
}

// configure will setup the registry with new options
func configure(e *etcdSync, opts ...sync.Option) error {
	for _, o := range opts {
		o(&e.options)
	}

	// setup the client
	cli, err := newClient(e)
	if err != nil {
		return err
	}

	if e.client != nil {
		e.client.Close()
	}

	// setup new client
	e.client = cli

	return nil
}
func NewSync(opts ...sync.Option) sync.Sync {
	e := &etcdSync{
		path:    "/micro/sync",
		options: sync.Options{},
		locks:   make(map[string]*etcdLock),
	}
	configure(e, opts...)
	return e
}
