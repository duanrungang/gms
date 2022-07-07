package request

import (
	"gms/model/common/request"
	"gms/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
