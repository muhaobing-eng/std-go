package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/muhaobing-eng/std-go/go-common/cache"
	"github.com/muhaobing-eng/std-go/go-common/database"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig    `yaml:"server"`
	Database *DatabaseConfig `yaml:"database"`
	Cache    *CacheConfig    `yaml:"cache"`
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

	restServerPrefix = "restserver"
)

var globalConfig *Config

func Get() *Config {
	if globalConfig == nil {
		panic("config not loaded, call Load() first")
	}
	return globalConfig
}

// Load loads configuration from file specified by environment variable or default path
func Load() (*Config, error) {
	conf := new(Config)
	if err := LoadWithPrefix(restServerPrefix, conf); err != nil {
		return nil, err
	}
	globalConfig = conf
	return globalConfig, nil
}

func LoadWithPrefix(prefix string, conf interface{}) error {
	if prefix == "" {
		return fmt.Errorf("prefix is empty")
	}

	configPath := os.Getenv(ConfigPathEnvName)
	if configPath == "" {
		configPath = defaultConfigPath
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %w", configPath, err)
	}

	var rawConfig map[string]interface{}
	if err = yaml.Unmarshal(data, &rawConfig); err != nil {
		return fmt.Errorf("failed to parse config file %s: %w", configPath, err)
	}
	prefixData, ok := rawConfig[prefix]
	if !ok {
		return fmt.Errorf("prefix '%s' not found in config file %s", prefix, configPath)
	}
	prefixBytes, err := yaml.Marshal(prefixData)
	if err != nil {
		return fmt.Errorf("failed to marshal prefix section: %w", err)
	}
	if err = yaml.Unmarshal(prefixBytes, conf); err != nil {
		return fmt.Errorf("failed to parse prefix section '%s': %w", prefix, err)
	}
	return nil
}
