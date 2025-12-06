package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const (
	redisContextKey        = "__redis_instance"
	redisClusterContextKey = "__redis_cluster_instance"
)

// Context stores a Redis client in the context
func Context(ctx context.Context, client *redis.Client) context.Context {
	return context.WithValue(ctx, redisContextKey, client)
}

// FromContext retrieves a Redis client from the context
func FromContext(ctx context.Context) *redis.Client {
	client, ok := ctx.Value(redisContextKey).(*redis.Client)
	if !ok {
		return nil
	}
	return client
}

// ClusterContext stores a Redis cluster client in the context
func ClusterContext(ctx context.Context, client *redis.ClusterClient) context.Context {
	return context.WithValue(ctx, redisClusterContextKey, client)
}

// ClusterFromContext retrieves a Redis cluster client from the context
func ClusterFromContext(ctx context.Context) *redis.ClusterClient {
	client, ok := ctx.Value(redisClusterContextKey).(*redis.ClusterClient)
	if !ok {
		return nil
	}
	return client
}
