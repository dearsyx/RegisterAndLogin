package repository

import (
	"encoding/json"
	"net/http"

	"goprojects.com/simple_regist_login/pkg/code"

	"goprojects.com/simple_regist_login/pkg/util"

	"goprojects.com/simple_regist_login/internal/dto/errs"
	"goprojects.com/simple_regist_login/internal/token"

	"goprojects.com/simple_regist_login/internal/domain/user/entity"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"goprojects.com/simple_regist_login/pkg/db"
)

type user struct {
	mysqlDB *gorm.DB
	redisDB *redis.Client
}

func NewRepoUser() User {
	return &user{
		mysqlDB: db.DB,
		redisDB: db.Redis,
	}
}

//IsTelephoneExist 判断手机号是否已经存在
func (u *user) IsTelephoneExist(tele string, userEntity *entity.User) bool {
	u.mysqlDB.Where("tele = ?", tele).First(userEntity)
	if userEntity.ID == 0 {
		return false
	}
	return true
}

func (u *user) SetToken(tokenKey string, tokenValue interface{}) (err interface{}) {
	tokenString, errMar := json.Marshal(tokenValue)
	if errMar != nil {
		return errs.NewDataError(http.StatusServiceUnavailable, code.ServerError, errMar.Error())
	}
	_ = u.redisDB.Set(tokenKey, string(tokenString), 0)
	return nil
}

func (u *user) GetToken(tokenKey string) (tokenValue *token.UserToken, err interface{}) {
	resp, errOrigin := u.redisDB.Get(tokenKey).Result()
	if errOrigin != nil {
		return nil, errs.NewDataError(http.StatusServiceUnavailable, code.ServerError, errOrigin.Error())
	}
	tokenObject := &token.UserToken{}
	errOrigin = json.Unmarshal([]byte(resp), tokenObject)
	if errOrigin != nil {
		return nil, errs.NewDataError(http.StatusServiceUnavailable, code.ServerError, errOrigin.Error())
	}
	return tokenObject, nil
}

func (u *user) CreateUserInMysql(userEntity *entity.User) {
	userEntity.Pass = util.HashPassword(userEntity.Pass)
	u.mysqlDB.Create(userEntity)
}

func (u *user) FindUser(query string, value interface{}, userEntity *entity.User) {
	u.mysqlDB.Where(query, value).First(userEntity)
}
