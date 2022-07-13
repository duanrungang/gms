package system

type ServiceGroup struct {
	ApiService
	JwtService
	MenuService
	UserService
	CasbinService
	InitDBService
	AutoCodeService
	BaseMenuService
	SystemConfigService
	AuthorityService
	DictionaryService
	AutoCodeHistoryService
	OperationRecordService
	DictionaryDetailService
	AuthorityBtnService
}
