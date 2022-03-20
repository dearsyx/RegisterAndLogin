package handler

import (
	"net/http"
	"time"

	"goprojects.com/simple_regist_login/pkg/code"

	"goprojects.com/simple_regist_login/internal/middleware"

	"goprojects.com/simple_regist_login/internal/dto/errs"

	"github.com/gin-gonic/gin"
	"goprojects.com/simple_regist_login/internal/domain/user/entity"
	"goprojects.com/simple_regist_login/internal/services"
	"goprojects.com/simple_regist_login/internal/token"
)

type User struct {
	serviceUser services.User
}

func NewHandlerUser(svcUser services.User) *User {
	return &User{serviceUser: svcUser}
}

//UserRegist 用户注册路由
func (u *User) UserRegist(c *gin.Context) {
	HandlerResponse(c, func() interface{} {
		jsonData := make(map[string]interface{})
		_ = c.BindJSON(&jsonData)
		name := jsonData["name"].(string)
		tele := jsonData["tele"].(string)
		pass := jsonData["pass"].(string)
		resp, err := u.serviceUser.RegistDataVerify(name, tele, pass, &entity.User{})
		if !resp {
			return err
		}
		// 创建新用户
		u.serviceUser.CreateUser(name, tele, pass)
		// 设置token
		tokenKey, err := u.serviceUser.SetUserToken(token.UserToken{
			Info:       gin.H{"user_name": name, "tele_number": tele},
			Token:      `this is a test token`,
			CreateTime: time.Now().Unix(),
		})
		if err != nil {
			return errs.NewDataError(http.StatusServiceUnavailable, 1001, "token设置失败")
		}
		// 保存token到cookie
		c.SetCookie(token.CookieTokenKey, tokenKey, token.TokenMaxAge, "/", "", false, true)
		return gin.H{
			"code": code.RegistSuccess,
			"msg":  "注册成功",
		}
	})
}

func (u *User) UserLogin(c *gin.Context) {
	HandlerResponse(c, func() interface{} {
		userEntity := &entity.User{}
		jsonData := make(map[string]interface{})
		_ = c.BindJSON(&jsonData)
		tele := jsonData["tele"].(string)
		pass := jsonData["pass"].(string)
		// 验证
		err := u.serviceUser.LoginDataVerify(tele, pass, userEntity)
		if err != nil {
			return err
		}
		// 数据验证成功，设置token
		tokenKey, err := u.serviceUser.SetUserToken(token.UserToken{
			Info:       gin.H{"user_name": userEntity.Name, "tele_number": tele},
			Token:      `this is a test token`,
			CreateTime: time.Now().Unix(),
		})
		if err != nil {
			return errs.NewDataError(http.StatusServiceUnavailable, code.ServerError, "token设置失败")
		}
		// 保存token到cookie
		c.SetCookie(token.CookieTokenKey, tokenKey, token.TokenMaxAge, "/", "", false, true)
		return gin.H{
			"code": code.LoginSuccess,
			"msg":  "登录成功",
		}
	})
}

//UserInfo 显示用户信息路由
func (u *User) UserInfo(c *gin.Context) {
	HandlerResponse(c, func() interface{} {
		tokenKey, ok := c.Get(token.CookieTokenKey)
		// 如果没有在context中取到token说明未登录
		if !ok {
			return errs.NewTokenError(http.StatusForbidden, code.RequestForbidden, "请先登录后再进行操作")
		}
		// 在redis中取不到token
		userToken, err := u.serviceUser.GetUserToken(tokenKey.(string))
		if err != nil {
			return errs.NewTokenError(http.StatusForbidden, code.RequestForbidden, "请先登录后再进行操作")
		}
		// 查看token是否已经过期
		if time.Now().Unix()-userToken.CreateTime > token.TokenMaxAge {
			return errs.NewTokenError(http.StatusForbidden, code.RequestForbidden, "请先登录后再进行操作")
		}
		// 提取电话号码并查询用户信息
		teleNumber := userToken.Info["tele_number"]
		var entityUser entity.User
		u.serviceUser.FindUserByTelephone(teleNumber.(string), &entityUser)
		if entityUser.ID == 0 {
			return errs.NewTokenError(http.StatusForbidden, code.RequestForbidden, "请先登录后再进行操作")
		}
		return gin.H{
			"uid":  entityUser.ID,
			"name": entityUser.Name,
		}
	})
}

func (u *User) Regist(r *gin.Engine) {
	// 注册登录路由
	rg := r.Group("/user")
	{
		rg.POST("/regist", u.UserRegist)
		rg.POST("/login", u.UserLogin)
	}

	// 用户信息路由
	rgInfo := rg.Group("/", middleware.UserMiddleware())
	{
		rgInfo.GET("/info", u.UserInfo)
	}
}
