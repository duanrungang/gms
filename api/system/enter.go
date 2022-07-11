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
	SystemApiApi
	AuthorityApi
	DictionaryApi
	AuthorityMenuApi
}

var (
	apiService          = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService          = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService         = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService         = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService       = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	casbinService       = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	baseMenuService     = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	authorityService    = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	dictionaryService   = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
)
