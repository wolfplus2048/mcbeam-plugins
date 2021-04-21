package cache

type Options struct {
	Addrs string
}

type Option func(options *Options)

func WithAddress(addrs string) Option {
	return func(options *Options) {
		options.Addrs = addrs
	}
}