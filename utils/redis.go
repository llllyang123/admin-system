package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var Redis *RedisService

// RedisClient 是一个Redis客户端接口
type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	// ... 其他你需要的Redis方法
}

// RedisService 是RedisClient的实现
type RedisService struct {
	client *redis.Client
}

func RedisServiceInit() {
	redisOpts := &redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // Redis密码
		DB:       0,                // 使用的数据库编号
	}
	Redis = NewRedisService(redisOpts)
}

// NewRedisService 创建一个新的RedisService实例
func NewRedisService(opts *redis.Options) *RedisService {
	return &RedisService{
		client: redis.NewClient(opts),
	}
}

// Set 实现RedisClient接口的Set方法
func (rs *RedisService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return rs.client.Set(ctx, key, value, expiration)
}

// Get 实现RedisClient接口的Get方法
func (rs *RedisService) Get(ctx context.Context, key string) *redis.StringCmd {
	return rs.client.Get(ctx, key)
}

// Del 实现RedisClient接口的Del方法
func (rs *RedisService) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return rs.client.Del(ctx, keys...)
}
