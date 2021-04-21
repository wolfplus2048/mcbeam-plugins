package cache

import (
	"context"
	"time"
)

type Z struct {
	Score  float64
	Member interface{}
}

type BitCount struct {
	Start, End int64
}

type ZStore struct {
	Keys    []string
	Weights []float64
	// Can be SUM, MIN or MAX.
	Aggregate string
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

type GeoLocation struct {
	Name                      string
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

type GeoPos struct {
	Longitude, Latitude float64
}

type GeoRadiusQuery struct {
	Radius float64
	// Can be m, km, ft, or mi. Default is km.
	Unit        string
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
	Count       int
	// Can be ASC or DESC. Default is no sort order.
	Sort      string
	Store     string
	StoreDist string
}

type Sort struct {
	By            string
	Offset, Count int64
	Get           []string
	Order         string
	Alpha         bool
}

var Default Cache

type Cache interface {
	Init(...Option) error
	Shutdown(ctx context.Context) error
	Del(ctx context.Context, keys ...string) (int64, error)
	Unlink(ctx context.Context, keys ...string) (int64, error)
	Dump(ctx context.Context, key string) (string, error)
	Exists(ctx context.Context, keys ...string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)

	ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error)
	Keys(ctx context.Context, pattern string) ([]string, error)
	Move(ctx context.Context, key string, db int) (bool, error)
	Persist(ctx context.Context, key string) (bool, error)
	PExpire(ctx context.Context, key string, expiration time.Duration) (bool, error)
	PExpireAt(ctx context.Context, key string, tm time.Time) (bool, error)
	PTTL(ctx context.Context, key string) (time.Duration, error)
	RandomKey(ctx context.Context, ) (string, error)
	Rename(ctx context.Context, key, newkey string) (string, error)
	RenameNX(ctx context.Context, key, newkey string) (bool, error)
	Sort(ctx context.Context, key string, sort *Sort) ([]string, error)
	SortStore(ctx context.Context, key, store string, sort *Sort) (int64, error)
	SortInterfaces(ctx context.Context, key string, sort *Sort) ([]interface{}, error)
	Touch(ctx context.Context, keys ...string) (int64, error)
	TTL(ctx context.Context, key string) (time.Duration, error)
	Type(ctx context.Context, key string) (string, error)
	Scan(ctx context.Context, cursor uint64, match string, count int64) (keys []string, cursor uint64, err error)
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, cursor uint64, err error)
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, cursor uint64, err error)
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, cursor uint64, err error)
	BitCount(ctx context.Context, key string, bitCount *BitCount) (int64, error)
	BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error)
	BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error)
	BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error)
	BitOpNot(ctx context.Context, destKey string, key string) (int64, error)
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) (int64, error)
	BitField(ctx context.Context, key string, args ...interface{}) ([]int64, error)
	Decr(ctx context.Context, key string) (int64, error)
	DecrBy(ctx context.Context, key string, decrement int64) (int64, error)
	Get(ctx context.Context, key string) (string, error)
	GetBit(ctx context.Context, key string, offset int64) (int64, error)
	GetRange(ctx context.Context, key string, start, end int64) (string, error)
	GetSet(ctx context.Context, key string, value interface{}) (string, error)
	Incr(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
	IncrByFloat(ctx context.Context, key string, value float64) (float64, error)

	MGet(ctx context.Context, keys ...string) ([]interface{}, error)
	MSet(ctx context.Context, values ...interface{}) (string, error)
	MSetNX(ctx context.Context, values ...interface{}) (bool, error)

	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error)
	SetBit(ctx context.Context, key string, offset int64, value int) (int64, error)
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	SetRange(ctx context.Context, key string, offset int64, value string) (int64, error)
	StrLen(ctx context.Context, key string) (int64, error)

	HDel(ctx context.Context, key string, fields ...string) (int64, error)
	HExists(ctx context.Context, key, field string) (bool, error)
	HGet(ctx context.Context, key, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error)
	HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error)
	HKeys(ctx context.Context, key string) ([]string, error)
	HLen(ctx context.Context, key string) (int64, error)
	HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	HSet(ctx context.Context, key string, values ...interface{}) (int64, error)
	HMSet(ctx context.Context, key string, values ...interface{}) (bool, error)
	HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error)
	HVals(ctx context.Context, key string) ([]string, error)

	BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (string, error)

	LIndex(ctx context.Context, key string, index int64) (string, error)
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) (int64, error)
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) (int64, error)
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) (int64, error)
	LLen(ctx context.Context, key string) (int64, error)
	LPop(ctx context.Context, key string) (string, error)
	LPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	LPushX(ctx context.Context, key string, values ...interface{}) (int64, error)
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error)
	LSet(ctx context.Context, key string, index int64, value interface{}) (string, error)
	LTrim(ctx context.Context, key string, start, stop int64) (string, error)
	RPop(ctx context.Context, key string) (string, error)
	RPopLPush(ctx context.Context, source, destination string) (string, error)
	RPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	RPushX(ctx context.Context, key string, values ...interface{}) (int64, error)

	SAdd(ctx context.Context, key string, members ...interface{}) (int64, error)
	SCard(ctx context.Context, key string) (int64, error)
	SDiff(ctx context.Context, keys ...string) ([]string, error)
	SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error)
	SInter(ctx context.Context, keys ...string) ([]string, error)
	SInterStore(ctx context.Context, destination string, keys ...string) (int64, error)
	SIsMember(ctx context.Context, key string, member interface{}) (bool, error)
	SMembers(ctx context.Context, key string) ([]string, error)
	SMembersMap(ctx context.Context, key string) (map[string]struct{}, error)
	SMove(ctx context.Context, source, destination string, member interface{}) (bool, error)
	SPop(ctx context.Context, key string) (string, error)
	SPopN(ctx context.Context, key string, count int64) ([]string, error)
	SRandMember(ctx context.Context, key string) (string, error)
	SRandMemberN(ctx context.Context, key string, count int64) ([]string, error)
	SRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	SUnion(ctx context.Context, keys ...string) ([]string, error)
	SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error)
	// TODO: X not support

	ZAdd(ctx context.Context, key string, members ...*Z) (int64, error)
	ZAddNX(ctx context.Context, key string, members ...*Z) (int64, error)
	ZAddXX(ctx context.Context, key string, members ...*Z) (int64, error)
	ZAddCh(ctx context.Context, key string, members ...*Z) (int64, error)
	ZAddNXCh(ctx context.Context, key string, members ...*Z) (int64, error)
	ZAddXXCh(ctx context.Context, key string, members ...*Z) (int64, error)
	ZIncr(ctx context.Context, key string, member *Z) (float64, error)
	ZIncrNX(ctx context.Context, key string, member *Z) (float64, error)
	ZIncrXX(ctx context.Context, key string, member *Z) (float64, error)
	ZCard(ctx context.Context, key string) (int64, error)
	ZCount(ctx context.Context, key, min, max string) (int64, error)
	ZLexCount(ctx context.Context, key, min, max string) (int64, error)
	ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error)
	ZInterStore(ctx context.Context, destination string, store *ZStore) (int64, error)
	ZPopMax(ctx context.Context, key string, count ...int64) ([]Z, error)
	ZPopMin(ctx context.Context, key string, count ...int64) ([]Z, error)
	ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error)
	ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)
	ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) ([]Z, error)
	ZRank(ctx context.Context, key, member string) (int64, error)
	ZRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error)
	ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error)
	ZRemRangeByLex(ctx context.Context, key, min, max string) (int64, error)
	ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error)
	ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)
	ZRevRangeByLex(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) ([]Z, error)
	ZRevRank(ctx context.Context, key, member string) (int64, error)
	ZScore(ctx context.Context, key, member string) (float64, error)
	ZUnionStore(ctx context.Context, dest string, store *ZStore) (int64, error)

	Publish(ctx context.Context, channel string, message interface{}) (int64, error)
	PubSubChannels(ctx context.Context, pattern string) ([]string, error)
	PubSubNumSub(ctx context.Context, channels ...string) (map[string]int64, error)
	PubSubNumPat(ctx context.Context, ) (int64, error)

	GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) (int64, error)
	GeoPos(ctx context.Context, key string, members ...string) ([]*GeoPos, error)
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) ([]GeoLocation, error)
	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) (int64, error)
	GeoRadiusByMember(ctx context.Context, key, member string, query *GeoRadiusQuery) ([]GeoLocation, error)
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *GeoRadiusQuery) (int64, error)
	GeoDist(ctx context.Context, key string, member1, member2, unit string) (float64, error)
	GeoHash(ctx context.Context, key string, members ...string) ([]string, error)
	ReadOnly(ctx context.Context) (string, error)
	ReadWrite(ctx context.Context) (string, error)
}
