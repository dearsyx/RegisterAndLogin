package config

import (
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

// MysqlConfig mysql配置文件结构体
type MysqlConfig struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Database string `ini:"database"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
	Charset  string `ini:"charset"`
}

//RedisConfig redis配置文件结构体
type RedisConfig struct {
	Address  string `ini:"address"`
	Password string `ini:"password"`
	DB       int    `ini:"DB"`
}

// StaticConfig 静态文件配置
type StaticConfig struct {
	TemplatePath   string `ini:"template_path"`
	JavascriptPath string `ini:"js_path"`
	CssPath        string `ini:"css_path"`
}

// Config 配置文件结构体
type Config struct {
	MysqlConfig  `ini:"mysql"`
	RedisConfig  `ini:"redis"`
	StaticConfig `ini:"static"`
}

var AppConfig Config

// InitConfig 初始化配置文件
func InitConfig() (err error) {
	var cfg *ini.File
	cfg, err = ini.Load("./config.ini")
	if err != nil {
		return err
	}
	err = cfg.MapTo(&AppConfig)
	if err != nil {
		return err
	}
	logrus.Info("config file load success!")
	return nil
}
