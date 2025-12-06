package cache

import "fmt"

// Option represents Redis connection configuration
type Option struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
}

// GetAddr returns the Redis address in host:port format
func (opt Option) GetAddr() string {
	return fmt.Sprintf("%s:%d", opt.Host, opt.Port)
}

// GetDefaultOption returns default Redis configuration
func GetDefaultOption() *Option {
	return &Option{
		Host:     "127.0.0.1",
		Port:     6379,
		Username: "",
		Password: "",
		DB:       0,
	}
}

// ClusterOption represents Redis cluster configuration
type ClusterOption struct {
	Addrs    []string
	Username string
	Password string
}

// GetDefaultClusterOption returns default Redis cluster configuration
func GetDefaultClusterOption() *ClusterOption {
	return &ClusterOption{
		Addrs:    []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
		Username: "",
		Password: "",
	}
}
