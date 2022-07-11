package system

type ServiceGroup struct {
	ApiService
	JwtService
	MenuService
	UserService
	CasbinService
	InitDBService
	BaseMenuService
	SystemConfigService
	AuthorityService
	DictionaryService
	OperationRecordService
	DictionaryDetailService
}
