package system

import (
	"github.com/gin-gonic/gin"
	"gms/api"
	"gms/middleware"
)

type DictionaryDetailRouter struct{}

func (s *DictionaryDetailRouter) InitSysDictionaryDetailRouter(Router *gin.RouterGroup) {
	dictionaryDetailRouter := Router.Group("sysDictionaryDetail").Use(middleware.OperationRecord())
	dictionaryDetailRouterWithoutRecord := Router.Group("sysDictionaryDetail")
	sysDictionaryDetailApi := api.ApiGroupApp.SystemApiGroup.DictionaryDetailApi
	{
		dictionaryDetailRouter.POST("createSysDictionaryDetail", sysDictionaryDetailApi.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
		dictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", sysDictionaryDetailApi.DeleteSysDictionaryDetail) // 删除SysDictionaryDetail
		dictionaryDetailRouter.PUT("updateSysDictionaryDetail", sysDictionaryDetailApi.UpdateSysDictionaryDetail)    // 更新SysDictionaryDetail
	}
	{
		dictionaryDetailRouterWithoutRecord.GET("findSysDictionaryDetail", sysDictionaryDetailApi.FindSysDictionaryDetail)       // 根据ID获取SysDictionaryDetail
		dictionaryDetailRouterWithoutRecord.GET("getSysDictionaryDetailList", sysDictionaryDetailApi.GetSysDictionaryDetailList) // 获取SysDictionaryDetail列表
	}
}
