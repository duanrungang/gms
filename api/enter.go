package api

import (
	"gms/api/system"
)

// ApiGroup api组
type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
