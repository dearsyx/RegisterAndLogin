package static

import (
	"net/http"

	"goprojects.com/simple_regist_login/pkg/config"

	"github.com/gin-gonic/gin"
)

type StaticRegistor struct {
}

func NewStaticRegistor() *StaticRegistor {
	return &StaticRegistor{}
}

func (s *StaticRegistor) Regist(r *gin.Engine) {
	rg := r.Group("")
	rg.GET("regist", s.RegistHTML)
	rg.GET("login", s.LoginHTML)
}

func LoadStatic(r *gin.Engine) {
	r.LoadHTMLGlob(config.AppConfig.StaticConfig.TemplatePath)
	r.Static("/js", config.AppConfig.StaticConfig.JavascriptPath)
}

func (s *StaticRegistor) RegistHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "regist.html", nil)
}

func (s *StaticRegistor) LoginHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
