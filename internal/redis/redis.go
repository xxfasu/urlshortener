package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/xxfasu/urlshortener/internal/conf"

	"time"
)

func InitRedis() (redis.UniversalClient, error) {
	var client redis.UniversalClient
	// 创建 Redis 客户端
	if conf.Config.Redis.UseCluster {
		// 使用集群模式
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        conf.Config.Redis.ClusterAddrs,
			Password:     conf.Config.Redis.Password,
			Username:     conf.Config.Redis.Username,
			PoolSize:     16, // 连接池大小
			MinIdleConns: 3,  // 最小空闲连接数
		})
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:         conf.Config.Redis.Addr,
			Password:     conf.Config.Redis.Password, // 如果没有密码，使用空字符串
			DB:           conf.Config.Redis.DB,
			Username:     conf.Config.Redis.Username,
			PoolSize:     16, // 连接池大小
			MinIdleConns: 3,  // 最小空闲连接数
		})
	}
	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}
	return client, nil
}
