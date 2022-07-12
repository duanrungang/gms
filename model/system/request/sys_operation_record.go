package request

import (
	"gms/model/common/request"
	"gms/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
