package system

type RouterGroup struct {
	JwtRouter
	BaseRouter
	InitRouter
	MenuRouter
	UserRouter
	AuthorityRouter
}
