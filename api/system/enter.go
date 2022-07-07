package system

import (
	"gms/service"
)

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	AuthorityMenuApi
}

var (
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService   = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)