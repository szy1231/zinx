package iface

//定义一个服务器接口
type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Serve()
	//路由功能，给当前服务注册一个路由
	AddRouter(router IRouter)
}
