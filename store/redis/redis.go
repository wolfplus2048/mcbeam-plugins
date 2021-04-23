package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
)

type rkv struct {
	options store.Options
	Client  *redis.Client
}

func (r *rkv) Init(opts ...store.Option) error {
	for _, o := range opts {
		o(&r.options)
	}

	return r.configure()
}

func (r *rkv) Close() error {
	return r.Client.Close()
}

func (r *rkv) Read(key string, opts ...store.ReadOption) ([]*store.Record, error) {
	options := store.ReadOptions{Context: context.Background()}
	options.Table = r.options.Table

	for _, o := range opts {
		o(&options)
	}

	var keys []string

	rkey := fmt.Sprintf("%s%s", options.Table, key)
	// Handle Prefix
	// TODO suffix
	if options.Prefix {
		prefixKey := fmt.Sprintf("%s*", rkey)
		fkeys, err := r.Client.Keys(options.Context, prefixKey).Result()
		if err != nil {
			return nil, err
		}
		// TODO Limit Offset

		keys = append(keys, fkeys...)

	} else {
		keys = []string{rkey}
	}

	if v := options.Context.Value(readZMemberKey{}); nil != v {
		field, ok := v.(*readZMember)
		if !ok {
			return nil, store.ErrNotFound
		}
		return r.readSortedSet(rkey, field, &options)
	} else if v := options.Context.Value(readZRangeByIndexKey{}); nil != v {
		return r.rangeByIndex(rkey, v.(*readZRangeByIndex), &options)
	} else if v := options.Context.Value(readZRangeWithScoreKey{}); nil != v {
		return r.rangeByScore(rkey, v.(*redis.ZRangeBy), &options)
	}

	records := make([]*store.Record, 0, len(keys))
	for _, rkey = range keys {
		val, err := r.Client.Get(options.Context, rkey).Bytes()

		if err != nil && err == redis.Nil {
			return nil, store.ErrNotFound
		} else if err != nil {
			return nil, err
		}

		if val == nil {
			return nil, store.ErrNotFound
		}

		d, err := r.Client.TTL(options.Context, rkey).Result()
		if err != nil {
			return nil, err
		}

		records = append(records, &store.Record{
			Key:    key,
			Value:  val,
			Expiry: d,
		})
	}

	return records, nil
}

func (r *rkv) rangeByIndex(key string, rangeBy *readZRangeByIndex, options *store.ReadOptions) ([]*store.Record, error) {
	records := make([]*store.Record, 0, 1)
	var err error
	var members []redis.Z
	if rangeBy.Asc {
		members, err = r.Client.ZRangeWithScores(options.Context, key, rangeBy.Start, rangeBy.End).Result()
	} else {
		members, err = r.Client.ZRevRangeWithScores(options.Context, key, rangeBy.Start, rangeBy.End).Result()
	}
	if nil != err {
		return nil, err
	}
	for idx, it := range members {
		records = append(records, &store.Record{
			Key:      key,
			Value:    []byte(it.Member.(string)),
			Metadata: map[string]interface{}{"rank": int64(idx) + rangeBy.Start, "score": it.Score},
			Expiry:   0,
		})
	}

	return records, nil
}

func (r *rkv) rangeByScore(key string, rangeBy *redis.ZRangeBy, options *store.ReadOptions) ([]*store.Record, error) {
	records := make([]*store.Record, 0, 1)
	var err error
	var members []redis.Z
	asc := true
	if v := options.Context.Value(readZRangeSortKey{}); nil != v {
		asc = v.(bool)
	}
	if asc {
		members, err = r.Client.ZRangeByScoreWithScores(options.Context, key, rangeBy).Result()
	} else {
		members, err = r.Client.ZRevRangeByScoreWithScores(options.Context, key, rangeBy).Result()
	}
	if nil != err {
		return nil, err
	}
	min := int64(0)
	if "-inf" != rangeBy.Min {
		min, err = strconv.ParseInt(rangeBy.Min, 10, 64)
		if nil != err {
			return nil, err
		}
	}

	for idx, it := range members {
		records = append(records, &store.Record{
			Key:      key,
			Value:    []byte(it.Member.(string)),
			Metadata: map[string]interface{}{"rank": int64(idx) + min, "score": it.Score},
			Expiry:   0,
		})
	}

	return records, nil
}

// sorted set 可以读取到的数据: 1.排名（小->大，大->小） 2.分数
func (r *rkv) readSortedSet(key string, field *readZMember, options *store.ReadOptions) ([]*store.Record, error) {
	records := make([]*store.Record, 0, 1)
	score, err := r.Client.ZScore(options.Context, key, field.Member).Result()
	if nil != err {
		return nil, err
	}
	rank := int64(0)
	if field.Asc {
		// 升序
		idx, err := r.Client.ZRank(options.Context, key, field.Member).Result()
		if nil != err {
			return nil, err
		}
		rank = idx
	} else {
		// 倒序
		idx, err := r.Client.ZRevRank(options.Context, key, field.Member).Result()
		if nil != err {
			return nil, err
		}
		rank = idx
	}
	log.Debugf("%+v", field)
	record := &store.Record{
		Key:      key,
		Value:    []byte(field.Member),
		Metadata: map[string]interface{}{"rank": rank, "score": score},
		Expiry:   0,
	}
	records = append(records, record)
	return records, nil
}

func (r *rkv) Delete(key string, opts ...store.DeleteOption) error {
	options := store.DeleteOptions{Context: context.Background()}
	options.Table = r.options.Table

	for _, o := range opts {
		o(&options)
	}
	rkey := fmt.Sprintf("%s%s", options.Table, key)

	if v := options.Context.Value(deleteZMemberKey{}); nil != v {
		return r.deleteSortedSetMember(rkey, v.(string), options)
	}

	return r.Client.Del(options.Context, rkey).Err()
}

func (r *rkv) deleteSortedSetMember(key string, member string, opts store.DeleteOptions) error {
	return r.Client.ZRem(opts.Context, key, member).Err()
}

func (r *rkv) writeSortedSet(key string, record *store.Record, score float64, options store.WriteOptions) error {
	return r.Client.ZAdd(options.Context, key, &redis.Z{
		Score:  score,
		Member: record.Value,
	}).Err()
}

func (r *rkv) Write(record *store.Record, opts ...store.WriteOption) error {
	options := store.WriteOptions{Context: context.Background()}
	options.Table = r.options.Table

	for _, o := range opts {
		o(&options)
	}
	rkey := fmt.Sprintf("%s%s", options.Table, record.Key)
	if v := options.Context.Value(writeZScoreKey{}); v != nil {
		return r.writeSortedSet(rkey, record, v.(float64), options)
	}
	return r.Client.Set(options.Context, rkey, record.Value, record.Expiry).Err()
}

func (r *rkv) List(opts ...store.ListOption) ([]string, error) {
	options := store.ListOptions{Context: context.Background()}
	options.Table = r.options.Table

	for _, o := range opts {
		o(&options)
	}

	keys, err := r.Client.Keys(options.Context, "*").Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (r *rkv) Options() store.Options {
	return r.options
}

func (r *rkv) String() string {
	return "redis"
}

func NewStore(opts ...store.Option) store.Store {
	var options store.Options
	for _, o := range opts {
		o(&options)
	}

	s := new(rkv)
	s.options = options

	if err := s.configure(); err != nil {
		log.Fatal(err)
	}

	return s
}

func (r *rkv) configure() error {
	var redisOptions *redis.Options
	nodes := r.options.Nodes

	if len(nodes) == 0 {
		nodes = []string{"redis://127.0.0.1:6379"}
	}

	redisOptions, err := redis.ParseURL(nodes[0])
	if err != nil {
		//Backwards compatibility
		redisOptions = &redis.Options{
			Addr:     nodes[0],
			Password: "", // no password set
			DB:       0,  // use default DB
		}
	}

	r.Client = redis.NewClient(redisOptions)
	return nil
}
