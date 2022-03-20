package main

import (
	"goprojects.com/simple_regist_login/internal/handler"
	"goprojects.com/simple_regist_login/pkg/config"
	"goprojects.com/simple_regist_login/pkg/db"
)

func main() {
	var err error

	// 读取配置文件
	err = config.InitConfig()
	if err != nil {
		panic(err)
	}

	// 初始化Mysql数据库
	err = db.MysqlInit()
	if err != nil {
		panic(err)
	}

	// 初始化Redis数据库
	err = db.RedisInit()
	if err != nil {
		panic(err)
	}

	handler.StartApi()
}
