package repository

import (
	"goprojects.com/simple_regist_login/internal/domain/user/entity"
	"goprojects.com/simple_regist_login/internal/token"
)

type User interface {
	IsTelephoneExist(tele string, userEntity *entity.User) bool
	SetToken(tokenKey string, tokenValue interface{}) (err interface{})
	GetToken(tokenKey string) (tokenValue *token.UserToken, err interface{})
	CreateUserInMysql(userEntity *entity.User)
	FindUser(query string, value interface{}, userEntity *entity.User)
}
