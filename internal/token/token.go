package token

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	letterList     = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	TokenMaxAge    = 259200       // Token在Cookie中最多保留三天
	CookieTokenKey = `user_token` // Cookie中的token key
)

type UserToken struct {
	Info       gin.H  `json:"info"`
	Token      string `json:"token"`
	CreateTime int64  `json:"create_time"`
}

func GenRandTokenKey() (tokenKey string) {
	rand.Seed(time.Now().Unix())
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterList[rand.Int63()%int64(len(letterList))]
	}
	return string(b)
}
