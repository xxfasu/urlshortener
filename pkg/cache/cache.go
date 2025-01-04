package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/coocood/freecache"
	"github.com/redis/go-redis/v9"
	"github.com/xxfasu/urlshortener/pkg/logs"
	"go.uber.org/zap"
	"time"
)

type cache struct {
	LocalCache *freecache.Cache
	RedisCache redis.UniversalClient
}

type StatusInfo struct {
	HitRate       string // 获取缓存命中率
	HitCount      int64  // 获取命中次数
	MissCount     int64  // 获取未命中次数
	EntryCount    int64  // 获取当前缓存条目数
	EvacuateCount int64  // 获取被清理的条目数
	TotalRequests int64  // 总请求次数
}

type KeyInfo struct {
	TTL   int64  // 过期时间
	Value string // 缓存的值
}

type FetcherFunc func(ctx context.Context, key string) (string, error)

func (f FetcherFunc) Fetch(ctx context.Context, key string) (string, error) {
	return f(ctx, key)
}

func InitLocalCache(client redis.UniversalClient) Cache {
	// 创建一个 10MB 大小的缓存
	cacheSize := 100 * 1024 * 1024 // 10MB
	localCache := freecache.NewCache(cacheSize)
	return &cache{LocalCache: localCache, RedisCache: client}
}

// GetCacheStatus 封装的方法：获取缓存命中率、缓存命中数、总请求数
func (c *cache) GetCacheStatus() StatusInfo {
	hitRate := c.LocalCache.HitRate()             // 获取缓存命中率
	hitCount := c.LocalCache.HitCount()           // 获取命中次数
	missCount := c.LocalCache.MissCount()         // 获取未命中次数
	evacuateCount := c.LocalCache.EvacuateCount() // 获取被清理的条目数
	entryCount := c.LocalCache.EntryCount()       // 获取当前缓存条目数
	totalRequests := hitCount + missCount         // 计算总请求次数
	return StatusInfo{
		HitRate:       fmt.Sprintf("%.2f%%", hitRate*100),
		HitCount:      hitCount,
		MissCount:     missCount,
		TotalRequests: totalRequests,
		EntryCount:    entryCount,
		EvacuateCount: evacuateCount,
	}
}

// GetKeyStatus 封装的方法：获取对应缓存的值和过期时间
func (c *cache) GetKeyStatus(key string) KeyInfo {
	value, ttl, err := c.LocalCache.GetWithExpiration([]byte(key))
	if err != nil {
		logs.Log.Error("Error getting key:", zap.Error(err))
	}

	return KeyInfo{
		TTL:   int64(ttl) - time.Now().Unix(),
		Value: string(value),
	}
}

// GetLocal 获取本地缓存中的值
func (c *cache) GetLocal(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key is empty")
	}
	if c == nil || c.LocalCache == nil {
		return "", nil
	}
	value, err := c.LocalCache.Get([]byte(key))
	if err != nil {
		return "", err
	}
	return string(value), nil
}

// GetCache 先从本地缓存获取，若不存在，则从redis获取
func (c *cache) GetCache(ctx context.Context, key string) (string, error) {
	value, err := c.GetLocal(key)
	if err != nil {
		if !errors.Is(err, freecache.ErrNotFound) {
			return "", err
		}
	}
	if len(value) != 0 {
		return value, nil
	}
	pipeline := c.RedisCache.Pipeline()
	getCmd := pipeline.Get(ctx, key)
	ttlCmd := pipeline.TTL(ctx, key)
	_, err = pipeline.Exec(ctx)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return "", err
		}
		return "", nil
	}
	value, err = getCmd.Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return "", err
		}
		return "", nil
	}
	ttl, err := ttlCmd.Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return "", err
		}
		return "", nil
	}
	duration := ttl / 3
	if int(duration.Seconds()) > 0 {
		err = c.SetLocal(key, value, duration)
		if err != nil {
			return "", errors.New("cache set err")
		}
	}
	return value, nil
}

// GetCacheOrElse 先从本地缓存获取，若不存在，则从redis获取，若redis也不存在，则调用fetcher
func (c *cache) GetCacheOrElse(ctx context.Context, key string, ttl time.Duration, fetcher Fetcher) (string, error) {
	value, err := c.GetLocal(key)
	if err != nil {
		if !errors.Is(err, freecache.ErrNotFound) {
			return "", err
		}
	}
	if len(value) != 0 {
		return value, nil
	}
	pipeline := c.RedisCache.Pipeline()
	getCmd := pipeline.Get(ctx, key)
	ttlCmd := pipeline.TTL(ctx, key)
	_, err = pipeline.Exec(ctx)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return "", err
		}
	}
	remainingTime, err := ttlCmd.Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return "", err
		}
	}
	value, err = getCmd.Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			value, err = fetcher.Fetch(ctx, key)
			if err != nil {
				return "", err
			}
			c.SetCache(ctx, key, value, ttl)
			return value, nil
		} else {
			return "", err
		}
	}
	duration := remainingTime / 3
	if int(duration.Seconds()) > 0 {
		err = c.SetLocal(key, value, duration)
		if err != nil {
			return "", errors.New("cache set err")
		}
	}
	return value, nil
}

// SetLocal 设置本地缓存值
func (c *cache) SetLocal(key, value string, ttl time.Duration) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if c == nil || c.LocalCache == nil {
		return nil
	}
	err := c.LocalCache.Set([]byte(key), []byte(value), int(ttl.Seconds()))
	if err != nil {
		logs.Log.Error("cache set err:", zap.Error(err))
		return errors.New("cache set err")
	}
	return nil
}

// SetCache 设置redis和本地缓存值
func (c *cache) SetCache(ctx context.Context, key, value string, ttl time.Duration) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	duration := ttl / 3
	if int(duration.Seconds()) > 0 {
		err := c.SetLocal(key, value, duration)
		if err != nil {
			return err
		}
	}
	result, err := c.RedisCache.Set(ctx, key, value, ttl).Result()
	logs.Log.Info("redis2 set", zap.String("result", result))
	if err != nil {
		return err
	}
	return nil
}

// DelLocal 删除本地缓存值
func (c *cache) DelLocal(key string) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if c == nil || c.LocalCache == nil {
		return nil
	}
	c.LocalCache.Del([]byte(key))
	return nil
}

// DelCache 删除redis和本地缓存值
func (c *cache) DelCache(key string) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	err := c.DelLocal(key)
	if err != nil {
		return err
	}
	result, err := c.RedisCache.Del(context.Background(), key).Result()
	logs.Log.Info("redis2 del num", zap.Int64("result", result))
	if err != nil {
		return err
	}
	return nil
}
