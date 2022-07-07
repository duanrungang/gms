package router

import (
	"gms/router/system"
)

// RouterGroup 路由组
type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
