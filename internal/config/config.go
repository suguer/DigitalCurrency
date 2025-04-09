package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

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
