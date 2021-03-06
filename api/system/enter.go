package system

import (
	"gms/service"
)

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	SystemApi
	CasbinApi
	AutoCodeApi
	SystemApiApi
	AuthorityApi
	DictionaryApi
	AuthorityMenuApi
	DictionaryDetailApi
	OperationRecordApi
	AutoCodeHistoryApi
	AuthorityBtnApi
}

var (
	apiService              = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService              = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	autoCodeService         = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
	baseMenuService         = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	authorityService        = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	autoCodeHistoryService  = service.ServiceGroupApp.SystemServiceGroup.AutoCodeHistoryService
	operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	authorityBtnService     = service.ServiceGroupApp.SystemServiceGroup.AuthorityBtnService
)
