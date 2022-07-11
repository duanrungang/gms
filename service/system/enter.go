package system

type ServiceGroup struct {
	ApiService
	JwtService
	MenuService
	UserService
	CasbinService
	InitDBService
	BaseMenuService
	AuthorityService
	OperationRecordService
}
