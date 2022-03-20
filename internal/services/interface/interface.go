package _interface

import "github.com/gin-gonic/gin"

//Migrator 自动迁移
type Migrator interface {
	Migrate()
}

// Register 路由注册器
type Register interface {
	Regist(r *gin.Engine)
}
