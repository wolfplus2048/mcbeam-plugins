package session

type Options struct {
	Sid int64
	Fid string
	Uid string
}
type Option func(o *Options)

func Sid(sid int64) Option {
	return func(o *Options) {
		o.Sid = sid
	}
}
func Fid(fid string) Option {
	return func(o *Options) {
		o.Fid = fid
	}
}
func Uid(uid string) Option {
	return func(o *Options) {
		o.Uid = uid
	}
}