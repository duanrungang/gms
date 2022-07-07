package initialize

import (
	_ "gms/source/example"
	_ "gms/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
