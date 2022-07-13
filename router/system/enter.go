package system

type RouterGroup struct {
	ApiRouter
	JwtRouter
	SysRouter
	BaseRouter
	InitRouter
	MenuRouter
	UserRouter
	AutoCodeRouter
	DictionaryRouter
	CasbinRouter
	AuthorityRouter
	AutoCodeHistoryRouter
	DictionaryDetailRouter
	OperationRecordRouter
	AuthorityBtnRouter
}
