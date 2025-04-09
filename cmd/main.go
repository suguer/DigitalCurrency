package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"DigitalCurrency/internal/config"
	"DigitalCurrency/internal/router"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "etc/config.yaml", "配置文件路径")
}

func main() {
	flag.Parse()

	// 获取工作目录
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取工作目录失败: %v", err)
	}

	// 加载配置文件
	cfg, err := config.Load(filepath.Join(wd, configFile))
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	// 初始化路由
	r := router.NewRouter()

	// 启动服务
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("服务启动在 http://localhost%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}