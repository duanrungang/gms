package system

import (
	"github.com/gin-gonic/gin"
	"gms/api"
	"gms/middleware"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := api.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu) // 获取菜单树
	}
	return menuRouter
}
