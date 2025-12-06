package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func New(opt *Option) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     opt.GetAddr(),
		Username: opt.Username,
		Password: opt.Password,
		DB:       opt.DB,
	})

	// Test connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewCluster(opt *ClusterOption) (*redis.ClusterClient, error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    opt.Addrs,
		Username: opt.Username,
		Password: opt.Password,
	})

	// Test connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
