package redis

import (
	"context"
	"errors"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"time"
)

// New Redis redis模式
func New(opts ...Option) (*Redis, error) {

	op := setDefault()
	for _, o := range opts {
		if o != nil {
			o(&op)
		}
	}

	r := &Redis{
		ctx: context.Background(),
	}
	r.client = redis.NewClient(&op).WithContext(r.ctx)
	err := r.connect()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Redis cache implement
type Redis struct {
	client *redis.Client
	mutex  *redislock.Client
	ctx    context.Context
}

func (r *Redis) WithContext(ctx context.Context) *Redis {
	r.client = r.client.WithContext(ctx)
	r.ctx = ctx
	return r
}

// connect connect test
func (r *Redis) connect() error {
	var err error
	_, err = r.client.Ping(r.ctx).Result()
	return err
}

func (r *Redis) Shutdown() {
	r.client.Shutdown(r.ctx)
}

// GetClient 暴露原生client
func (r *Redis) GetClient() *redis.Client {
	return r.client
}

func (r *Redis) Ping() error {
	_, err := r.GetClient().Ping(r.ctx).Result()
	return err
}

//Keys 根据正则获取keys
func (r *Redis) Keys(key string) ([]string, error) {
	if key == "*" {
		return nil, errors.New("can not search key：* use this redis client")
	}
	return r.client.Keys(r.ctx, key).Result()
}

//Type 获取key对应值得类型
func (r *Redis) Type(key string) (string, error) {
	return r.client.Type(r.ctx, key).Result()
}

//Del 删除缓存项
func (r *Redis) Del(keys ...string) (int64, error) {
	return r.client.Del(r.ctx, keys...).Result()
}

//Exists 检测缓存项是否存在
func (r *Redis) Exists(keys ...string) (int64, error) {
	return r.client.Exists(r.ctx, keys...).Result()
}

//Expire 方法是设置某个时间段(time.Duration)后过期
func (r *Redis) Expire(key string, expire time.Duration) (bool, error) {
	return r.client.Expire(r.ctx, key, expire).Result()
}

//ExpireAt 方法是在某个时间点(time.Time)过期失效
func (r *Redis) ExpireAt(key string, expireAt time.Time) (bool, error) {
	return r.client.ExpireAt(r.ctx, key, expireAt).Result()
}

//TTL 方法可以获取某个键的剩余有效期
func (r *Redis) TTL(key string) (time.Duration, error) {
	return r.client.TTL(r.ctx, key).Result()
}

//PTTL 获取有效期
func (r *Redis) PTTL(key string) (time.Duration, error) {
	return r.client.PTTL(r.ctx, key).Result()
}

// Get from key -- type string
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

//Set 设置Key = Value
func (r *Redis) Set(key string, val interface{}) (string, error) {
	return r.client.Set(r.ctx, key, val, 0).Result()
}

//SetWithExpire 设置Key=Value并指定 n 秒后过期
func (r *Redis) SetWithExpire(key string, val interface{}, expire int) (string, error) {
	return r.client.Set(r.ctx, key, val, time.Duration(expire)*time.Second).Result()
}

//SetEX 设置Key=Value并指定过期时间
func (r *Redis) SetEX(key string, val interface{}, expire time.Duration) (string, error) {
	return r.client.SetEX(r.ctx, key, val, expire).Result()
}

//SetNX 设置Key=Value并指定过期时间
func (r *Redis) SetNX(key string, val interface{}, expire time.Duration) (bool, error) {
	return r.client.SetNX(r.ctx, key, val, expire).Result()
}

//GetRange 字符串截取
func (r *Redis) GetRange(key string, start, end int64) (string, error) {
	return r.client.GetRange(r.ctx, key, start, end).Result()
}

//Incr 增加+1
func (r *Redis) Incr(key string) (int64, error) {
	return r.client.Incr(r.ctx, key).Result()
}

//IncrBy 按指定步长增加
func (r *Redis) IncrBy(key string, incrCount int64) (int64, error) {
	return r.client.IncrBy(r.ctx, key, incrCount).Result()
}

//Decr 减少-1
func (r *Redis) Decr(key string) (int64, error) {
	return r.client.Decr(r.ctx, key).Result()
}

//DecrBy 按指定步长减少
func (r *Redis) DecrBy(key string, decrCount int64) (int64, error) {
	return r.client.DecrBy(r.ctx, key, decrCount).Result()
}

//Append 表示往字符串后面追加元素，返回值是字符串的总长度
func (r *Redis) Append(key string, val string) (int64, error) {
	return r.client.Append(r.ctx, key, val).Result()
}

//StrLen 方法可以获取字符串的长度
func (r *Redis) StrLen(key string) (int64, error) {
	return r.client.StrLen(r.ctx, key).Result()
}

//LPush 方法将数据从左侧压入链表
func (r *Redis) LPush(key string, values ...interface{}) (int64, error) {
	return r.client.LPush(r.ctx, key, values...).Result()
}

//RPush 方法将数据从右侧压入链表
func (r *Redis) RPush(key string, values ...interface{}) (int64, error) {
	return r.client.RPush(r.ctx, key, values...).Result()
}

//LInsert 在某个位置插入新元素
func (r *Redis) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	return r.client.LInsert(r.ctx, key, op, pivot, value).Result()
}

//LSet 设置某个元素的值
func (r *Redis) LSet(key string, index int64, value interface{}) (string, error) {
	return r.client.LSet(r.ctx, key, index, value).Result()
}

//LLen 获取链表元素个数
func (r *Redis) LLen(key string) (int64, error) {
	return r.client.LLen(r.ctx, key).Result()
}

//LIndex 获取链表下标对应的元素
func (r *Redis) LIndex(key string, index int64) (string, error) {
	return r.client.LIndex(r.ctx, key, index).Result()
}

//LRange 获取某个选定范围的元素集
func (r *Redis) LRange(key string, start, stop int64) ([]string, error) {
	return r.client.LRange(r.ctx, key, start, stop).Result()
}

//LPop 从链表左侧弹出数据
func (r *Redis) LPop(key string) (string, error) {
	return r.client.LPop(r.ctx, key).Result()
}

//RPop 从链表右侧弹出数据
func (r *Redis) RPop(key string) (string, error) {
	return r.client.RPop(r.ctx, key).Result()
}

//LRem 根据值移除元素
func (r *Redis) LRem(key string, count int64, value interface{}) (int64, error) {
	return r.client.LRem(r.ctx, key, count, value).Result()
}

//SAdd 添加元素
func (r *Redis) SAdd(key string, count int64, value interface{}) (int64, error) {
	return r.client.SAdd(r.ctx, key, count, value).Result()
}

//SPop 随机获取一个元素
func (r *Redis) SPop(key string) (string, error) {
	return r.client.SPop(r.ctx, key).Result()
}

//SRem 删除集合里指定的值
func (r *Redis) SRem(keys string, members ...interface{}) (int64, error) {
	return r.client.SRem(r.ctx, keys, members).Result()
}

//SMembers 获取所有成员
func (r *Redis) SMembers(key string) ([]string, error) {
	return r.client.SMembers(r.ctx, key).Result()
}

//SIsMember 判断元素是否在集合中
func (r *Redis) SIsMember(key string, item interface{}) (bool, error) {
	return r.client.SIsMember(r.ctx, key, item).Result()
}

//SCard 获取集合元素个数
func (r *Redis) SCard(key string) (int64, error) {
	return r.client.SCard(r.ctx, key).Result()
}

//SUnion 并集
func (r *Redis) SUnion(key string) ([]string, error) {
	return r.client.SUnion(r.ctx, key).Result()
}

//SDiff 差集
func (r *Redis) SDiff(key string) ([]string, error) {
	return r.client.SDiff(r.ctx, key).Result()
}

//SInter 交集
func (r *Redis) SInter(key string) ([]string, error) {
	return r.client.SInter(r.ctx, key).Result()
}

//ZAdd 添加元素
func (r *Redis) ZAdd(key string, zItems ...*redis.Z) (int64, error) {
	return r.client.ZAdd(r.ctx, key, zItems...).Result()
}

//ZIncrBy 增加元素分值
func (r *Redis) ZIncrBy(key string, incr float64, item string) (float64, error) {
	return r.client.ZIncrBy(r.ctx, key, incr, item).Result()
}

//ZRange 获取根据score排序后的数据段
func (r *Redis) ZRange(key string, start, stop int64) ([]string, error) {
	return r.client.ZRange(r.ctx, key, start, stop).Result()
}

//ZRevRange 获取根据score排序后的数据段
func (r *Redis) ZRevRange(key string, start, stop int64) ([]string, error) {
	return r.client.ZRevRange(r.ctx, key, start, stop).Result()
}

//ZRangeByScore 获取score过滤后排序的数据段
func (r *Redis) ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	return r.client.ZRangeByScore(r.ctx, key, opt).Result()
}

//ZCount 获取区间内元素个数
func (r *Redis) ZCount(key string, min, max int) (int64, error) {
	return r.client.ZCount(r.ctx, key, cast.ToString(min), cast.ToString(max)).Result()
}

//ZScore 获取元素的score
func (r *Redis) ZScore(key, item string) (float64, error) {
	return r.client.ZScore(r.ctx, key, item).Result()
}

//ZRank 方法是返回元素在集合中的升序排名情况，从0开始
func (r *Redis) ZRank(key, item string) (int64, error) {
	return r.client.ZRank(r.ctx, key, item).Result()
}

//ZRevRank 方法是返回元素在集合中的降序排名情况
func (r *Redis) ZRevRank(key, item string) (int64, error) {
	return r.client.ZRevRank(r.ctx, key, item).Result()
}

//ZRem 方法支持通过元素的值来删除元素
func (r *Redis) ZRem(key, item string) (int64, error) {
	return r.client.ZRem(r.ctx, key, item).Result()
}

//ZRemRangeByRank 根据排名来删除
func (r *Redis) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	return r.client.ZRemRangeByRank(r.ctx, key, start, stop).Result()
}

//ZRemRangeByScore 根据分值区间来删除
func (r *Redis) ZRemRangeByScore(key string, min, max int) (int64, error) {
	return r.client.ZRemRangeByScore(r.ctx, key, cast.ToString(min), cast.ToString(max)).Result()
}

// HSet 设置Hash
func (r *Redis) HSet(hk string, kv map[string]interface{}) (int64, error) {
	return r.client.HSet(r.ctx, hk, kv).Result()
}

// HMSet 批量设置
func (r *Redis) HMSet(hk string, kv map[string]interface{}) (bool, error) {
	return r.client.HMSet(r.ctx, hk, kv).Result()
}

//HGet 获取某个元素
func (r *Redis) HGet(hk, key string) (string, error) {
	return r.client.HGet(r.ctx, hk, key).Result()
}

//HGetAll 获取全部元素
func (r *Redis) HGetAll(hk string) (map[string]string, error) {
	return r.client.HGetAll(r.ctx, hk).Result()
}

//HDel 删除某个元素或某几个元素
func (r *Redis) HDel(hk string, keys ...string) (int64, error) {
	return r.client.HDel(r.ctx, hk, keys...).Result()
}

//HExists 判断元素是否存在
func (r *Redis) HExists(hk, key string) (bool, error) {
	return r.client.HExists(r.ctx, hk, key).Result()
}

//HLen 获取长度
func (r *Redis) HLen(hk string) (int64, error) {
	return r.client.HLen(r.ctx, hk).Result()
}
