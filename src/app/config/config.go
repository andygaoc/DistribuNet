package config

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/toolkits/file"
	"os"
	"sync"
)

type DbInfo struct {
	Host           string
	Port           string
	Database       string
	Username       string
	Password       string
	DbMaxIdleConns int
	DbMaxOpenConns int
}

// 应用配置
type Server struct {
	Host string
	Port string
}

type Config struct {
	Server *Server
	DbInfo *DbInfo
	LogDir string
}

var (
	lock   sync.RWMutex
	config *Config
)

func ParseConfig(cfg string) {
	if cfg == "" {
		logs.Critical("没有指定配置文件")
		os.Exit(1)
	}

	if !file.IsExist(cfg) {
		logs.Critical("配置文件:", cfg, "不存在")
		os.Exit(1)
	}

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		logs.Critical("读取配置文件:", cfg, "失败:", err)
		os.Exit(1)
	}

	var c Config
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		logs.Critical("解析配置文件:", cfg, "失败:", err)
		os.Exit(1)
	}

	lock.Lock()
	defer lock.Unlock()
	config = &c
}

func GetConfig() *Config {
	lock.RLock()
	defer lock.RUnlock()
	return config
}
