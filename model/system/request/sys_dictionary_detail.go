package request

import (
	"gms/model/common/request"
	"gms/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
