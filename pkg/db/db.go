package db

import (
	"fmt"

	"github.com/go-redis/redis"
	"goprojects.com/simple_regist_login/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var Redis *redis.Client

func MysqlInit() (err error) {
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		config.AppConfig.MysqlConfig.UserName,
		config.AppConfig.MysqlConfig.Password,
		config.AppConfig.MysqlConfig.Host,
		config.AppConfig.MysqlConfig.Port,
		config.AppConfig.MysqlConfig.Database,
		config.AppConfig.MysqlConfig.Charset,
	)
	// 打开数据库
	DB, err = gorm.Open("mysql", args)
	if err != nil {
		//fmt.Println("连接数据库失败")
		return err
	}

	// 检查数据库连接
	err = DB.DB().Ping()
	if err != nil {
		return err
	}

	return DB.Error
}

func RedisInit() (err error) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisConfig.Address,
		Password: config.AppConfig.RedisConfig.Password,
		DB:       config.AppConfig.RedisConfig.DB,
	})

	_, err = Redis.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
