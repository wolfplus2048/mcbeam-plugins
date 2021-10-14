package session

type Options struct {
	SessionID string
	ServerID string
}
type Option func(o *Options)

func SessionID(id string) Option {
	return func(o *Options) {
		o.SessionID = id
	}
}
func ServerID(id string) Option {
	return func(o *Options) {
		o.ServerID = id
	}
}
