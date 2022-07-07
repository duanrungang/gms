package system

import (
	"github.com/gin-gonic/gin"
	"gms/api"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := api.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)   // 创建Api
		initRouter.POST("checkdb", dbApi.CheckDB) // 创建Api
	}
}
