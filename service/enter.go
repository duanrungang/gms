package service

import (
	"gms/service/example"
	"gms/service/system"
)

// ServiceGroup service组
type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
