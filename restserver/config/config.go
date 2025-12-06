package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"std-go/go-common/cache"
	"std-go/go-common/database"

	"gopkg.in/yaml.v3"
)

// RestserverConfig represents the main restserver configuration section
type RestserverConfig struct {
	Server   ServerConfig   `yaml:"restserver.server"`
	Database DatabaseConfig `yaml:"restserver.database"`
	Cache    CacheConfig    `yaml:"restserver.cache"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Log  string `yaml:"log"`
}

type DatabaseConfig struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Other    string `yaml:"other"`
}

type CacheConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func (dc *DatabaseConfig) GetDatabaseOption() *database.Option {
	return &database.Option{
		IP:       dc.IP,
		Port:     dc.Port,
		User:     dc.User,
		Password: dc.Password,
		DB:       dc.DB,
		Other:    dc.Other,
	}
}

func (cc *CacheConfig) GetCacheOption() *cache.Option {
	return &cache.Option{
		Host:     cc.Host,
		Port:     cc.Port,
		Username: cc.Username,
		Password: cc.Password,
		DB:       cc.DB,
	}
}

const (
	ConfigPathEnvName = "STD_REST_CONF_PATH"

	defaultConfigPath = "./conf/restserver.yaml"
)

var globalConfig *RestserverConfig

// Load loads configuration from file specified by environment variable or default path
func Load() (*RestserverConfig, error) {
	configPath := os.Getenv(ConfigPathEnvName)
	if configPath == "" {
		configPath = defaultConfigPath
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", configPath, err)
	}

	var config RestserverConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", configPath, err)
	}

	// Set default values if not specified
	if config.Server.Port == 0 {
		config.Server.Port = 8080
	}

	if config.Database.IP == "" {
		config.Database.IP = "127.0.0.1"
	}
	if config.Database.Port == 0 {
		config.Database.Port = 3306
	}
	if config.Database.User == "" {
		config.Database.User = "root"
	}
	if config.Database.Other == "" {
		config.Database.Other = "charset=utf8&parseTime=True&loc=Local"
	}

	if config.Cache.Host == "" {
		config.Cache.Host = "127.0.0.1"
	}
	if config.Cache.Port == 0 {
		config.Cache.Port = 6379
	}

	globalConfig = &config
	return &config, nil
}

func Get() *RestserverConfig {
	if globalConfig == nil {
		panic("config not loaded, call Load() first")
	}
	return globalConfig
}
