package redis

import (
	"context"
	"github.com/micro/micro/v3/service/store"
)

type clusterKey struct{}

func WithCluster(addrs string) store.Option {
	return func(o *store.Options) {
		if nil == o.Context {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, clusterKey{}, addrs)
	}
}

type memberKey struct{}

func WriteScore(filed string, score float64) store.WriteOption {
	return func(w *store.WriteOptions) {
		if nil == w.Context {
			w.Context = context.Background()
		}
		w.Context = context.WithValue(w.Context, memberKey{}, score)
	}
}

type memberNameKey struct{}

func ReadSortedSet(field string) store.ReadOption {
	return func(r *store.ReadOptions) {
		if nil == r.Context {
			r.Context = context.Background()
		}
		r.Context = context.WithValue(r.Context, memberNameKey{}, field)
	}
}

type sortTypeKey struct{}

func ReadSortedSetBySort(asc bool) store.ReadOption {
	return func(r *store.ReadOptions) {
		if nil == r.Context {
			r.Context = context.Background()
		}
		r.Context = context.WithValue(r.Context, sortTypeKey{}, asc)
	}
}

func DeleteSortedSetMember(field string) store.DeleteOption {
	return func(r *store.DeleteOptions) {
		if nil == r.Context {
			r.Context = context.Background()
		}
		r.Context = context.WithValue(r.Context, memberNameKey{}, field)
	}
}

type sortBy struct {
	Key        string
	Start, End int64
	Asc        bool
}
type sortByKey struct{}

func ListSortedSetBySort(key string, start, end int64, asc bool) store.ListOption {
	return func(r *store.ListOptions) {
		if nil == r.Context {
			r.Context = context.Background()
		}
		r.Context = context.WithValue(r.Context, sortByKey{}, &sortBy{
			Key:   key,
			Start: start,
			End:   end,
			Asc:   asc,
		})
	}
}
