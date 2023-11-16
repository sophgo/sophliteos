package config

import (
	"algoliteos/logger"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	configurePath = "/etc/sophliteos/config" // 配置文件所在目录
)

var Conf Config
var JsonConf Config
var Event Config

func LoadConfig() {
	Conf = Config{}
	Conf.name = "algoliteos"
	Conf.v = viper.New()

	v := Conf.Application.v
	v.AddConfigPath(configurePath)
	v.SetConfigName(Conf.name)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil { // viper解析配置文件
		logger.Info("load config path: %s, error: %s", configurePath, err)
	}

	Event = Config{}
	Event.name = "events"
	Event.v = viper.New()

	k := Event.Application.v
	k.AddConfigPath(configurePath)
	k.SetConfigName(Event.name)
	k.SetConfigType("yaml")

	if err := k.ReadInConfig(); err != nil { // viper解析配置文件
		logger.Info("load config path: %s, error: %s", configurePath, err)
	}

	JsonConf = Config{}
	JsonConf.name = "event"
	JsonConf.v = viper.New()

	j := JsonConf.Application.v
	j.AddConfigPath(configurePath)
	j.SetConfigName(JsonConf.name)
	j.SetConfigType("json")

	if err := j.ReadInConfig(); err != nil { // viper解析配置文件
		logger.Info("load config path: %s, error: %s", configurePath, err)
	}

	exists := true
	if exists {
		// 监控配置文件变化并热加载程序
		watchConfig(v, func(in fsnotify.Event) {
			logger.Info("Config file changed: %s", in.Name)
			v.ReadInConfig()
		})

		watchConfig(k, func(in fsnotify.Event) {
			logger.Info("Config file changed: %s", in.Name)
			k.ReadInConfig()
		})

		watchConfig(j, func(in fsnotify.Event) {
			logger.Info("Config file changed: %s", in.Name)
			j.ReadInConfig()
		})
	}

}

type Application struct {
	name string
	v    *viper.Viper
}

func (c *Application) GetName() string {
	return c.name
}

func (c *Application) GetViper() *viper.Viper {
	return c.v
}

type Config struct {
	Application
	rwMutex sync.RWMutex
}

func (sc *Config) RLock() {
	sc.rwMutex.RLock()
}

func (sc *Config) RUnlock() {
	sc.rwMutex.RUnlock()
}

func (sc *Config) Lock() {
	sc.rwMutex.Lock()
}

func (sc *Config) Unlock() {
	sc.rwMutex.Unlock()
}

// 监控配置文件变化并热加载程序
func watchConfig(v *viper.Viper, callback func(in fsnotify.Event)) {
	v.OnConfigChange(callback)
	v.WatchConfig()
}
