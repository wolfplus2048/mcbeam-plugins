package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/micro/micro/v3/service/store"
)

type writeZScoreKey struct{}

func WriteZScore(score float64) store.WriteOption {
	return func(w *store.WriteOptions) {
		if nil == w.Context {
			w.Context = context.Background()
		}
		w.Context = context.WithValue(w.Context, writeZScoreKey{}, score)
	}
}

type readZMemberKey struct{}

type readZMember struct {
	Member string
	Asc    bool
}

// 读取成员对应的分数和值
func ReadZMember(member string, sortAsc bool) store.ReadOption {
	return func(r *store.ReadOptions) {
		if nil == r.Context {
			r.Context = context.Background()
		}
		r.Context = context.WithValue(r.Context, readZMemberKey{}, &readZMember{Member: member, Asc: sortAsc})
	}
}

type readZRangeWithScoreKey struct{}
type readZRangeSortKey struct{}

// 读取指定分数区间的成员
func ReadZRangeByScore(min, max string, offset, count int64, asc bool) store.ReadOption {
	return func(r *store.ReadOptions) {
		if nil == r.Context {
			r.Context = context.Background()
		}
		r.Context = context.WithValue(r.Context, readZRangeWithScoreKey{}, &redis.ZRangeBy{Min: min, Max: max, Offset: offset, Count: count})
		r.Context = context.WithValue(r.Context, readZRangeSortKey{}, asc)
	}
}

type readZRangeByIndexKey struct{}
type readZRangeByIndex struct {
	Start, End int64
	Asc        bool
}

// 按顺序读取分数排名区间的成员列表
func ReadZRange(start, end int64, asc bool) store.ReadOption {
	return func(r *store.ReadOptions) {
		if nil == r.Context {
			r.Context = context.Background()
		}
		r.Context = context.WithValue(r.Context, readZRangeByIndexKey{}, &readZRangeByIndex{Start: start, End: end, Asc: asc})
	}
}

type deleteZMemberKey struct{}

func DeleteZMember(member string) store.DeleteOption {
	return func(r *store.DeleteOptions) {
		if nil == r.Context {
			r.Context = context.Background()
		}
		r.Context = context.WithValue(r.Context, deleteZMemberKey{}, member)
	}
}
