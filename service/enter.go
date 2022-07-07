package service

import (
	"gms/service/system"
)

// ServiceGroup service组
type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
