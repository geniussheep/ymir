package router

const (
	V1 = "/api/v1"
)

var (
	ApiRouterScan = make([]func(), 0)
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	for _, routerScan := range ApiRouterScan {
		routerScan()
	}
}
