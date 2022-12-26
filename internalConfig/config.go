package internalConfig

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

// Configuration 项目配置
type Configuration struct {
	// gtp apikey
	ApiKey string `json:"api_key"`
	// 自动通过好友
	AutoPass bool `json:"auto_pass"`
	// 会话超时时间
	SessionTimeout time.Duration `json:"session_timeout"`
	// 只接受指定群组的消息
	AcceptGroups []string `json:"accept_groups"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{
			SessionTimeout: 1,
		}
		// 文件位置可能有变化
		f, err := os.Open("./internalConfig/config.json") // 读取文件
		if err != nil {
			log.Fatalf("open config err: %v", err)
			return
		}
		defer f.Close()
		encoder := json.NewDecoder(f)
		err = encoder.Decode(config) // 将 config.json 中的文件序列化到config结构体中
		if err != nil {
			log.Fatalf("decode config err: %v", err)
			return
		}

		// 如果环境变量有配置，优先读取环境变量
		ApiKey := os.Getenv("ApiKey")
		AutoPass := os.Getenv("AutoPass")
		SessionTimeout := os.Getenv("SessionTimeout")
		AcceptGroups := os.Getenv("AcceptGroups")
		if ApiKey != "" {
			config.ApiKey = ApiKey
		}
		if AutoPass == "true" {
			config.AutoPass = true
		}
		if SessionTimeout != "" {
			duration, err := time.ParseDuration(SessionTimeout)
			if err != nil {
				log.Fatalf("config decode session timeout err: %v ,get is %v", err, SessionTimeout)
				return
			}
			config.SessionTimeout = duration
		}
		if AcceptGroups != "" {
			err := json.Unmarshal([]byte(AcceptGroups), &config.AcceptGroups)
			if err != nil {
				log.Fatalf("Unmarshal AcceptGroups error:%v", err)
			}
		}
	})
	return config
}
