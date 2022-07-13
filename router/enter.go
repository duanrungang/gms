package router

import (
	"gms/router/example"
	"gms/router/system"
)

// RouterGroup 路由组
type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
