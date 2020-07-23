package iface

type IRouter interface {
	//在处理conn业务之前的
	PreHandle(request IRequest)
	//处理conn业务的主方法
	Handle(request IRequest)
	//处理conn业务之后的方法
	PostHandle(request IRequest)
}