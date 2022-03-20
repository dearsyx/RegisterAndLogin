package services

import (
	"net/http"

	"goprojects.com/simple_regist_login/pkg/code"

	"goprojects.com/simple_regist_login/pkg/util"

	"goprojects.com/simple_regist_login/internal/dto/errs"

	"goprojects.com/simple_regist_login/internal/token"

	"goprojects.com/simple_regist_login/internal/domain/user/entity"
	"goprojects.com/simple_regist_login/internal/domain/user/repository"
)

type user struct {
	repoUser repository.User
}

type User interface {
	RegistDataVerify(name, tele, pass string, userEntity *entity.User) (is bool, err interface{})
	CreateUser(name, tele, pass string)
	SetUserToken(token token.UserToken) (tokenKey string, err interface{})
	GetUserToken(tokenKey string) (userToken *token.UserToken, err interface{})
	FindUserByTelephone(tele string, userEntity *entity.User)
	LoginDataVerify(tele string, pass string, userEntity *entity.User) (err interface{})
}

func NewServiceUser(repoUser repository.User) User {
	return &user{repoUser: repoUser}
}

func (u *user) RegistDataVerify(name, tele, pass string, userEntity *entity.User) (is bool, err interface{}) {
	if name == "" || pass == "" || tele == "" {
		return false, errs.NewDataError(http.StatusBadRequest, code.InputDataError, "参数不能为空")
	}
	if len(tele) != 11 {
		return false, errs.NewDataError(http.StatusBadRequest, code.InputDataError, "电话号码格式错误")
	}
	if u.repoUser.IsTelephoneExist(tele, userEntity) {
		return false, errs.NewDataError(http.StatusBadRequest, code.InputDataError, "电话号码已被注册")
	}
	return true, nil
}

func (u *user) CreateUser(name, tele, pass string) {
	u.repoUser.CreateUserInMysql(&entity.User{
		Name: name,
		Pass: pass,
		Tele: tele,
	})
}

func (u *user) SetUserToken(userToken token.UserToken) (tokenKey string, err interface{}) {
	tokenKey = token.GenRandTokenKey()
	err = u.repoUser.SetToken(tokenKey, userToken)
	if err != nil {
		return "", err
	}
	return tokenKey, nil
}

func (u *user) GetUserToken(tokenKey string) (userToken *token.UserToken, err interface{}) {
	userToken, err = u.repoUser.GetToken(tokenKey)
	if err != nil {
		return nil, err
	}
	return userToken, err
}

func (u *user) FindUserByTelephone(tele string, userEntity *entity.User) {
	query := "tele = ?"
	u.repoUser.FindUser(query, tele, userEntity)
}

func (u *user) LoginDataVerify(tele string, pass string, userEntity *entity.User) (err interface{}) {
	u.FindUserByTelephone(tele, userEntity)
	if userEntity.ID == 0 {
		return errs.NewDataError(http.StatusBadRequest, code.InputDataError, "用户不存在")
	}
	if !util.VerifyPassword(userEntity.Pass, pass) {
		return errs.NewDataError(http.StatusBadRequest, code.InputDataError, "用户名或密码错误")
	}
	return nil
}
