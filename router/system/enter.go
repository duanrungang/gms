package system

type RouterGroup struct {
	ApiRouter
	JwtRouter
	BaseRouter
	InitRouter
	MenuRouter
	UserRouter
	AuthorityRouter
}
