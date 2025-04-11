package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type EthConfig struct {
	Chain    string
	Enable   int
	ApiKey   string
	Node     string
	Contract []ContractConfig
}
type ContractConfig struct {
	Token   string
	Address string
	Decimal int
}

func (c *EthConfig) ContractMap() map[string]ContractConfig {
	m := make(map[string]ContractConfig)
	for _, v := range c.Contract {
		m[v.Address] = v
	}
	return m
}

type RedisConf struct {
	Host     string
	Port     int
	Database int
	Password string
}
type MysqlConf struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Charset  string
}

// Config 配置结构体
type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	}
	Database struct {
		Driver string `yaml:"driver"`
		Source string `yaml:"source"`
	}
	Blockchain struct {
		BlockInterval int `yaml:"block_interval"`
	}
	Log struct {
		Level    string `yaml:"level"`
		File     string `yaml:"file"`
		Console  bool   `yaml:"console"`
		Rotation struct {
			MaxSize    int `yaml:"max_size"`
			MaxAge     int `yaml:"max_age"`
			MaxBackups int `yaml:"max_backups"`
		} `yaml:"rotation"`
	}
	Tron  EthConfig
	Redis RedisConf
	Mysql MysqlConf
}

// Load 加载配置文件
func Load(file string) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
