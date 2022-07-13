package api

import (
	"gms/api/example"
	"gms/api/system"
)

// ApiGroup api组
type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
