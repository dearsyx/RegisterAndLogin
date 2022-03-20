package entity

import (
	"github.com/jinzhu/gorm"
	_interface "goprojects.com/simple_regist_login/internal/services/interface"
	"goprojects.com/simple_regist_login/pkg/db"
)

type User struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	Pass string `gorm:"column:pass" json:"pass"`
	Tele string `gorm:"column:tele" json:"tele"`
}

//NewEntityUser 返回一个entity.User实体
func NewEntityUser(name, pass, tele string) *User {
	return &User{
		Name: name,
		Pass: pass,
		Tele: tele,
	}
}

func (u *User) Migrate() {
	db.DB.AutoMigrate(&User{})
}

//NewEntityUserMigrator 返回一个用于Migrate User的自动迁移接口
func NewEntityUserMigrator() _interface.Migrator {
	return &User{}
}
