package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/wolfplus2048/mcbeam-plugins/cache"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func init()  {
	cache.Default = New()
}
type srv struct {
	client  *redis.ClusterClient
	cmd     redis.Cmdable
	mutex   sync.Mutex
	options cache.Options
}

func New() cache.Cache {
	s := &srv{}
	if err := s.Init(); nil != err {
		log.Fatal(err)
		return nil
	}
	return s
}

func (p *srv) Init(opts ...cache.Option) error {
	for _, o := range opts {
		o(&p.options)
	}
	return p.configure()
}

func (p *srv) Shutdown(ctx context.Context) error {
	return p.close()
}

func newClusterOptions(rawurl string) (*redis.ClusterOptions, error) {
	addr := rawurl
	var hosts []string
	// 检查是否配置多个host
	if pos := strings.Index(rawurl, ","); pos > -1 {
		lpos := strings.Index(rawurl, "?")
		addr = rawurl[0:pos]
		if lpos != -1 {
			addr = addr + rawurl[lpos:]
		}
		hs := ""
		if pos := strings.Index(rawurl, "@"); pos > -1 {
			if lpos > -1 {
				hs = rawurl[pos+1 : lpos]
			} else {
				hs = rawurl[pos+1:]
			}
		}
		hosts = strings.Split(hs, ",")
	}
	// 解析配置参数
	option, err := redis.ParseURL(addr)
	if nil != err {
		log.Panicf("parse url, err:%s", err.Error())
		return nil, err
	}
	if len(hosts) == 0 {
		hosts = []string{option.Addr,}
	}
	log.Printf("redis cluster options, uri:%v addrs:%+v", addr, hosts)
	return &redis.ClusterOptions{
		Addrs:    hosts,
		Username: option.Username,
		Password: option.Password,

		DialTimeout:        option.DialTimeout,
		ReadTimeout:        option.ReadTimeout,
		WriteTimeout:       option.WriteTimeout,
		PoolSize:           option.PoolSize,
		PoolTimeout:        option.PoolTimeout,
		IdleTimeout:        option.IdleTimeout,
		IdleCheckFrequency: option.IdleCheckFrequency,
		MaxRetries:         option.MaxRetries,
	}, nil
}

func (p *srv) configure() error {

	addr := os.Getenv("MICRO_REDIS_ADDRESS")
	if addr == "" {
		addr = "redis://:root@redis-cluster-0:6379,redis-cluster-1:6379,redis-cluster-2:6379,redis-cluster-3:6379,redis-cluster-4:6379,redis-cluster-5:6379"
	}

	clusterOps, _ := newClusterOptions(addr)
	p.client = redis.NewClusterClient(clusterOps)
	if _, err := p.client.Ping(context.Background()).Result(); nil != err {
		log.Panicf("addr:%s err:%s", addr, err.Error())
		return err
	}
	p.cmd = p.client

	return nil
}

func (p *srv) close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if nil != p.client {
		defer func() { p.client = nil }()
		if err := p.client.Close(); nil != err {
			log.Panicf("err:%v\n", err)
			return err
		}
	}
	return nil
}

func (p *srv) Del(ctx context.Context, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Del(ctx, keys...).Result()
}
func (p *srv) Unlink(ctx context.Context, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Unlink(ctx, keys...).Result()
}
func (p *srv) Dump(ctx context.Context, key string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Dump(ctx, key).Result()
}
func (p *srv) Exists(ctx context.Context, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Exists(ctx, keys...).Result()
}
func (p *srv) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Expire(ctx, key, expiration).Result()
}
func (p *srv) ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ExpireAt(ctx, key, tm).Result()
}
func (p *srv) Keys(ctx context.Context, pattern string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Keys(ctx, pattern).Result()
}
func (p *srv) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Migrate(ctx, host, port, key, db, timeout).Result()
}
func (p *srv) Move(ctx context.Context, key string, db int) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Move(ctx, key, db).Result()
}
func (p *srv) ObjectRefCount(ctx context.Context, key string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ObjectRefCount(ctx, key).Result()
}
func (p *srv) ObjectEncoding(ctx context.Context, key string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ObjectEncoding(ctx, key).Result()
}
func (p *srv) ObjectIdleTime(ctx context.Context, key string) (time.Duration, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ObjectIdleTime(ctx, key).Result()
}
func (p *srv) Persist(ctx context.Context, key string) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Persist(ctx, key).Result()
}
func (p *srv) PExpire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.PExpire(ctx, key, expiration).Result()
}
func (p *srv) PExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.PExpireAt(ctx, key, tm).Result()
}
func (p *srv) PTTL(ctx context.Context, key string) (time.Duration, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.PTTL(ctx, key).Result()
}
func (p *srv) RandomKey(ctx context.Context) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.RandomKey(ctx).Result()
}
func (p *srv) Rename(ctx context.Context, key, newkey string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Rename(ctx, key, newkey).Result()
}
func (p *srv) RenameNX(ctx context.Context, key, newkey string) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.RenameNX(ctx, key, newkey).Result()
}
func (p *srv) Restore(ctx context.Context, key string, ttl time.Duration, value string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Restore(ctx, key, ttl, value).Result()
}
func (p *srv) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.RestoreReplace(ctx, key, ttl, value).Result()
}
func (p *srv) Sort(ctx context.Context, key string, sort *cache.Sort) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Sort(ctx, key, &redis.Sort{
		By:     sort.By,
		Offset: sort.Offset,
		Count:  sort.Count,
		Get:    sort.Get,
		Order:  sort.Order,
		Alpha:  sort.Alpha,
	}).Result()
}
func (p *srv) SortStore(ctx context.Context, key, store string, sort *cache.Sort) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SortStore(ctx, key, store, &redis.Sort{
		By:     sort.By,
		Offset: sort.Offset,
		Count:  sort.Count,
		Get:    sort.Get,
		Order:  sort.Order,
		Alpha:  sort.Alpha,
	}).Result()
}
func (p *srv) SortInterfaces(ctx context.Context, key string, sort *cache.Sort) ([]interface{}, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SortInterfaces(ctx, key, &redis.Sort{
		By:     sort.By,
		Offset: sort.Offset,
		Count:  sort.Count,
		Get:    sort.Get,
		Order:  sort.Order,
		Alpha:  sort.Alpha,
	}).Result()
}
func (p *srv) Touch(ctx context.Context, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Touch(ctx, keys...).Result()
}
func (p *srv) TTL(ctx context.Context, key string) (time.Duration, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.TTL(ctx, key).Result()
}
func (p *srv) Type(ctx context.Context, key string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Type(ctx, key).Result()
}
func (p *srv) Scan(ctx context.Context, cursor uint64, match string, count int64) (keys []string, cursor uint64, err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Scan(ctx, cursor, match, count).Result()
}
func (p *srv) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, cursor uint64, err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SScan(ctx, key, cursor, match, count).Result()
}
func (p *srv) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, cursor uint64, err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HScan(ctx, key, cursor, match, count).Result()
}
func (p *srv) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, cursor uint64, err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZScan(ctx, key, cursor, match, count).Result()
}
func (p *srv) Append(ctx context.Context, key, value string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Append(ctx, key, value).Result()
}
func (p *srv) BitCount(ctx context.Context, key string, bitCount *cache.BitCount) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BitCount(ctx, key, &redis.BitCount{
		Start: bitCount.Start,
		End:   bitCount.End,
	}).Result()
}
func (p *srv) BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BitOpAnd(ctx, destKey, keys...).Result()
}
func (p *srv) BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BitOpOr(ctx, destKey, keys...).Result()
}
func (p *srv) BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BitOpXor(ctx, destKey, keys...).Result()
}
func (p *srv) BitOpNot(ctx context.Context, destKey string, key string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BitOpNot(ctx, destKey, key).Result()
}
func (p *srv) BitPos(ctx context.Context, key string, bit int64, pos ...int64) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BitPos(ctx, key, bit, pos...).Result()
}
func (p *srv) BitField(ctx context.Context, key string, args ...interface{}) ([]int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BitField(ctx, key, args...).Result()
}
func (p *srv) Decr(ctx context.Context, key string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Decr(ctx, key).Result()
}
func (p *srv) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.DecrBy(ctx, key, decrement).Result()
}

func (p *srv) Get(ctx context.Context, key string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Get(ctx, key).Result()
}

func (p *srv) GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.GetBit(ctx, key, offset).Result()
}
func (p *srv) GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.GetRange(ctx, key, start, end).Result()
}
func (p *srv) GetSet(ctx context.Context, key string, value interface{}) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.GetSet(ctx, key, value).Result()
}
func (p *srv) Incr(ctx context.Context, key string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Incr(ctx, key).Result()
}
func (p *srv) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.IncrBy(ctx, key, value).Result()
}
func (p *srv) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.IncrByFloat(ctx, key, value).Result()
}
func (p *srv) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.MGet(ctx, keys...).Result()
}
func (p *srv) MSet(ctx context.Context, values ...interface{}) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.MSet(ctx, values...).Result()
}
func (p *srv) MSetNX(ctx context.Context, values ...interface{}) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.MSetNX(ctx, values...).Result()
}
func (p *srv) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Set(ctx, key, value, expiration).Result()
}
func (p *srv) SetBit(ctx context.Context, key string, offset int64, value int) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SetBit(ctx, key, offset, value).Result()
}
func (p *srv) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SetNX(ctx, key, value, expiration).Result()
}
func (p *srv) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SetXX(ctx, key, value, expiration).Result()
}
func (p *srv) SetRange(ctx context.Context, key string, offset int64, value string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SetRange(ctx, key, offset, value).Result()
}
func (p *srv) StrLen(ctx context.Context, key string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.StrLen(ctx, key).Result()
}
func (p *srv) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HDel(ctx, key, fields...).Result()
}
func (p *srv) HExists(ctx context.Context, key, field string) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HExists(ctx, key, field).Result()
}
func (p *srv) HGet(ctx context.Context, key, field string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HGet(ctx, key, field).Result()
}
func (p *srv) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HGetAll(ctx, key).Result()
}
func (p *srv) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HIncrBy(ctx, key, field, incr).Result()
}
func (p *srv) HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HIncrByFloat(ctx, key, field, incr).Result()
}
func (p *srv) HKeys(ctx context.Context, key string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HKeys(ctx, key).Result()
}
func (p *srv) HLen(ctx context.Context, key string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HLen(ctx, key).Result()
}
func (p *srv) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HMGet(ctx, key, fields...).Result()
}
func (p *srv) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HSet(ctx, key, values).Result()
}
func (p *srv) HMSet(ctx context.Context, key string, values ...interface{}) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HMSet(ctx, key, values...).Result()
}
func (p *srv) HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HSetNX(ctx, key, field, value).Result()
}
func (p *srv) HVals(ctx context.Context, key string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.HVals(ctx, key).Result()
}
func (p *srv) BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BLPop(ctx, timeout, keys...).Result()
}
func (p *srv) BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BRPop(ctx, timeout, keys...).Result()
}
func (p *srv) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BRPopLPush(ctx, source, destination, timeout).Result()
}
func (p *srv) LIndex(ctx context.Context, key string, index int64) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LIndex(ctx, key, index).Result()
}
func (p *srv) LInsert(ctx context.Context, key, op string, pivot, value interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LInsert(ctx, key, op, pivot, value).Result()
}
func (p *srv) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LInsertBefore(ctx, key, pivot, value).Result()
}
func (p *srv) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LInsertAfter(ctx, key, pivot, value).Result()
}
func (p *srv) LLen(ctx context.Context, key string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LLen(ctx, key).Result()
}
func (p *srv) LPop(ctx context.Context, key string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LPop(ctx, key).Result()
}
func (p *srv) LPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LPush(ctx, key, values...).Result()
}
func (p *srv) LPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.cmd.LPushX(ctx, key, values...).Result()
}
func (p *srv) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LRange(ctx, key, start, stop).Result()
}
func (p *srv) LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LRem(ctx, key, count, value).Result()
}
func (p *srv) LSet(ctx context.Context, key string, index int64, value interface{}) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LSet(ctx, key, index, value).Result()
}
func (p *srv) LTrim(ctx context.Context, key string, start, stop int64) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LTrim(ctx, key, start, stop).Result()
}
func (p *srv) RPop(ctx context.Context, key string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.RPop(ctx, key).Result()
}
func (p *srv) RPopLPush(ctx context.Context, source, destination string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.RPopLPush(ctx, source, destination).Result()
}
func (p *srv) RPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.RPush(ctx, key, values...).Result()
}
func (p *srv) RPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.RPushX(ctx, key, values...).Result()
}
func (p *srv) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SAdd(ctx, key, members...).Result()
}
func (p *srv) SCard(ctx context.Context, key string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SCard(ctx, key).Result()
}
func (p *srv) SDiff(ctx context.Context, keys ...string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SDiff(ctx, keys...).Result()
}
func (p *srv) SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SDiffStore(ctx, destination, keys...).Result()
}
func (p *srv) SInter(ctx context.Context, keys ...string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SInter(ctx, keys...).Result()
}
func (p *srv) SInterStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SInterStore(ctx, destination, keys...).Result()
}
func (p *srv) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SIsMember(ctx, key, member).Result()
}
func (p *srv) SMembers(ctx context.Context, key string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SMembers(ctx, key).Result()
}
func (p *srv) SMembersMap(ctx context.Context, key string) (map[string]struct{}, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SMembersMap(ctx, key).Result()
}
func (p *srv) SMove(ctx context.Context, source, destination string, member interface{}) (bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SMove(ctx, source, destination, member.Result()
}
func (p *srv) SPop(ctx context.Context, key string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SPop(ctx, key).Result()
}
func (p *srv) SPopN(ctx context.Context, key string, count int64) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SPopN(ctx, key, count).Result()
}
func (p *srv) SRandMember(ctx context.Context, key string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SRandMember(ctx, key).Result()
}
func (p *srv) SRandMemberN(ctx context.Context, key string, count int64) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SRandMemberN(ctx, key, count).Result()
}
func (p *srv) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SRem(ctx, key, members...).Result()
}
func (p *srv) SUnion(ctx context.Context, keys ...string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SUnion(ctx, keys...).Result()
}
func (p *srv) SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.SUnionStore(ctx, destination, keys...).Result()
}

func (p *srv) ZAdd(ctx context.Context, key string, members ...*cache.Z) (int64, error) {
	fields := make([]*cache.Z, 0, len(members))
	for _, it := range members {
		fields = append(fields, &redis.Z{Score: it.Score, Member: it.Member})
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZAdd(ctx, key, fields...).Result()
}
func (p *srv) ZAddNX(ctx context.Context, key string, members ...*cache.Z) (int64, error) {
	fields := make([]*cache.Z, 0, len(members))
	for _, it := range members {
		fields = append(fields, &redis.Z{Score: it.Score, Member: it.Member})
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZAddNX(ctx, key, fields...).Result()
}
func (p *srv) ZAddXX(ctx context.Context, key string, members ...*cache.Z) (int64, error) {
	fields := make([]*cache.Z, 0, len(members))
	for _, it := range members {
		fields = append(fields, &redis.Z{Score: it.Score, Member: it.Member})
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZAddXX(ctx, key, fields...).Result()
}
func (p *srv) ZAddCh(ctx context.Context, key string, members ...*cache.Z) (int64, error) {
	fields := make([]*cache.Z, 0, len(members))
	for _, it := range members {
		fields = append(fields, &redis.Z{Score: it.Score, Member: it.Member})
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZAddCh(ctx, key, fields...).Result()
}
func (p *srv) ZAddNXCh(ctx context.Context, key string, members ...*cache.Z) (int64, error) {
	fields := make([]*cache.Z, 0, len(members))
	for _, it := range members {
		fields = append(fields, &redis.Z{Score: it.Score, Member: it.Member})
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZAddNXCh(ctx, key, fields...).Result()
}
func (p *srv) ZAddXXCh(ctx context.Context, key string, members ...*cache.Z) (int64, error) {
	fields := make([]*cache.Z, 0, len(members))
	for _, it := range members {
		fields = append(fields, &redis.Z{Score: it.Score, Member: it.Member})
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZAddXXCh(ctx, key, fields...).Result()
}
func (p *srv) ZIncr(ctx context.Context, key string, member *cache.Z) (float64, error) {

	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZIncr(ctx, key, &redis.Z{
		Score:  member.Score,
		Member: member.Member,
	}).Result()
}
func (p *srv) ZIncrNX(ctx context.Context, key string, member *cache.Z) (float64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZIncrNX(ctx, key, &redis.Z{
		Score:  member.Score,
		Member: member.Member,
	}).Result()
}
func (p *srv) ZIncrXX(ctx context.Context, key string, member *cache.Z) (float64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZIncrXX(ctx, key, &redis.Z{
		Score:  member.Score,
		Member: member.Member,
	}).Result()
}
func (p *srv) ZCard(ctx context.Context, key string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZCard(ctx, key).Result()
}
func (p *srv) ZCount(ctx context.Context, key, min, max string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZCount(ctx, key, min, max).Result()
}
func (p *srv) ZLexCount(ctx context.Context, key, min, max string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZLexCount(ctx, key, min, max).Result()
}
func (p *srv) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZIncrBy(ctx, key, increment, member).Result()
}
func (p *srv) ZInterStore(ctx context.Context, destination string, store *cache.ZStore) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZInterStore(ctx, destination, &redis.ZStore{
		Keys:      store.Keys,
		Weights:   store.Weights,
		Aggregate: store.Aggregate,
	}).Result()
}
func (p *srv) ZPopMax(ctx context.Context, key string, count ...int64) ([]Z, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZPopMax(ctx, key, count...).Result()
}
func (p *srv) ZPopMin(ctx context.Context, key string, count ...int64) ([]Z, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZPopMin(ctx, key, count...).Result()
}
func (p *srv) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRange(ctx, key, start, stop).Result()
}
func (p *srv) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRangeWithScores(ctx, key, start, stop).Result()
}
func (p *srv) ZRangeByScore(ctx context.Context, key string, opt *cache.ZRangeBy) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Min:    opt.Min,
		Max:    opt.Max,
		Offset: opt.Offset,
		Count:  opt.Count,
	}).Result()
}
func (p *srv) ZRangeByLex(ctx context.Context, key string, opt *cache.ZRangeBy) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRangeByLex(ctx, key, &redis.ZRangeBy{
		Min:    opt.Min,
		Max:    opt.Max,
		Offset: opt.Offset,
		Count:  opt.Count,
	}).Result()
}
func (p *srv) ZRangeByScoreWithScores(ctx context.Context, key string, opt *cache.ZRangeBy) ([]Z, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min:    opt.Min,
		Max:    opt.Max,
		Offset: opt.Offset,
		Count:  opt.Count,
	}).Result()
}
func (p *srv) ZRank(ctx context.Context, key, member string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRank(ctx, key, member).Result()
}
func (p *srv) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRem(ctx, key, members...).Result()
}
func (p *srv) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRemRangeByRank(ctx, key, start, stop).Result()
}
func (p *srv) ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRemRangeByScore(ctx, key, min, max).Result()
}
func (p *srv) ZRemRangeByLex(ctx context.Context, key, min, max string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRemRangeByLex(ctx, key, min, max).Result()
}
func (p *srv) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRevRange(ctx, key, start, stop).Result()
}
func (p *srv) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]*cache.Z, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRevRangeWithScores(ctx, key, start, stop).Result()
}
func (p *srv) ZRevRangeByScore(ctx context.Context, key string, opt *cache.ZRangeBy) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRevRangeByScore(ctx, key, opt).Result()
}
func (p *srv) ZRevRangeByLex(ctx context.Context, key string, opt *cache.ZRangeBy) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRevRangeByLex(ctx, key, opt).Result()
}
func (p *srv) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *cache.ZRangeBy) ([]Z, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRevRangeByScoreWithScores(ctx, key, opt).Result()
}
func (p *srv) ZRevRank(ctx context.Context, key, member string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZRevRank(ctx, key, member).Result()
}
func (p *srv) ZScore(ctx context.Context, key, member string) (float64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZScore(ctx, key, member).Result()
}
func (p *srv) ZUnionStore(ctx context.Context, dest string, store *cache.ZStore) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ZUnionStore(ctx, dest, &redis.ZStore{
		Keys:      store.Keys,
		Weights:   store.Weights,
		Aggregate: store.Aggregate,
	}).Result()
}
func (p *srv) PFAdd(ctx context.Context, key string, els ...interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.PFAdd(ctx, key, els...).Result()
}
func (p *srv) PFCount(ctx context.Context, keys ...string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.PFCount(ctx, keys...).Result()
}
func (p *srv) PFMerge(ctx context.Context, dest string, keys ...string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.PFMerge(ctx, dest, keys...).Result()
}
func (p *srv) BgRewriteAOF(ctx context.Context) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BgRewriteAOF(ctx).Result()
}
func (p *srv) BgSave(ctx context.Context) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.BgSave(ctx).Result()
}
func (p *srv) ClientID(ctx context.Context) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ClientID(ctx).Result()
}
func (p *srv) ConfigGet(ctx context.Context, parameter string) ([]interface{}, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ConfigGet(ctx, parameter).Result()
}
func (p *srv) ConfigSet(ctx context.Context, parameter, value string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ConfigSet(ctx, parameter, value).Result()
}
func (p *srv) DBSize(ctx context.Context) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.DBSize(ctx).Result()
}
func (p *srv) Info(ctx context.Context, section ...string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Info(ctx, section...).Result()
}
func (p *srv) LastSave(ctx context.Context) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.LastSave(ctx).Result()
}
func (p *srv) Save(ctx context.Context) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Save(ctx).Result()
}

func (p *srv) Time(ctx context.Context) (time.Time, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Time(ctx).Result()
}

func (p *srv) Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Eval(ctx, script, keys, args...).Result()
}
func (p *srv) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.EvalSha(ctx, sha1, keys, args...).Result()
}
func (p *srv) ScriptExists(ctx context.Context, hashes ...string) ([]bool, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ScriptExists(ctx, hashes...).Result()
}
func (p *srv) ScriptFlush(ctx context.Context, ) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ScriptFlush(ctx).Result()
}
func (p *srv) ScriptKill(ctx context.Context, ) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ScriptKill(ctx).Result()
}
func (p *srv) ScriptLoad(ctx context.Context, script string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ScriptLoad(ctx, script).Result()
}
func (p *srv) DebugObject(ctx context.Context, key string) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.DebugObject(ctx, key).Result()
}
func (p *srv) Publish(ctx context.Context, channel string, message interface{}) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.Publish(ctx, channel, message).Result()
}
func (p *srv) PubSubChannels(ctx context.Context, pattern string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.PubSubChannels(ctx, pattern).Result()
}
func (p *srv) PubSubNumSub(ctx context.Context, channels ...string) (map[string]int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.PubSubNumSub(ctx, channels...).Result()
}
func (p *srv) PubSubNumPat(ctx context.Context, ) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.PubSubNumPat(ctx).Result()
}
func (p *srv) GeoAdd(ctx context.Context, key string, geoLocation ...*cache.GeoLocation) (int64, error) {
	locations := make([]*redis.GeoLocation, 0, len(geoLocation))
	for _, it := range geoLocation {
		loctions = append(loctions, &redis.GeoLocation{
			Name:      it.Name,
			Longitude: it.Longitude,
			Latitude:  it.Latitude,
			Dist:      it.Dist,
			GeoHash:   it.GeoHash,
		})
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.GeoAdd(ctx, key, locations...).Result()
}
func (p *srv) GeoPos(ctx context.Context, key string, members ...string) ([]*cache.GeoPos, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	res, err := p.cmd.GeoPos(ctx, key, members...).Result()
	if nil != err {
		return nil, err
	}
	positions := make([]*cache.GeoPos, 0, len(res))
	for _, it := range res {
		positions = append(positions, &cache.GeoPos{Longitude: it.Longitude, Latitude: it.Latitude})
	}
	return positions, nil
}
func (p *srv) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *cache.GeoRadiusQuery) ([]*cache.GeoLocation, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	res, err := p.cmd.GeoRadius(ctx, key, longitude, latitude, &redis.GeoRadiusQuery{
		Radius:      query.Radius,
		Unit:        query.Unit,
		WithCoord:   query.WithCoord,
		WithDist:    query.WithDist,
		WithGeoHash: query.WithGeoHash,
		Count:       query.Count,
		Sort:        query.Sort,
		Store:       query.Store,
		StoreDist:   query.StoreDist,
	}).Result()
	if nil != err {
		return nil, err
	}
	items := make([]*cache.GeoLocation, 0, len(res))
	for _, it := range res {
		items = append(items, &cache.GeoLocation{Name: it.Name, Longitude: it.Longitude, Latitude: it.Latitude, Dist: it.Dist, GeoHash: it.GeoHash})
	}
	return items, nil
}
func (p *srv) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *cache.GeoRadiusQuery) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.GeoRadiusStore(ctx, key, longitude, latitude, &redis.GeoRadiusQuery{
		Radius:      query.Radius,
		Unit:        query.Unit,
		WithCoord:   query.WithCoord,
		WithDist:    query.WithDist,
		WithGeoHash: query.WithGeoHash,
		Count:       query.Count,
		Sort:        query.Sort,
		Store:       query.Store,
		StoreDist:   query.StoreDist,
	}).Result()
}
func (p *srv) GeoRadiusByMember(ctx context.Context, key, member string, query *cache.GeoRadiusQuery) ([]*cache.GeoLocation, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	res, err := p.cmd.GeoRadiusByMember(ctx, key, member, &redis.GeoRadiusQuery{
		Radius:      query.Radius,
		Unit:        query.Unit,
		WithCoord:   query.WithCoord,
		WithDist:    query.WithDist,
		WithGeoHash: query.WithGeoHash,
		Count:       query.Count,
		Sort:        query.Sort,
		Store:       query.Store,
		StoreDist:   query.StoreDist,
	}).Result()
	if nil != err {
		return nil, err
	}
	items := make([]*cache.GeoLocation, 0, len(res))
	for _, it := range res {
		items = append(items, &cache.GeoLocation{Name: it.Name, Longitude: it.Longitude, Latitude: it.Latitude, Dist: it.Dist, GeoHash: it.GeoHash})
	}
	return items, nil
}
func (p *srv) GeoRadiusByMemberStore(ctx context.Context, key, member string, query *cache.GeoRadiusQuery) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.GeoRadiusByMemberStore(ctx, key, member, &redis.GeoRadiusQuery{
		Radius:      query.Radius,
		Unit:        query.Unit,
		WithCoord:   query.WithCoord,
		WithDist:    query.WithDist,
		WithGeoHash: query.WithGeoHash,
		Count:       query.Count,
		Sort:        query.Sort,
		Store:       query.Store,
		StoreDist:   query.StoreDist,
	}).Result()
}
func (p *srv) GeoDist(ctx context.Context, key string, member1, member2, unit string) (float64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.GeoDist(ctx, key, member1, member2, unit).Result()
}
func (p *srv) GeoHash(ctx context.Context, key string, members ...string) ([]string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.GeoHash(ctx, key, members...).Result()
}
func (p *srv) ReadOnly(ctx context.Context, ) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ReadOnly(ctx).Result()
}
func (p *srv) ReadWrite(ctx context.Context, ) (string, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.cmd.ReadWrite(ctx).Result()
}
