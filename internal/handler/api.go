package handler

import (
	"net/http"

	"goprojects.com/simple_regist_login/internal/static"

	"goprojects.com/simple_regist_login/internal/dto/errs"

	"goprojects.com/simple_regist_login/internal/domain/user/entity"

	"goprojects.com/simple_regist_login/internal/domain/user/repository"
	"goprojects.com/simple_regist_login/internal/services"

	_interface "goprojects.com/simple_regist_login/internal/services/interface"

	"github.com/gin-gonic/gin"
)

func GroupRegist(r *gin.Engine, registerGroup ..._interface.Register) {
	for _, item := range registerGroup {
		item.Regist(r)
	}
}

func GroupMigrate(migratorGroup ..._interface.Migrator) {
	for _, item := range migratorGroup {
		item.Migrate()
	}
}

func HandlerResponse(c *gin.Context, f func() interface{}) {
	resp := f()
	switch resp.(type) {
	case *errs.DataError:
		err := resp.(*errs.DataError)
		c.JSON(err.Status, gin.H{
			"code": err.Code,
			"msg":  err.Message,
		})
	case *errs.TokenError:
		err := resp.(*errs.TokenError)
		c.JSON(err.Status, gin.H{
			"code": err.Code,
			"msg":  err.Message,
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"data": resp,
		})
	}
}

func StartApi() {
	serviceUser := services.NewServiceUser(repository.NewRepoUser())

	r := gin.Default()

	static.LoadStatic(r)
	GroupRegist(
		r,
		NewHandlerUser(serviceUser),
		static.NewStaticRegistor(),
	)

	GroupMigrate(
		entity.NewEntityUserMigrator(),
	)

	_ = r.Run()
}
